package analytics_api_service

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

// DashboardCount, veritabanındaki `analytics.rpc__dashboard_count` fonksiyonunu çağırır.
func (s *RPCService) DashboardCount() (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM analytics.rpc__dashboard_count()"
	row := s.DB.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// EventCountBySeverity, veritabanındaki `analytics.rpc__event_count_by_severity` fonksiyonunu çağırır.
func (s *RPCService) EventCountBySeverity(pSeverity string) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM analytics.rpc__event_count_by_severity(?)"
	row := s.DB.Raw(query, pSeverity).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// UnreadNotificationCount, veritabanındaki `analytics.rpc__unread_notification_count` fonksiyonunu çağırır.
func (s *RPCService) UnreadNotificationCount(pUserId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM analytics.rpc__unread_notification_count(?)"
	row := s.DB.Raw(query, pUserId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}
