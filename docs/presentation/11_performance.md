
# Performance Engineering

## Concurrency and Resource Management

DataBridge V2 is designed for high throughput code generation.
Multiple mechanisms ensure fast execution even for large schemas.

### Performance Techniques

| Technique | Component | Purpose |
|-----------|-----------|---------|
| Worker Pool | `parser/worker_pool.go` | Parallel table processing |
| Parallel Creator | `parser/parallel_creator.go` | Concurrent file generation |
| Template Cache | `parser/template_cache.go` | One-time template parsing |
| File Writer Pool | `helpers/optimized_file_writer.go` | Buffered I/O pooling |
| String Pool | `helpers/string_pool.go` | String builder reuse |
| Project Semaphore | `app/run.go` | Bounded concurrent projects |

---

# Worker Pool Architecture

The `WorkerPool` provides bounded concurrency with panic recovery:

```go
type WorkerPool struct {
    workers    int
    jobQueue   chan Job
    results    chan Result
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

type Job struct {
    ID       int
    Execute  func() (interface{}, error)
}

type Result struct {
    JobID    int
    Output   interface{}
    Error    error
    Duration time.Duration
}
```

**Default workers:** `runtime.NumCPU()`

Each worker runs in its own goroutine with `defer recover()`
to prevent a single table failure from crashing the entire run.

---

# Parallel Table Generation

The `ParallelCreator` distributes table processing across workers:

```
                    Schema: "iam" (11 tables)
                              |
                    ParallelCreator.Generate()
                              |
              +-------+-------+-------+-------+
              |       |       |       |       |
           Worker1  Worker2  Worker3  Worker4  ...
           users    roles    perms   sessions
              |       |       |       |
            struct  struct  struct  struct
            service service service service
            ctrl    ctrl    ctrl    ctrl
            test    test    test    test
              |       |       |       |
              +-------+-------+-------+
                              |
                    All files written
                              |
                    EndPointCreator (sequential)
```

Struct, service, controller, and test files for different tables
are independent -- they can safely run in parallel.

Route registration runs sequentially after all tables complete
because it aggregates all table metadata into a single file.

---

# Object Pooling

Three `sync.Pool` implementations reduce GC pressure:

### FileWriterPool
```go
var writerPool = sync.Pool{
    New: func() interface{} {
        return bufio.NewWriterSize(nil, 64*1024) // 64KB buffer
    },
}
```
Reuses buffered writers across file generation calls.
Directory existence is cached to avoid redundant `os.MkdirAll` calls.

### StringBuilderPool
```go
var builderPool = sync.Pool{
    New: func() interface{} {
        return &strings.Builder{}
    },
}

func FastStringConcat(parts ...string) string {
    b := builderPool.Get().(*strings.Builder)
    defer builderPool.Put(b)
    b.Reset()
    // Pre-calculate total length
    totalLen := 0
    for _, p := range parts {
        totalLen += len(p)
    }
    b.Grow(totalLen)
    for _, p := range parts {
        b.WriteString(p)
    }
    return b.String()
}
```

---

# Template Cache Performance

Templates are parsed once and cloned for each execution:

```
Startup:
  Parse 13 templates -> Cache (sync.Once)

Per Table Execution:
  1. template.Clone()        ~50ns (vs ~5ms parse)
  2. Execute(clone, data)    ~1ms
  3. Write to buffered I/O   ~0.5ms

Total per table: ~1.5ms (vs ~6.5ms without cache)
Speedup: 4.3x per table operation
```

For a schema with 50 tables generating 4 files each,
the cache saves approximately **1 second** of template parsing time.

---

# Performance Monitor

The `PerformanceMonitor` tracks real-time metrics during generation:

```go
type PerformanceMonitor struct {
    operations map[string]*OperationMetrics
    mu         sync.RWMutex
    startTime  time.Time
}

type OperationMetrics struct {
    Count       int64
    TotalTime   time.Duration
    MinTime     time.Duration
    MaxTime     time.Duration
    MemoryDelta int64
}
```

**Tracked Operations:**
- Schema introspection time
- Per-table struct/service/controller generation
- Template execution duration
- File I/O duration
- Total memory allocation

---

# Benchmark Results

Measured on a schema with 50+ tables across 5 schemas:

### Sequential vs Parallel Generation

| Mode | Tables | Duration | Files | Speed |
|------|--------|----------|-------|-------|
| Sequential | 54 | ~8.2s | 312 | 38 files/s |
| Parallel (8 workers) | 54 | ~2.4s | 312 | 130 files/s |

**Speedup: 3.4x** with NumCPU workers.

### Per-Operation Breakdown

| Operation | Avg Duration | Percentage |
|-----------|-------------|-----------|
| Schema Introspection | 120ms | 5% |
| Template Execution | 0.8ms/file | 35% |
| File I/O | 0.4ms/file | 15% |
| goimports Post-Processing | 1.1s | 45% |

The bottleneck is `goimports` post-processing, which runs
sequentially across all generated files.

---

# Memory Efficiency

### Allocation Strategy

| Technique | Savings |
|-----------|---------|
| `sync.Pool` for writers | ~60% fewer allocations per file |
| `sync.Pool` for builders | ~40% fewer string allocations |
| Template clone vs re-parse | ~90% fewer template allocations |
| Pre-sized slice/map init | ~30% fewer slice grows |
| `sonic.Marshal` vs `json` | ~50% faster JSON encoding |

### Memory Profile (54 tables, 312 files)

```
Before optimization:
  Total Alloc: 285 MB
  Num GC:      142

After optimization:
  Total Alloc: 127 MB   (-55%)
  Num GC:       61       (-57%)
```

---

# Concurrency Safety

All concurrent access points are protected:

| Resource | Protection |
|----------|-----------|
| Template Cache | `sync.RWMutex` + `sync.Once` |
| File Writer Pool | `sync.Pool` (inherently safe) |
| String Builder Pool | `sync.Pool` (inherently safe) |
| SSE Subscribers | `sync.RWMutex` |
| Event Handlers | `sync.RWMutex` |
| Performance Monitor | `sync.RWMutex` |
| Worker Pool Jobs | Buffered channels |
| Project Semaphore | Buffered channel (cap 2) |

**No global mutable state.**
All state is either immutable (`GeneratorContext`), pool-managed,
or mutex-protected. This enables safe concurrent project generation
without race conditions.
