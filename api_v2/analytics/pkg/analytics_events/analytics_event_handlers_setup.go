package analytics_handlers

import (
	"github.com/maple-tech/baseline/events"
)

// RegisterAllAnalyticsEventHandlers, bu şemadaki tüm tablo hook'larını
// tek seferde kaydeder. Varsayılan olarak tüm tablolar kapalıdır.
// İstediğiniz tablonun satırındaki "//" işaretini kaldırarak aktif edin.
//
// Bu fonksiyonu main.go veya uygulama setup kodunuzda çağırın:
//
//	em := events.NewEventManager()
//	RegisterAllAnalyticsEventHandlers(em)
//	api.Setup(app, db, em)
func RegisterAllAnalyticsEventHandlers(em *events.EventManager) {
	// RegisterAlertHistoryEventHooks(em)
	// RegisterAlertRulesEventHooks(em)
	// RegisterAuditLogsEventHooks(em)
	// RegisterDashboardWidgetsEventHooks(em)
	// RegisterDashboardsEventHooks(em)
	// RegisterEventsEventHooks(em)
	// RegisterMetricsEventHooks(em)
	// RegisterNotificationsEventHooks(em)
	// RegisterRecentEventsEventHooks(em)
	// RegisterReportExecutionsEventHooks(em)
	// RegisterReportsEventHooks(em)
	// RegisterUnreadNotificationsEventHooks(em)
}

// RegisterAlertHistoryEventHooks, AlertHistory tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterAlertHistoryEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/alert-history
	//   Tekil işlem  : /analytics/alert-history/with-id/:id
	//   Sayfalama    : /analytics/alert-history/pagination
	//   Toplu işlem  : /analytics/alert-history/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-history/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-history/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.AlertHistoryForm)
	// 	// if created, ok := event.Data.(structures.AlertHistoryForm); ok { ... }
	// })
	// em.OnEvent("/analytics/alert-history", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-history/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-history/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.AlertHistoryForm)
	// 	// if created, ok := event.Data.([]structures.AlertHistoryForm); ok { ... }
	// })
	// em.OnEvent("/analytics/alert-history/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.AlertHistoryBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.AlertHistoryBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/alert-history/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-history/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-history/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-history/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterAlertRulesEventHooks, AlertRules tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterAlertRulesEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/alert-rules
	//   Tekil işlem  : /analytics/alert-rules/with-id/:id
	//   Sayfalama    : /analytics/alert-rules/pagination
	//   Toplu işlem  : /analytics/alert-rules/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-rules/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-rules/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.AlertRulesForm)
	// 	// if created, ok := event.Data.(structures.AlertRulesForm); ok { ... }
	// })
	// em.OnEvent("/analytics/alert-rules", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-rules/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-rules/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.AlertRulesForm)
	// 	// if created, ok := event.Data.([]structures.AlertRulesForm); ok { ... }
	// })
	// em.OnEvent("/analytics/alert-rules/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.AlertRulesBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.AlertRulesBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/alert-rules/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/alert-rules/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/alert-rules/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/alert-rules/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterAuditLogsEventHooks, AuditLogs tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterAuditLogsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/audit-logs
	//   Tekil işlem  : /analytics/audit-logs/with-id/:id
	//   Sayfalama    : /analytics/audit-logs/pagination
	//   Toplu işlem  : /analytics/audit-logs/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/audit-logs/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/audit-logs/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.AuditLogsForm)
	// 	// if created, ok := event.Data.(structures.AuditLogsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/audit-logs", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/audit-logs/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/audit-logs/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.AuditLogsForm)
	// 	// if created, ok := event.Data.([]structures.AuditLogsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/audit-logs/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.AuditLogsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.AuditLogsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/audit-logs/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/audit-logs/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/audit-logs/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/audit-logs/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterDashboardWidgetsEventHooks, DashboardWidgets tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterDashboardWidgetsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/dashboard-widgets
	//   Tekil işlem  : /analytics/dashboard-widgets/with-id/:id
	//   Sayfalama    : /analytics/dashboard-widgets/pagination
	//   Toplu işlem  : /analytics/dashboard-widgets/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboard-widgets/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboard-widgets/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.DashboardWidgetsForm)
	// 	// if created, ok := event.Data.(structures.DashboardWidgetsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/dashboard-widgets", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboard-widgets/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboard-widgets/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.DashboardWidgetsForm)
	// 	// if created, ok := event.Data.([]structures.DashboardWidgetsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/dashboard-widgets/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.DashboardWidgetsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.DashboardWidgetsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/dashboard-widgets/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboard-widgets/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboard-widgets/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboard-widgets/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterDashboardsEventHooks, Dashboards tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterDashboardsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/dashboards
	//   Tekil işlem  : /analytics/dashboards/with-id/:id
	//   Sayfalama    : /analytics/dashboards/pagination
	//   Toplu işlem  : /analytics/dashboards/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboards/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboards/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.DashboardsForm)
	// 	// if created, ok := event.Data.(structures.DashboardsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/dashboards", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboards/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboards/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.DashboardsForm)
	// 	// if created, ok := event.Data.([]structures.DashboardsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/dashboards/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.DashboardsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.DashboardsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/dashboards/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/dashboards/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/dashboards/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/dashboards/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterEventsEventHooks, Events tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterEventsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/events
	//   Tekil işlem  : /analytics/events/with-id/:id
	//   Sayfalama    : /analytics/events/pagination
	//   Toplu işlem  : /analytics/events/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/events/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/events/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.EventsForm)
	// 	// if created, ok := event.Data.(structures.EventsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/events", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/events/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/events/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.EventsForm)
	// 	// if created, ok := event.Data.([]structures.EventsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/events/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.EventsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.EventsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/events/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/events/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/events/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/events/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterMetricsEventHooks, Metrics tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterMetricsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/metrics
	//   Tekil işlem  : /analytics/metrics/with-id/:id
	//   Sayfalama    : /analytics/metrics/pagination
	//   Toplu işlem  : /analytics/metrics/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/metrics/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/metrics/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.MetricsForm)
	// 	// if created, ok := event.Data.(structures.MetricsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/metrics", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/metrics/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/metrics/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.MetricsForm)
	// 	// if created, ok := event.Data.([]structures.MetricsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/metrics/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.MetricsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.MetricsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/metrics/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/metrics/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/metrics/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/metrics/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterNotificationsEventHooks, Notifications tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterNotificationsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/notifications
	//   Tekil işlem  : /analytics/notifications/with-id/:id
	//   Sayfalama    : /analytics/notifications/pagination
	//   Toplu işlem  : /analytics/notifications/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/notifications/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/notifications/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.NotificationsForm)
	// 	// if created, ok := event.Data.(structures.NotificationsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/notifications", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/notifications/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/notifications/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.NotificationsForm)
	// 	// if created, ok := event.Data.([]structures.NotificationsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/notifications/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.NotificationsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.NotificationsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/notifications/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/notifications/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/notifications/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/notifications/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterRecentEventsEventHooks, RecentEvents tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterRecentEventsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/recent-events
	//   Tekil işlem  : /analytics/recent-events/with-id
	//   Sayfalama    : /analytics/recent-events/pagination
	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/recent-events/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/recent-events/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/recent-events/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/recent-events/with-id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/recent-events/with-id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/recent-events/with-id", events.FindQueryError, func(event events.Event) {})

}

// RegisterReportExecutionsEventHooks, ReportExecutions tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterReportExecutionsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/report-executions
	//   Tekil işlem  : /analytics/report-executions/with-id/:id
	//   Sayfalama    : /analytics/report-executions/pagination
	//   Toplu işlem  : /analytics/report-executions/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/report-executions/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/report-executions/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ReportExecutionsForm)
	// 	// if created, ok := event.Data.(structures.ReportExecutionsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/report-executions", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/report-executions/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/report-executions/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ReportExecutionsForm)
	// 	// if created, ok := event.Data.([]structures.ReportExecutionsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/report-executions/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ReportExecutionsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ReportExecutionsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/report-executions/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/report-executions/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/report-executions/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/report-executions/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterReportsEventHooks, Reports tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterReportsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/reports
	//   Tekil işlem  : /analytics/reports/with-id/:id
	//   Sayfalama    : /analytics/reports/pagination
	//   Toplu işlem  : /analytics/reports/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/reports/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/reports/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ReportsForm)
	// 	// if created, ok := event.Data.(structures.ReportsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/reports", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/reports/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/reports/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ReportsForm)
	// 	// if created, ok := event.Data.([]structures.ReportsForm); ok { ... }
	// })
	// em.OnEvent("/analytics/reports/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ReportsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ReportsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/analytics/reports/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/reports/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/reports/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/reports/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterUnreadNotificationsEventHooks, UnreadNotifications tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllAnalyticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterUnreadNotificationsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /analytics/unread-notifications
	//   Tekil işlem  : /analytics/unread-notifications/with-id
	//   Sayfalama    : /analytics/unread-notifications/pagination
	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/unread-notifications/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/unread-notifications/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/unread-notifications/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/analytics/unread-notifications/with-id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/analytics/unread-notifications/with-id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/analytics/unread-notifications/with-id", events.FindQueryError, func(event events.Event) {})

}
