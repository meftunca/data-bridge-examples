// Endpoint Registration Template - Enterprise-Grade Route Management
//
// This template generates comprehensive HTTP endpoint registration for generated resources.
// It implements enterprise patterns including:
//
// • RESTful API conventions with consistent URL patterns
// • Dependency injection for controllers and services
// • Event-driven architecture integration for request monitoring
// • Route grouping for logical resource organization
// • Comprehensive CRUD operations with bulk processing support
// • Type-safe parameter binding and validation
// • RBAC endpoint-based permission middleware (when RBACEnabled=true)
// • SSE real-time subscription endpoint (when SSEEnabled=true)
//
// Generated Route Patterns:
// • GET    /{resource}/pagination     - Paginated listing with filtering
// • GET    /{resource}/with-id/{id}   - Single resource retrieval
// • POST   /{resource}                - Create new resource
// • POST   /{resource}/bulk           - Bulk resource creation
// • PUT    /{resource}/with-id/{id}   - Update existing resource
// • PUT    /{resource}/bulk           - Bulk resource updates
// • DELETE /{resource}/with-id/{id}   - Delete single resource
// • DELETE /{resource}/bulk           - Bulk resource deletion
// • GET    /{resource}/events         - SSE real-time subscription (SSEEnabled=true)
//
// Permission Key Format (RBACEnabled=true):
//
//	{schema}.{resource}.{action}
//	e.g.: core.users.read, core.users.write, core.users.delete
package catalog

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Dinamik olarak ilgili paketleri import et
	controller "data-bridge-examples/api_v2/catalog/controllers"
	services "data-bridge-examples/api_v2/catalog/services"

	"github.com/maple-tech/baseline/events" // EventManager için

	"backend-generator/apiv2/sse"
)

// Run initializes and registers all HTTP endpoints for the generated API schema.
//
// This function sets up the complete routing infrastructure with:
// • Dependency injection for all controllers and services
// • Event manager integration for request lifecycle monitoring
// • Database connection propagation to service layer
// • Route group organization for clean URL structure
// • RESTful endpoint registration following HTTP conventions
//
// The registration process creates route groups for each resource with consistent
// patterns and comprehensive CRUD operations. All endpoints include proper error
// handling, validation, and observability through the event system.
//
// Parameters:
//   - app: Fiber router instance for route registration
//   - db: GORM database connection for service layer operations
//   - eventManager: Event manager for request monitoring and observability
//
// Generated Structure:
//
//	Each resource gets its own route group with standardized endpoints
//	supporting both single and bulk operations for optimal performance.
func Run(app fiber.Router, db *gorm.DB, eventManager *events.EventManager) {

	// --- brands için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		BrandsController := &controller.BrandsController{
			Svc: &services.BrandsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		BrandsGroup := app.Group("/brands")

		// Rotaları tanımla
		BrandsGroup.Get("/pagination", BrandsController.GetBrandsWithPagination)

		BrandsGroup.Get("/with-id/:id", BrandsController.GetBrandsById)
		BrandsGroup.Post("/", BrandsController.CreateBrands)
		BrandsGroup.Post("/bulk", BrandsController.CreateBrandsMultiple)
		BrandsGroup.Put("/with-id/:id", BrandsController.UpdateBrands)
		BrandsGroup.Delete("/with-id/:id", BrandsController.DeleteBrands)
		BrandsGroup.Delete("/bulk", BrandsController.DeleteBrandsMultiple)
		BrandsGroup.Put("/bulk", BrandsController.UpdateBrandsMultiple)

		BrandsGroup.Get("/events", sse.StreamResource("catalog.brands"))
	}

	// --- categories için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		CategoriesController := &controller.CategoriesController{
			Svc: &services.CategoriesService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		CategoriesGroup := app.Group("/categories")

		// Rotaları tanımla
		CategoriesGroup.Get("/pagination", CategoriesController.GetCategoriesWithPagination)

		CategoriesGroup.Get("/with-id/:id", CategoriesController.GetCategoriesById)
		CategoriesGroup.Post("/", CategoriesController.CreateCategories)
		CategoriesGroup.Post("/bulk", CategoriesController.CreateCategoriesMultiple)
		CategoriesGroup.Put("/with-id/:id", CategoriesController.UpdateCategories)
		CategoriesGroup.Delete("/with-id/:id", CategoriesController.DeleteCategories)
		CategoriesGroup.Delete("/bulk", CategoriesController.DeleteCategoriesMultiple)
		CategoriesGroup.Put("/bulk", CategoriesController.UpdateCategoriesMultiple)

		CategoriesGroup.Get("/events", sse.StreamResource("catalog.categories"))
	}

	// --- collections için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		CollectionsController := &controller.CollectionsController{
			Svc: &services.CollectionsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		CollectionsGroup := app.Group("/collections")

		// Rotaları tanımla
		CollectionsGroup.Get("/pagination", CollectionsController.GetCollectionsWithPagination)

		CollectionsGroup.Get("/with-id/:id", CollectionsController.GetCollectionsById)
		CollectionsGroup.Post("/", CollectionsController.CreateCollections)
		CollectionsGroup.Post("/bulk", CollectionsController.CreateCollectionsMultiple)
		CollectionsGroup.Put("/with-id/:id", CollectionsController.UpdateCollections)
		CollectionsGroup.Delete("/with-id/:id", CollectionsController.DeleteCollections)
		CollectionsGroup.Delete("/bulk", CollectionsController.DeleteCollectionsMultiple)
		CollectionsGroup.Put("/bulk", CollectionsController.UpdateCollectionsMultiple)

		CollectionsGroup.Get("/events", sse.StreamResource("catalog.collections"))
	}

	// --- products için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ProductsController := &controller.ProductsController{
			Svc: &services.ProductsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ProductsGroup := app.Group("/products")

		// Rotaları tanımla
		ProductsGroup.Get("/pagination", ProductsController.GetProductsWithPagination)

		ProductsGroup.Get("/with-id/:id", ProductsController.GetProductsById)
		ProductsGroup.Post("/", ProductsController.CreateProducts)
		ProductsGroup.Post("/bulk", ProductsController.CreateProductsMultiple)
		ProductsGroup.Put("/with-id/:id", ProductsController.UpdateProducts)
		ProductsGroup.Delete("/with-id/:id", ProductsController.DeleteProducts)
		ProductsGroup.Delete("/bulk", ProductsController.DeleteProductsMultiple)
		ProductsGroup.Put("/bulk", ProductsController.UpdateProductsMultiple)

		ProductsGroup.Get("/events", sse.StreamResource("catalog.products"))
	}

	// --- tags için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		TagsController := &controller.TagsController{
			Svc: &services.TagsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		TagsGroup := app.Group("/tags")

		// Rotaları tanımla
		TagsGroup.Get("/pagination", TagsController.GetTagsWithPagination)

		TagsGroup.Get("/with-id/:id", TagsController.GetTagsById)
		TagsGroup.Post("/", TagsController.CreateTags)
		TagsGroup.Post("/bulk", TagsController.CreateTagsMultiple)
		TagsGroup.Put("/with-id/:id", TagsController.UpdateTags)
		TagsGroup.Delete("/with-id/:id", TagsController.DeleteTags)
		TagsGroup.Delete("/bulk", TagsController.DeleteTagsMultiple)
		TagsGroup.Put("/bulk", TagsController.UpdateTagsMultiple)

		TagsGroup.Get("/events", sse.StreamResource("catalog.tags"))
	}

	// --- collection_products için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		CollectionProductsController := &controller.CollectionProductsController{
			Svc: &services.CollectionProductsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		CollectionProductsGroup := app.Group("/collection-products")

		// Rotaları tanımla
		CollectionProductsGroup.Get("/pagination", CollectionProductsController.GetCollectionProductsWithPagination)

		CollectionProductsGroup.Get("/with-id/:id", CollectionProductsController.GetCollectionProductsById)
		CollectionProductsGroup.Post("/", CollectionProductsController.CreateCollectionProducts)
		CollectionProductsGroup.Post("/bulk", CollectionProductsController.CreateCollectionProductsMultiple)
		CollectionProductsGroup.Put("/with-id/:id", CollectionProductsController.UpdateCollectionProducts)
		CollectionProductsGroup.Delete("/with-id/:id", CollectionProductsController.DeleteCollectionProducts)
		CollectionProductsGroup.Delete("/bulk", CollectionProductsController.DeleteCollectionProductsMultiple)
		CollectionProductsGroup.Put("/bulk", CollectionProductsController.UpdateCollectionProductsMultiple)

		CollectionProductsGroup.Get("/events", sse.StreamResource("catalog.collection-products"))
	}

	// --- price_history için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		PriceHistoryController := &controller.PriceHistoryController{
			Svc: &services.PriceHistoryService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		PriceHistoryGroup := app.Group("/price-history")

		// Rotaları tanımla
		PriceHistoryGroup.Get("/pagination", PriceHistoryController.GetPriceHistoryWithPagination)

		PriceHistoryGroup.Get("/with-id/:id", PriceHistoryController.GetPriceHistoryById)
		PriceHistoryGroup.Post("/", PriceHistoryController.CreatePriceHistory)
		PriceHistoryGroup.Post("/bulk", PriceHistoryController.CreatePriceHistoryMultiple)
		PriceHistoryGroup.Put("/with-id/:id", PriceHistoryController.UpdatePriceHistory)
		PriceHistoryGroup.Delete("/with-id/:id", PriceHistoryController.DeletePriceHistory)
		PriceHistoryGroup.Delete("/bulk", PriceHistoryController.DeletePriceHistoryMultiple)
		PriceHistoryGroup.Put("/bulk", PriceHistoryController.UpdatePriceHistoryMultiple)

		PriceHistoryGroup.Get("/events", sse.StreamResource("catalog.price-history"))
	}

	// --- product_media için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ProductMediaController := &controller.ProductMediaController{
			Svc: &services.ProductMediaService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ProductMediaGroup := app.Group("/product-media")

		// Rotaları tanımla
		ProductMediaGroup.Get("/pagination", ProductMediaController.GetProductMediaWithPagination)

		ProductMediaGroup.Get("/with-id/:id", ProductMediaController.GetProductMediaById)
		ProductMediaGroup.Post("/", ProductMediaController.CreateProductMedia)
		ProductMediaGroup.Post("/bulk", ProductMediaController.CreateProductMediaMultiple)
		ProductMediaGroup.Put("/with-id/:id", ProductMediaController.UpdateProductMedia)
		ProductMediaGroup.Delete("/with-id/:id", ProductMediaController.DeleteProductMedia)
		ProductMediaGroup.Delete("/bulk", ProductMediaController.DeleteProductMediaMultiple)
		ProductMediaGroup.Put("/bulk", ProductMediaController.UpdateProductMediaMultiple)

		ProductMediaGroup.Get("/events", sse.StreamResource("catalog.product-media"))
	}

	// --- product_reviews için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ProductReviewsController := &controller.ProductReviewsController{
			Svc: &services.ProductReviewsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ProductReviewsGroup := app.Group("/product-reviews")

		// Rotaları tanımla
		ProductReviewsGroup.Get("/pagination", ProductReviewsController.GetProductReviewsWithPagination)

		ProductReviewsGroup.Get("/with-id/:id", ProductReviewsController.GetProductReviewsById)
		ProductReviewsGroup.Post("/", ProductReviewsController.CreateProductReviews)
		ProductReviewsGroup.Post("/bulk", ProductReviewsController.CreateProductReviewsMultiple)
		ProductReviewsGroup.Put("/with-id/:id", ProductReviewsController.UpdateProductReviews)
		ProductReviewsGroup.Delete("/with-id/:id", ProductReviewsController.DeleteProductReviews)
		ProductReviewsGroup.Delete("/bulk", ProductReviewsController.DeleteProductReviewsMultiple)
		ProductReviewsGroup.Put("/bulk", ProductReviewsController.UpdateProductReviewsMultiple)

		ProductReviewsGroup.Get("/events", sse.StreamResource("catalog.product-reviews"))
	}

	// --- product_tags için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ProductTagsController := &controller.ProductTagsController{
			Svc: &services.ProductTagsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ProductTagsGroup := app.Group("/product-tags")

		// Rotaları tanımla
		ProductTagsGroup.Get("/pagination", ProductTagsController.GetProductTagsWithPagination)

		ProductTagsGroup.Get("/with-id/:productId/:tagId", ProductTagsController.GetProductTagsById)
		ProductTagsGroup.Post("/", ProductTagsController.CreateProductTags)
		ProductTagsGroup.Post("/bulk", ProductTagsController.CreateProductTagsMultiple)
		ProductTagsGroup.Put("/with-id/:productId/:tagId", ProductTagsController.UpdateProductTags)
		ProductTagsGroup.Delete("/with-id/:productId/:tagId", ProductTagsController.DeleteProductTags)
		ProductTagsGroup.Delete("/bulk", ProductTagsController.DeleteProductTagsMultiple)
		ProductTagsGroup.Put("/bulk", ProductTagsController.UpdateProductTagsMultiple)

		ProductTagsGroup.Get("/events", sse.StreamResource("catalog.product-tags"))
	}

	// --- product_variants için Rota Grubu ---
	{
		// Controller ve Service katmanlarını ilklendir (initialize)
		ProductVariantsController := &controller.ProductVariantsController{
			Svc: &services.ProductVariantsService{DB: db},
			EM:  eventManager,
		}

		// Fiber rota grubunu oluştur
		ProductVariantsGroup := app.Group("/product-variants")

		// Rotaları tanımla
		ProductVariantsGroup.Get("/pagination", ProductVariantsController.GetProductVariantsWithPagination)

		ProductVariantsGroup.Get("/with-id/:id", ProductVariantsController.GetProductVariantsById)
		ProductVariantsGroup.Post("/", ProductVariantsController.CreateProductVariants)
		ProductVariantsGroup.Post("/bulk", ProductVariantsController.CreateProductVariantsMultiple)
		ProductVariantsGroup.Put("/with-id/:id", ProductVariantsController.UpdateProductVariants)
		ProductVariantsGroup.Delete("/with-id/:id", ProductVariantsController.DeleteProductVariants)
		ProductVariantsGroup.Delete("/bulk", ProductVariantsController.DeleteProductVariantsMultiple)
		ProductVariantsGroup.Put("/bulk", ProductVariantsController.UpdateProductVariantsMultiple)

		ProductVariantsGroup.Get("/events", sse.StreamResource("catalog.product-variants"))
	}

	rpcSvc := &services.RPCService{DB: db}
	rpcController := &controller.RPCController{Svc: rpcSvc}
	rpcGroup := app.Group("/rpc")
	{
		rpcGroup.Post("/avg-product-rating", rpcController.AvgProductRating)
		rpcGroup.Post("/count-active-products", rpcController.CountActiveProducts)
		rpcGroup.Post("/products-by-category", rpcController.ProductsByCategory)
	}
}
