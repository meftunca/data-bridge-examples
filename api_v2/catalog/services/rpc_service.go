package catalog_api_service

import (
	"time" // Zaman tipleri için

	"github.com/maple-tech/baseline/types"
	"gorm.io/gorm"
)

var (
	_ = types.ID(0)
	_ = types.URID("")
	_ = time.Time{}
)

// RPCService, RPC fonksiyonlarını çağıran metodları içerir.
type RPCService struct {
	DB *gorm.DB
}

// AvgProductRating, veritabanındaki `catalog.rpc__avg_product_rating` fonksiyonunu çağırır.
func (s *RPCService) AvgProductRating(pProductId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM catalog.rpc__avg_product_rating(?)"
	row := s.DB.Raw(query, pProductId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// CountActiveProducts, veritabanındaki `catalog.rpc__count_active_products` fonksiyonunu çağırır.
func (s *RPCService) CountActiveProducts() (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM catalog.rpc__count_active_products()"
	row := s.DB.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// ProductsByCategory, veritabanındaki `catalog.rpc__products_by_category` fonksiyonunu çağırır.
func (s *RPCService) ProductsByCategory(pCategoryId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM catalog.rpc__products_by_category(?)"
	row := s.DB.Raw(query, pCategoryId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}
