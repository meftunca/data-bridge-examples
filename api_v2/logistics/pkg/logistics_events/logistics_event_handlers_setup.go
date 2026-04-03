package logistics_handlers

import (
	"github.com/maple-tech/baseline/events"
)

// RegisterAllLogisticsEventHandlers, bu şemadaki tüm tablo hook'larını
// tek seferde kaydeder. Varsayılan olarak tüm tablolar kapalıdır.
// İstediğiniz tablonun satırındaki "//" işaretini kaldırarak aktif edin.
//
// Bu fonksiyonu main.go veya uygulama setup kodunuzda çağırın:
//
//	em := events.NewEventManager()
//	RegisterAllLogisticsEventHandlers(em)
//	api.Setup(app, db, em)
func RegisterAllLogisticsEventHandlers(em *events.EventManager) {
	// RegisterInventoryEventHooks(em)
	// RegisterPurchaseOrderItemsEventHooks(em)
	// RegisterPurchaseOrdersEventHooks(em)
	// RegisterShipmentItemsEventHooks(em)
	// RegisterShipmentTrackingEventHooks(em)
	// RegisterShipmentsEventHooks(em)
	// RegisterStockMovementsEventHooks(em)
	// RegisterStorageBinsEventHooks(em)
	// RegisterStorageZonesEventHooks(em)
	// RegisterSuppliersEventHooks(em)
	// RegisterWarehousesEventHooks(em)
}

// RegisterInventoryEventHooks, Inventory tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterInventoryEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/inventory
	//   Tekil işlem  : /logistics/inventory/with-id/:id
	//   Sayfalama    : /logistics/inventory/pagination
	//   Toplu işlem  : /logistics/inventory/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/inventory/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/inventory/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.InventoryForm)
	// 	// if created, ok := event.Data.(structures.InventoryForm); ok { ... }
	// })
	// em.OnEvent("/logistics/inventory", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/inventory/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/inventory/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.InventoryForm)
	// 	// if created, ok := event.Data.([]structures.InventoryForm); ok { ... }
	// })
	// em.OnEvent("/logistics/inventory/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.InventoryBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.InventoryBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/inventory/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/inventory/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/inventory/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/inventory/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterPurchaseOrderItemsEventHooks, PurchaseOrderItems tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterPurchaseOrderItemsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/purchase-order-items
	//   Tekil işlem  : /logistics/purchase-order-items/with-id/:id
	//   Sayfalama    : /logistics/purchase-order-items/pagination
	//   Toplu işlem  : /logistics/purchase-order-items/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-order-items/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-order-items/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.PurchaseOrderItemsForm)
	// 	// if created, ok := event.Data.(structures.PurchaseOrderItemsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/purchase-order-items", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-order-items/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-order-items/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.PurchaseOrderItemsForm)
	// 	// if created, ok := event.Data.([]structures.PurchaseOrderItemsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/purchase-order-items/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.PurchaseOrderItemsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.PurchaseOrderItemsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/purchase-order-items/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-order-items/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-order-items/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-order-items/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterPurchaseOrdersEventHooks, PurchaseOrders tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterPurchaseOrdersEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/purchase-orders
	//   Tekil işlem  : /logistics/purchase-orders/with-id/:id
	//   Sayfalama    : /logistics/purchase-orders/pagination
	//   Toplu işlem  : /logistics/purchase-orders/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-orders/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-orders/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.PurchaseOrdersForm)
	// 	// if created, ok := event.Data.(structures.PurchaseOrdersForm); ok { ... }
	// })
	// em.OnEvent("/logistics/purchase-orders", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-orders/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-orders/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.PurchaseOrdersForm)
	// 	// if created, ok := event.Data.([]structures.PurchaseOrdersForm); ok { ... }
	// })
	// em.OnEvent("/logistics/purchase-orders/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.PurchaseOrdersBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.PurchaseOrdersBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/purchase-orders/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/purchase-orders/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/purchase-orders/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/purchase-orders/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterShipmentItemsEventHooks, ShipmentItems tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterShipmentItemsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/shipment-items
	//   Tekil işlem  : /logistics/shipment-items/with-id/:id
	//   Sayfalama    : /logistics/shipment-items/pagination
	//   Toplu işlem  : /logistics/shipment-items/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-items/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-items/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ShipmentItemsForm)
	// 	// if created, ok := event.Data.(structures.ShipmentItemsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/shipment-items", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-items/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-items/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ShipmentItemsForm)
	// 	// if created, ok := event.Data.([]structures.ShipmentItemsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/shipment-items/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ShipmentItemsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ShipmentItemsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/shipment-items/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-items/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-items/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-items/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterShipmentTrackingEventHooks, ShipmentTracking tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterShipmentTrackingEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/shipment-tracking
	//   Tekil işlem  : /logistics/shipment-tracking/with-id/:id
	//   Sayfalama    : /logistics/shipment-tracking/pagination
	//   Toplu işlem  : /logistics/shipment-tracking/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-tracking/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-tracking/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ShipmentTrackingForm)
	// 	// if created, ok := event.Data.(structures.ShipmentTrackingForm); ok { ... }
	// })
	// em.OnEvent("/logistics/shipment-tracking", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-tracking/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-tracking/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ShipmentTrackingForm)
	// 	// if created, ok := event.Data.([]structures.ShipmentTrackingForm); ok { ... }
	// })
	// em.OnEvent("/logistics/shipment-tracking/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ShipmentTrackingBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ShipmentTrackingBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/shipment-tracking/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipment-tracking/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipment-tracking/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipment-tracking/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterShipmentsEventHooks, Shipments tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterShipmentsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/shipments
	//   Tekil işlem  : /logistics/shipments/with-id/:id
	//   Sayfalama    : /logistics/shipments/pagination
	//   Toplu işlem  : /logistics/shipments/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipments/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipments/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ShipmentsForm)
	// 	// if created, ok := event.Data.(structures.ShipmentsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/shipments", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipments/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipments/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ShipmentsForm)
	// 	// if created, ok := event.Data.([]structures.ShipmentsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/shipments/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ShipmentsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ShipmentsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/shipments/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/shipments/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/shipments/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/shipments/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterStockMovementsEventHooks, StockMovements tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterStockMovementsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/stock-movements
	//   Tekil işlem  : /logistics/stock-movements/with-id/:id
	//   Sayfalama    : /logistics/stock-movements/pagination
	//   Toplu işlem  : /logistics/stock-movements/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/stock-movements/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/stock-movements/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.StockMovementsForm)
	// 	// if created, ok := event.Data.(structures.StockMovementsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/stock-movements", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/stock-movements/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/stock-movements/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.StockMovementsForm)
	// 	// if created, ok := event.Data.([]structures.StockMovementsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/stock-movements/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.StockMovementsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.StockMovementsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/stock-movements/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/stock-movements/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/stock-movements/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/stock-movements/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterStorageBinsEventHooks, StorageBins tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterStorageBinsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/storage-bins
	//   Tekil işlem  : /logistics/storage-bins/with-id/:id
	//   Sayfalama    : /logistics/storage-bins/pagination
	//   Toplu işlem  : /logistics/storage-bins/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-bins/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-bins/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.StorageBinsForm)
	// 	// if created, ok := event.Data.(structures.StorageBinsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/storage-bins", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-bins/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-bins/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.StorageBinsForm)
	// 	// if created, ok := event.Data.([]structures.StorageBinsForm); ok { ... }
	// })
	// em.OnEvent("/logistics/storage-bins/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.StorageBinsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.StorageBinsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/storage-bins/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-bins/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-bins/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-bins/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterStorageZonesEventHooks, StorageZones tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterStorageZonesEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/storage-zones
	//   Tekil işlem  : /logistics/storage-zones/with-id/:id
	//   Sayfalama    : /logistics/storage-zones/pagination
	//   Toplu işlem  : /logistics/storage-zones/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-zones/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-zones/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.StorageZonesForm)
	// 	// if created, ok := event.Data.(structures.StorageZonesForm); ok { ... }
	// })
	// em.OnEvent("/logistics/storage-zones", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-zones/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-zones/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.StorageZonesForm)
	// 	// if created, ok := event.Data.([]structures.StorageZonesForm); ok { ... }
	// })
	// em.OnEvent("/logistics/storage-zones/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.StorageZonesBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.StorageZonesBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/storage-zones/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/storage-zones/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/storage-zones/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/storage-zones/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterSuppliersEventHooks, Suppliers tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterSuppliersEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/suppliers
	//   Tekil işlem  : /logistics/suppliers/with-id/:id
	//   Sayfalama    : /logistics/suppliers/pagination
	//   Toplu işlem  : /logistics/suppliers/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/suppliers/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/suppliers/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.SuppliersForm)
	// 	// if created, ok := event.Data.(structures.SuppliersForm); ok { ... }
	// })
	// em.OnEvent("/logistics/suppliers", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/suppliers/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/suppliers/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.SuppliersForm)
	// 	// if created, ok := event.Data.([]structures.SuppliersForm); ok { ... }
	// })
	// em.OnEvent("/logistics/suppliers/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.SuppliersBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.SuppliersBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/suppliers/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/suppliers/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/suppliers/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/suppliers/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterWarehousesEventHooks, Warehouses tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllLogisticsEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterWarehousesEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /logistics/warehouses
	//   Tekil işlem  : /logistics/warehouses/with-id/:id
	//   Sayfalama    : /logistics/warehouses/pagination
	//   Toplu işlem  : /logistics/warehouses/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/warehouses/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/warehouses/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.WarehousesForm)
	// 	// if created, ok := event.Data.(structures.WarehousesForm); ok { ... }
	// })
	// em.OnEvent("/logistics/warehouses", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/warehouses/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/warehouses/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.WarehousesForm)
	// 	// if created, ok := event.Data.([]structures.WarehousesForm); ok { ... }
	// })
	// em.OnEvent("/logistics/warehouses/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.WarehousesBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.WarehousesBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/logistics/warehouses/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/logistics/warehouses/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/logistics/warehouses/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/logistics/warehouses/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}
