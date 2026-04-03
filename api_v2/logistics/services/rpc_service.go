package logistics_api_service

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

// LowStockCount, veritabanındaki `logistics.rpc__low_stock_count` fonksiyonunu çağırır.
func (s *RPCService) LowStockCount(pWarehouseId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM logistics.rpc__low_stock_count(?)"
	row := s.DB.Raw(query, pWarehouseId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// WarehouseUtilization, veritabanındaki `logistics.rpc__warehouse_utilization` fonksiyonunu çağırır.
func (s *RPCService) WarehouseUtilization(pWarehouseId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM logistics.rpc__warehouse_utilization(?)"
	row := s.DB.Raw(query, pWarehouseId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}
