package catalog_handlers

import (
	"github.com/maple-tech/baseline/events"
)

// RegisterAllCatalogEventHandlers, bu şemadaki tüm tablo hook'larını
// tek seferde kaydeder. Varsayılan olarak tüm tablolar kapalıdır.
// İstediğiniz tablonun satırındaki "//" işaretini kaldırarak aktif edin.
//
// Bu fonksiyonu main.go veya uygulama setup kodunuzda çağırın:
//
//	em := events.NewEventManager()
//	RegisterAllCatalogEventHandlers(em)
//	api.Setup(app, db, em)
func RegisterAllCatalogEventHandlers(em *events.EventManager) {
	// RegisterBrandsEventHooks(em)
	// RegisterCategoriesEventHooks(em)
	// RegisterCollectionProductsEventHooks(em)
	// RegisterCollectionsEventHooks(em)
	// RegisterPriceHistoryEventHooks(em)
	// RegisterProductMediaEventHooks(em)
	// RegisterProductReviewsEventHooks(em)
	// RegisterProductTagsEventHooks(em)
	// RegisterProductVariantsEventHooks(em)
	// RegisterProductsEventHooks(em)
	// RegisterTagsEventHooks(em)
}

// RegisterBrandsEventHooks, Brands tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterBrandsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/brands
	//   Tekil işlem  : /catalog/brands/with-id/:id
	//   Sayfalama    : /catalog/brands/pagination
	//   Toplu işlem  : /catalog/brands/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/brands/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/brands/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.BrandsForm)
	// 	// if created, ok := event.Data.(structures.BrandsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/brands", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/brands/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/brands/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.BrandsForm)
	// 	// if created, ok := event.Data.([]structures.BrandsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/brands/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.BrandsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.BrandsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/brands/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/brands/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/brands/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/brands/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterCategoriesEventHooks, Categories tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterCategoriesEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/categories
	//   Tekil işlem  : /catalog/categories/with-id/:id
	//   Sayfalama    : /catalog/categories/pagination
	//   Toplu işlem  : /catalog/categories/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/categories/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/categories/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.CategoriesForm)
	// 	// if created, ok := event.Data.(structures.CategoriesForm); ok { ... }
	// })
	// em.OnEvent("/catalog/categories", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/categories/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/categories/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.CategoriesForm)
	// 	// if created, ok := event.Data.([]structures.CategoriesForm); ok { ... }
	// })
	// em.OnEvent("/catalog/categories/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.CategoriesBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.CategoriesBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/categories/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/categories/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/categories/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/categories/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterCollectionProductsEventHooks, CollectionProducts tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterCollectionProductsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/collection-products
	//   Tekil işlem  : /catalog/collection-products/with-id/:id
	//   Sayfalama    : /catalog/collection-products/pagination
	//   Toplu işlem  : /catalog/collection-products/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collection-products/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collection-products/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.CollectionProductsForm)
	// 	// if created, ok := event.Data.(structures.CollectionProductsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/collection-products", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collection-products/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collection-products/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.CollectionProductsForm)
	// 	// if created, ok := event.Data.([]structures.CollectionProductsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/collection-products/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.CollectionProductsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.CollectionProductsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/collection-products/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collection-products/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collection-products/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collection-products/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterCollectionsEventHooks, Collections tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterCollectionsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/collections
	//   Tekil işlem  : /catalog/collections/with-id/:id
	//   Sayfalama    : /catalog/collections/pagination
	//   Toplu işlem  : /catalog/collections/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collections/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collections/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.CollectionsForm)
	// 	// if created, ok := event.Data.(structures.CollectionsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/collections", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collections/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collections/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.CollectionsForm)
	// 	// if created, ok := event.Data.([]structures.CollectionsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/collections/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.CollectionsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.CollectionsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/collections/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/collections/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/collections/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/collections/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterPriceHistoryEventHooks, PriceHistory tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterPriceHistoryEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/price-history
	//   Tekil işlem  : /catalog/price-history/with-id/:id
	//   Sayfalama    : /catalog/price-history/pagination
	//   Toplu işlem  : /catalog/price-history/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/price-history/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/price-history/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.PriceHistoryForm)
	// 	// if created, ok := event.Data.(structures.PriceHistoryForm); ok { ... }
	// })
	// em.OnEvent("/catalog/price-history", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/price-history/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/price-history/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.PriceHistoryForm)
	// 	// if created, ok := event.Data.([]structures.PriceHistoryForm); ok { ... }
	// })
	// em.OnEvent("/catalog/price-history/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.PriceHistoryBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.PriceHistoryBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/price-history/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/price-history/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/price-history/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/price-history/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterProductMediaEventHooks, ProductMedia tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterProductMediaEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/product-media
	//   Tekil işlem  : /catalog/product-media/with-id/:id
	//   Sayfalama    : /catalog/product-media/pagination
	//   Toplu işlem  : /catalog/product-media/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-media/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-media/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ProductMediaForm)
	// 	// if created, ok := event.Data.(structures.ProductMediaForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-media", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-media/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-media/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ProductMediaForm)
	// 	// if created, ok := event.Data.([]structures.ProductMediaForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-media/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ProductMediaBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ProductMediaBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/product-media/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-media/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-media/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-media/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterProductReviewsEventHooks, ProductReviews tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterProductReviewsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/product-reviews
	//   Tekil işlem  : /catalog/product-reviews/with-id/:id
	//   Sayfalama    : /catalog/product-reviews/pagination
	//   Toplu işlem  : /catalog/product-reviews/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-reviews/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-reviews/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ProductReviewsForm)
	// 	// if created, ok := event.Data.(structures.ProductReviewsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-reviews", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-reviews/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-reviews/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ProductReviewsForm)
	// 	// if created, ok := event.Data.([]structures.ProductReviewsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-reviews/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ProductReviewsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ProductReviewsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/product-reviews/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-reviews/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-reviews/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-reviews/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterProductTagsEventHooks, ProductTags tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterProductTagsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/product-tags
	//   Tekil işlem  : /catalog/product-tags/with-id/:productId/:tagId
	//   Sayfalama    : /catalog/product-tags/pagination
	//   Toplu işlem  : /catalog/product-tags/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-tags/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags/with-id/:productId/:tagId", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags/with-id/:productId/:tagId", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-tags/with-id/:productId/:tagId", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ProductTagsForm)
	// 	// if created, ok := event.Data.(structures.ProductTagsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-tags", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags/with-id/:productId/:tagId", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags/with-id/:productId/:tagId", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-tags/with-id/:productId/:tagId", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags/with-id/:productId/:tagId", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags/with-id/:productId/:tagId", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-tags/with-id/:productId/:tagId", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ProductTagsForm)
	// 	// if created, ok := event.Data.([]structures.ProductTagsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-tags/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ProductTagsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ProductTagsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/product-tags/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-tags/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-tags/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-tags/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterProductVariantsEventHooks, ProductVariants tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterProductVariantsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/product-variants
	//   Tekil işlem  : /catalog/product-variants/with-id/:id
	//   Sayfalama    : /catalog/product-variants/pagination
	//   Toplu işlem  : /catalog/product-variants/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-variants/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-variants/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ProductVariantsForm)
	// 	// if created, ok := event.Data.(structures.ProductVariantsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-variants", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-variants/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-variants/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ProductVariantsForm)
	// 	// if created, ok := event.Data.([]structures.ProductVariantsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/product-variants/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ProductVariantsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ProductVariantsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/product-variants/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/product-variants/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/product-variants/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/product-variants/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterProductsEventHooks, Products tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterProductsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/products
	//   Tekil işlem  : /catalog/products/with-id/:id
	//   Sayfalama    : /catalog/products/pagination
	//   Toplu işlem  : /catalog/products/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/products/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/products/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ProductsForm)
	// 	// if created, ok := event.Data.(structures.ProductsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/products", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/products/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/products/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ProductsForm)
	// 	// if created, ok := event.Data.([]structures.ProductsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/products/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ProductsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ProductsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/products/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/products/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/products/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/products/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterTagsEventHooks, Tags tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllCatalogEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterTagsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /catalog/tags
	//   Tekil işlem  : /catalog/tags/with-id/:id
	//   Sayfalama    : /catalog/tags/pagination
	//   Toplu işlem  : /catalog/tags/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/tags/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/tags/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.TagsForm)
	// 	// if created, ok := event.Data.(structures.TagsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/tags", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/tags/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/tags/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.TagsForm)
	// 	// if created, ok := event.Data.([]structures.TagsForm); ok { ... }
	// })
	// em.OnEvent("/catalog/tags/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.TagsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.TagsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/catalog/tags/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/catalog/tags/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/catalog/tags/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/catalog/tags/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}
