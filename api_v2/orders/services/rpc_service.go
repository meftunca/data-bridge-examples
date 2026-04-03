package orders_api_service

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

// CustomerTotalSpent, veritabanındaki `orders.rpc__customer_total_spent` fonksiyonunu çağırır.
func (s *RPCService) CustomerTotalSpent(pCustomerId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM orders.rpc__customer_total_spent(?)"
	row := s.DB.Raw(query, pCustomerId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// OrdersByStatus, veritabanındaki `orders.rpc__orders_by_status` fonksiyonunu çağırır.
func (s *RPCService) OrdersByStatus(pStatus string) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM orders.rpc__orders_by_status(?)"
	row := s.DB.Raw(query, pStatus).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// TotalRevenue, veritabanındaki `orders.rpc__total_revenue` fonksiyonunu çağırır.
func (s *RPCService) TotalRevenue() (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM orders.rpc__total_revenue()"
	row := s.DB.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}
