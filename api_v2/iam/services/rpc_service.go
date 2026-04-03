package iam_api_service

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

// CountActiveUsers, veritabanındaki `iam.rpc__count_active_users` fonksiyonunu çağırır.
func (s *RPCService) CountActiveUsers() (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM iam.rpc__count_active_users()"
	row := s.DB.Raw(query).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// UserPermissions, veritabanındaki `iam.rpc__user_permissions` fonksiyonunu çağırır.
func (s *RPCService) UserPermissions(pUserId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM iam.rpc__user_permissions(?)"
	row := s.DB.Raw(query, pUserId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}

// UsersByOrganization, veritabanındaki `iam.rpc__users_by_organization` fonksiyonunu çağırır.
func (s *RPCService) UsersByOrganization(pOrgId types.URID) (interface{}, error) {
	var result interface{} // Genel bir sonuç tipi

	query := "SELECT * FROM iam.rpc__users_by_organization(?)"
	row := s.DB.Raw(query, pOrgId).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}

	if result == nil {
		return []map[string]interface{}{}, nil
	}

	return result, nil
}
