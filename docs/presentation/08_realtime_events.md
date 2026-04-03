
# Real-Time Events and SSE

## Server-Sent Events Hub

The SSE package provides a production-grade Server-Sent Events
implementation for real-time data streaming over HTTP.

### Architecture

```
Client A ----\                              /----> Subscriber 1
Client B -----+----> SSE Hub (per topic) --+-----> Subscriber 2
Client C ----/       |                      \----> Subscriber 3
                     |
                 Backplane (optional)
                 Redis / NATS
                     |
                 Other Pods
```

Each topic (e.g., `iam.users`) maintains its own subscriber list.
Events are fan-out broadcast to all subscribers of a topic.

---

# SSE Hub Implementation

```go
type Hub struct {
    subscribers map[string]map[string]chan Event
    mu          sync.RWMutex
    backplane   Backplane
    bufferSize  int
    pingInterval time.Duration
}

type Event struct {
    Topic   string
    Type    string          // "created", "updated", "deleted"
    Payload json.RawMessage
}
```

**Key Characteristics:**
- Thread-safe subscriber management via `sync.RWMutex`
- Non-blocking fan-out: slow consumers receive dropped events, not deadlocks
- Configurable buffer size per subscriber channel
- Automatic keep-alive pings at configurable intervals (default: 30s)
- Graceful cleanup on client disconnect

---

# SSE Stream Handler

The SSE hub exposes a Fiber-compatible handler:

```go
func (h *Hub) StreamResource(topic string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        c.Set("Content-Type", "text/event-stream")
        c.Set("Cache-Control", "no-cache")
        c.Set("Connection", "keep-alive")
        c.Set("X-Accel-Buffering", "no")

        subscriberID := uuid.New().String()
        ch := h.Subscribe(topic, subscriberID)
        defer h.Unsubscribe(topic, subscriberID)

        c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
            ticker := time.NewTicker(h.pingInterval)
            defer ticker.Stop()

            for {
                select {
                case event := <-ch:
                    fmt.Fprintf(w, "event: %s\ndata: %s\n\n",
                        event.Type, event.Payload)
                    w.Flush()
                case <-ticker.C:
                    fmt.Fprintf(w, ": ping\n\n")
                    w.Flush()
                case <-c.Context().Done():
                    return
                }
            }
        })
        return nil
    }
}
```

---

# Backplane Interface

For multi-pod deployments, the SSE hub supports pluggable backplanes:

```go
type Backplane interface {
    Publish(topic string, event Event) error
    Subscribe(topic string) (<-chan Event, error)
    Unsubscribe(topic string) error
    Close() error
}
```

**Supported Implementations:**
- **Redis Pub/Sub** -- Standard cross-pod broadcast
- **NATS** -- High-throughput message bus

**Configuration:**
```go
hub := sse.NewHub(
    sse.WithBackplane(redisBackplane),
    sse.WithBufferSize(256),
    sse.WithPingInterval(15 * time.Second),
)
```

Without a backplane, events are local to the process.
Adding a backplane requires zero changes to generated code.

---

# Generated SSE Integration

When `sseEnabled: true`, the generator produces:

**1. Stream endpoint in route registration:**
```go
users.Get("/stream", hub.StreamResource("iam.users"))
```

**2. Event emission in controllers:**
```go
func (ctrl *UserController) Create(c *fiber.Ctx) error {
    result, err := ctrl.Service.Create(&form)
    if err == nil {
        // Publish to SSE subscribers
        sse.Publish("iam.users", sse.Event{
            Type:    "created",
            Payload: sonic.Marshal(result),
        })
    }
    return c.Status(201).JSON(result)
}
```

Clients connect via `EventSource` and receive real-time updates
for any table they subscribe to.

---

# Event Manager System

The `EventManager` provides route-scoped lifecycle hooks
for cross-cutting concerns.

### Event Types (Organized in Blocks of 100)

| Block | Event Types |
|-------|------------|
| 100-199 | Creation: `CreationSuccess`, `CreationError`, `CreationRequest` |
| 200-299 | Update: `UpdateSuccess`, `UpdateError`, `UpdateRequest` |
| 300-399 | Deletion: `DeletionSuccess`, `DeletionError`, `DeletionRequest` |
| 400-499 | Query: `PaginationQuerySuccess`, `SearchSuccess`, `FindSuccess` |
| 500-599 | Auth: `AuthorizationError`, `AuthenticationError` |
| 600-699 | Validation: `ValidationError`, `ParseError` |

Each event carries the Fiber context, error (if any), status code, and data payload.

---

# Event Handler Registration

Three dispatch modes for different use cases:

```go
em := events.NewEventManager()

// 1. Request interceptor (can abort the request)
em.OnRequest("/api/v1/iam/users", events.CreationRequest,
    func(e events.Event) error {
        // Validate, transform, or reject
        if !isAllowed(e.Ctx) {
            return fiber.ErrForbidden
        }
        return nil // continue processing
    },
)

// 2. Synchronous post-event handler
em.OnEvent("/api/v1/iam/users", events.CreationSuccess,
    func(e events.Event) {
        auditLog.Record(e)  // runs in request goroutine
    },
)

// 3. Asynchronous post-event handler
em.OnEventAsync("/api/v1/iam/users", events.CreationSuccess,
    func(e events.Event) {
        notificationService.Send(e)  // runs in separate goroutine
    },
)
```

---

# Event Manager: Wildcard Handlers

Global handlers that apply to all routes:

```go
// Log every successful creation across all resources
em.OnAnyEvent(events.CreationSuccess, func(e events.Event) {
    log.Info("Resource created",
        "route", e.RoutePath,
        "status", e.StatusCode,
    )
})

// Global error tracking
em.OnAnyEvent(events.CreationError, func(e events.Event) {
    errorTracker.Capture(e.Err, map[string]string{
        "route":  e.RoutePath,
        "status": fmt.Sprint(e.StatusCode),
    })
})

// Global request audit
em.OnAnyRequest(events.CreationRequest, func(e events.Event) error {
    rateLimiter.Check(e.Ctx)
    return nil
})
```

**Thread Safety**: All handler registration and dispatch
is protected by `sync.RWMutex` for concurrent access.

---

# Outbox Pattern

When `outboxEnabled: true`, the generator adds transactional
event persistence alongside database writes.

```
+------------------+       +------------------+
|  Service.Create  |       |   Outbox Table   |
|  (Transaction)   | ----> |   outbox_events  |
+------------------+       +------------------+
                                    |
                            Async Processor
                                    |
                            +------------------+
                            | Message Broker   |
                            | (Kafka, RabbitMQ)|
                            +------------------+
```

**Guarantees:**
- Business data and event are committed atomically
- No event is lost if the application crashes after commit
- Events are processed exactly-once by the outbox consumer
- Decouples synchronous API response from async processing

This eliminates the dual-write problem common in microservice architectures.
