package orders_handlers

import (
	"github.com/maple-tech/baseline/events"
)

// RegisterAllOrdersEventHandlers, bu şemadaki tüm tablo hook'larını
// tek seferde kaydeder. Varsayılan olarak tüm tablolar kapalıdır.
// İstediğiniz tablonun satırındaki "//" işaretini kaldırarak aktif edin.
//
// Bu fonksiyonu main.go veya uygulama setup kodunuzda çağırın:
//
//	em := events.NewEventManager()
//	RegisterAllOrdersEventHandlers(em)
//	api.Setup(app, db, em)
func RegisterAllOrdersEventHandlers(em *events.EventManager) {
	// RegisterCartItemsEventHooks(em)
	// RegisterCartsEventHooks(em)
	// RegisterCouponsEventHooks(em)
	// RegisterCustomersEventHooks(em)
	// RegisterOrderItemsEventHooks(em)
	// RegisterOrderStatusHistoryEventHooks(em)
	// RegisterOrdersEventHooks(em)
	// RegisterPaymentsEventHooks(em)
	// RegisterRefundsEventHooks(em)
}

// RegisterCartItemsEventHooks, CartItems tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterCartItemsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/cart-items
	//   Tekil işlem  : /orders/cart-items/with-id/:id
	//   Sayfalama    : /orders/cart-items/pagination
	//   Toplu işlem  : /orders/cart-items/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/cart-items/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/cart-items/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.CartItemsForm)
	// 	// if created, ok := event.Data.(structures.CartItemsForm); ok { ... }
	// })
	// em.OnEvent("/orders/cart-items", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/cart-items/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/cart-items/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.CartItemsForm)
	// 	// if created, ok := event.Data.([]structures.CartItemsForm); ok { ... }
	// })
	// em.OnEvent("/orders/cart-items/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.CartItemsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.CartItemsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/cart-items/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/cart-items/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/cart-items/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/cart-items/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterCartsEventHooks, Carts tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterCartsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/carts
	//   Tekil işlem  : /orders/carts/with-id/:id
	//   Sayfalama    : /orders/carts/pagination
	//   Toplu işlem  : /orders/carts/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/carts/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/carts/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.CartsForm)
	// 	// if created, ok := event.Data.(structures.CartsForm); ok { ... }
	// })
	// em.OnEvent("/orders/carts", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/carts/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/carts/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.CartsForm)
	// 	// if created, ok := event.Data.([]structures.CartsForm); ok { ... }
	// })
	// em.OnEvent("/orders/carts/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.CartsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.CartsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/carts/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/carts/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/carts/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/carts/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterCouponsEventHooks, Coupons tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterCouponsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/coupons
	//   Tekil işlem  : /orders/coupons/with-id/:id
	//   Sayfalama    : /orders/coupons/pagination
	//   Toplu işlem  : /orders/coupons/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/coupons/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/coupons/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.CouponsForm)
	// 	// if created, ok := event.Data.(structures.CouponsForm); ok { ... }
	// })
	// em.OnEvent("/orders/coupons", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/coupons/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/coupons/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.CouponsForm)
	// 	// if created, ok := event.Data.([]structures.CouponsForm); ok { ... }
	// })
	// em.OnEvent("/orders/coupons/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.CouponsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.CouponsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/coupons/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/coupons/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/coupons/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/coupons/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterCustomersEventHooks, Customers tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterCustomersEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/customers
	//   Tekil işlem  : /orders/customers/with-id/:id
	//   Sayfalama    : /orders/customers/pagination
	//   Toplu işlem  : /orders/customers/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/customers/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/customers/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.CustomersForm)
	// 	// if created, ok := event.Data.(structures.CustomersForm); ok { ... }
	// })
	// em.OnEvent("/orders/customers", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/customers/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/customers/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.CustomersForm)
	// 	// if created, ok := event.Data.([]structures.CustomersForm); ok { ... }
	// })
	// em.OnEvent("/orders/customers/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.CustomersBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.CustomersBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/customers/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/customers/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/customers/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/customers/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterOrderItemsEventHooks, OrderItems tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterOrderItemsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/order-items
	//   Tekil işlem  : /orders/order-items/with-id/:id
	//   Sayfalama    : /orders/order-items/pagination
	//   Toplu işlem  : /orders/order-items/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-items/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-items/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.OrderItemsForm)
	// 	// if created, ok := event.Data.(structures.OrderItemsForm); ok { ... }
	// })
	// em.OnEvent("/orders/order-items", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-items/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-items/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.OrderItemsForm)
	// 	// if created, ok := event.Data.([]structures.OrderItemsForm); ok { ... }
	// })
	// em.OnEvent("/orders/order-items/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.OrderItemsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.OrderItemsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/order-items/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-items/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-items/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-items/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterOrderStatusHistoryEventHooks, OrderStatusHistory tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterOrderStatusHistoryEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/order-status-history
	//   Tekil işlem  : /orders/order-status-history/with-id/:id
	//   Sayfalama    : /orders/order-status-history/pagination
	//   Toplu işlem  : /orders/order-status-history/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-status-history/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-status-history/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.OrderStatusHistoryForm)
	// 	// if created, ok := event.Data.(structures.OrderStatusHistoryForm); ok { ... }
	// })
	// em.OnEvent("/orders/order-status-history", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-status-history/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-status-history/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.OrderStatusHistoryForm)
	// 	// if created, ok := event.Data.([]structures.OrderStatusHistoryForm); ok { ... }
	// })
	// em.OnEvent("/orders/order-status-history/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.OrderStatusHistoryBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.OrderStatusHistoryBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/order-status-history/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/order-status-history/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/order-status-history/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/order-status-history/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterOrdersEventHooks, Orders tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterOrdersEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/orders
	//   Tekil işlem  : /orders/orders/with-id/:id
	//   Sayfalama    : /orders/orders/pagination
	//   Toplu işlem  : /orders/orders/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/orders/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/orders/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.OrdersForm)
	// 	// if created, ok := event.Data.(structures.OrdersForm); ok { ... }
	// })
	// em.OnEvent("/orders/orders", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/orders/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/orders/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.OrdersForm)
	// 	// if created, ok := event.Data.([]structures.OrdersForm); ok { ... }
	// })
	// em.OnEvent("/orders/orders/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.OrdersBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.OrdersBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/orders/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/orders/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/orders/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/orders/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterPaymentsEventHooks, Payments tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterPaymentsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/payments
	//   Tekil işlem  : /orders/payments/with-id/:id
	//   Sayfalama    : /orders/payments/pagination
	//   Toplu işlem  : /orders/payments/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/payments/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/payments/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.PaymentsForm)
	// 	// if created, ok := event.Data.(structures.PaymentsForm); ok { ... }
	// })
	// em.OnEvent("/orders/payments", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/payments/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/payments/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.PaymentsForm)
	// 	// if created, ok := event.Data.([]structures.PaymentsForm); ok { ... }
	// })
	// em.OnEvent("/orders/payments/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.PaymentsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.PaymentsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/payments/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/payments/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/payments/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/payments/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterRefundsEventHooks, Refunds tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllOrdersEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterRefundsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /orders/refunds
	//   Tekil işlem  : /orders/refunds/with-id/:id
	//   Sayfalama    : /orders/refunds/pagination
	//   Toplu işlem  : /orders/refunds/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/refunds/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/orders/refunds/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.RefundsForm)
	// 	// if created, ok := event.Data.(structures.RefundsForm); ok { ... }
	// })
	// em.OnEvent("/orders/refunds", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/refunds/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/refunds/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.RefundsForm)
	// 	// if created, ok := event.Data.([]structures.RefundsForm); ok { ... }
	// })
	// em.OnEvent("/orders/refunds/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.RefundsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.RefundsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/orders/refunds/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/orders/refunds/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/orders/refunds/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/orders/refunds/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}
