package analytics_api_structure

// Bu dosya, veritabanındaki ENUM tiplerinden otomatik olarak üretilmiştir.

// AnalyticsEventSeverity represents the 'event_severity' enum type from the 'analytics' schema.
type AnalyticsEventSeverity string

const (
	AnalyticsEventSeverityDebug    AnalyticsEventSeverity = "debug"
	AnalyticsEventSeverityInfo     AnalyticsEventSeverity = "info"
	AnalyticsEventSeverityWarning  AnalyticsEventSeverity = "warning"
	AnalyticsEventSeverityError    AnalyticsEventSeverity = "error"
	AnalyticsEventSeverityCritical AnalyticsEventSeverity = "critical"
)

// AnalyticsReportType represents the 'report_type' enum type from the 'analytics' schema.
type AnalyticsReportType string

const (
	AnalyticsReportTypeDaily     AnalyticsReportType = "daily"
	AnalyticsReportTypeWeekly    AnalyticsReportType = "weekly"
	AnalyticsReportTypeMonthly   AnalyticsReportType = "monthly"
	AnalyticsReportTypeQuarterly AnalyticsReportType = "quarterly"
	AnalyticsReportTypeAnnual    AnalyticsReportType = "annual"
	AnalyticsReportTypeCustom    AnalyticsReportType = "custom"
)
