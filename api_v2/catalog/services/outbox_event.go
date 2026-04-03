package catalog_api_service

import "gorm.io/gorm"

// OutboxEvent represents a transactional outbox record.
// Generated once per schema to avoid duplicate type declarations.
type OutboxEvent struct {
	AggregateType string `gorm:"column:aggregate_type"`
	AggregateID   string `gorm:"column:aggregate_id"`
	EventType     string `gorm:"column:event_type"`
	Payload       string `gorm:"column:payload"`
}

func (OutboxEvent) TableName() string { return "outbox_events" }

var _ = gorm.Model{}
