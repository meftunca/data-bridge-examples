package iam_handlers

import (
	"github.com/maple-tech/baseline/events"
)

// RegisterAllIamEventHandlers, bu şemadaki tüm tablo hook'larını
// tek seferde kaydeder. Varsayılan olarak tüm tablolar kapalıdır.
// İstediğiniz tablonun satırındaki "//" işaretini kaldırarak aktif edin.
//
// Bu fonksiyonu main.go veya uygulama setup kodunuzda çağırın:
//
//	em := events.NewEventManager()
//	RegisterAllIamEventHandlers(em)
//	api.Setup(app, db, em)
func RegisterAllIamEventHandlers(em *events.EventManager) {
	// RegisterApiKeysEventHooks(em)
	// RegisterInvitationsEventHooks(em)
	// RegisterOrganizationsEventHooks(em)
	// RegisterPermissionsEventHooks(em)
	// RegisterRolePermissionsEventHooks(em)
	// RegisterRolesEventHooks(em)
	// RegisterSessionsEventHooks(em)
	// RegisterTeamMembersEventHooks(em)
	// RegisterTeamsEventHooks(em)
	// RegisterUserRolesEventHooks(em)
	// RegisterUsersEventHooks(em)
}

// RegisterApiKeysEventHooks, ApiKeys tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterApiKeysEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/api-keys
	//   Tekil işlem  : /iam/api-keys/with-id/:id
	//   Sayfalama    : /iam/api-keys/pagination
	//   Toplu işlem  : /iam/api-keys/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/api-keys/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/api-keys/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.ApiKeysForm)
	// 	// if created, ok := event.Data.(structures.ApiKeysForm); ok { ... }
	// })
	// em.OnEvent("/iam/api-keys", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/api-keys/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/api-keys/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.ApiKeysForm)
	// 	// if created, ok := event.Data.([]structures.ApiKeysForm); ok { ... }
	// })
	// em.OnEvent("/iam/api-keys/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.ApiKeysBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.ApiKeysBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/api-keys/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/api-keys/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/api-keys/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/api-keys/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterInvitationsEventHooks, Invitations tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterInvitationsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/invitations
	//   Tekil işlem  : /iam/invitations/with-id/:id
	//   Sayfalama    : /iam/invitations/pagination
	//   Toplu işlem  : /iam/invitations/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/invitations/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/invitations/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.InvitationsForm)
	// 	// if created, ok := event.Data.(structures.InvitationsForm); ok { ... }
	// })
	// em.OnEvent("/iam/invitations", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/invitations/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/invitations/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.InvitationsForm)
	// 	// if created, ok := event.Data.([]structures.InvitationsForm); ok { ... }
	// })
	// em.OnEvent("/iam/invitations/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.InvitationsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.InvitationsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/invitations/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/invitations/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/invitations/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/invitations/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterOrganizationsEventHooks, Organizations tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterOrganizationsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/organizations
	//   Tekil işlem  : /iam/organizations/with-id/:id
	//   Sayfalama    : /iam/organizations/pagination
	//   Toplu işlem  : /iam/organizations/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/organizations/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/organizations/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.OrganizationsForm)
	// 	// if created, ok := event.Data.(structures.OrganizationsForm); ok { ... }
	// })
	// em.OnEvent("/iam/organizations", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/organizations/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/organizations/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.OrganizationsForm)
	// 	// if created, ok := event.Data.([]structures.OrganizationsForm); ok { ... }
	// })
	// em.OnEvent("/iam/organizations/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.OrganizationsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.OrganizationsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/organizations/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/organizations/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/organizations/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/organizations/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterPermissionsEventHooks, Permissions tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterPermissionsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/permissions
	//   Tekil işlem  : /iam/permissions/with-id/:id
	//   Sayfalama    : /iam/permissions/pagination
	//   Toplu işlem  : /iam/permissions/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/permissions/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/permissions/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.PermissionsForm)
	// 	// if created, ok := event.Data.(structures.PermissionsForm); ok { ... }
	// })
	// em.OnEvent("/iam/permissions", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/permissions/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/permissions/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.PermissionsForm)
	// 	// if created, ok := event.Data.([]structures.PermissionsForm); ok { ... }
	// })
	// em.OnEvent("/iam/permissions/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.PermissionsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.PermissionsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/permissions/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/permissions/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/permissions/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/permissions/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterRolePermissionsEventHooks, RolePermissions tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterRolePermissionsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/role-permissions
	//   Tekil işlem  : /iam/role-permissions/with-id/:id
	//   Sayfalama    : /iam/role-permissions/pagination
	//   Toplu işlem  : /iam/role-permissions/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/role-permissions/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/role-permissions/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.RolePermissionsForm)
	// 	// if created, ok := event.Data.(structures.RolePermissionsForm); ok { ... }
	// })
	// em.OnEvent("/iam/role-permissions", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/role-permissions/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/role-permissions/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.RolePermissionsForm)
	// 	// if created, ok := event.Data.([]structures.RolePermissionsForm); ok { ... }
	// })
	// em.OnEvent("/iam/role-permissions/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.RolePermissionsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.RolePermissionsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/role-permissions/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/role-permissions/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/role-permissions/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/role-permissions/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterRolesEventHooks, Roles tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterRolesEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/roles
	//   Tekil işlem  : /iam/roles/with-id/:id
	//   Sayfalama    : /iam/roles/pagination
	//   Toplu işlem  : /iam/roles/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/roles/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/roles/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.RolesForm)
	// 	// if created, ok := event.Data.(structures.RolesForm); ok { ... }
	// })
	// em.OnEvent("/iam/roles", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/roles/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/roles/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.RolesForm)
	// 	// if created, ok := event.Data.([]structures.RolesForm); ok { ... }
	// })
	// em.OnEvent("/iam/roles/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.RolesBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.RolesBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/roles/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/roles/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/roles/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/roles/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterSessionsEventHooks, Sessions tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterSessionsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/sessions
	//   Tekil işlem  : /iam/sessions/with-id/:id
	//   Sayfalama    : /iam/sessions/pagination
	//   Toplu işlem  : /iam/sessions/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/sessions/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/sessions/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.SessionsForm)
	// 	// if created, ok := event.Data.(structures.SessionsForm); ok { ... }
	// })
	// em.OnEvent("/iam/sessions", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/sessions/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/sessions/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.SessionsForm)
	// 	// if created, ok := event.Data.([]structures.SessionsForm); ok { ... }
	// })
	// em.OnEvent("/iam/sessions/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.SessionsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.SessionsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/sessions/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/sessions/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/sessions/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/sessions/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterTeamMembersEventHooks, TeamMembers tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterTeamMembersEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/team-members
	//   Tekil işlem  : /iam/team-members/with-id/:id
	//   Sayfalama    : /iam/team-members/pagination
	//   Toplu işlem  : /iam/team-members/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/team-members/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/team-members/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.TeamMembersForm)
	// 	// if created, ok := event.Data.(structures.TeamMembersForm); ok { ... }
	// })
	// em.OnEvent("/iam/team-members", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/team-members/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/team-members/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.TeamMembersForm)
	// 	// if created, ok := event.Data.([]structures.TeamMembersForm); ok { ... }
	// })
	// em.OnEvent("/iam/team-members/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.TeamMembersBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.TeamMembersBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/team-members/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/team-members/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/team-members/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/team-members/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterTeamsEventHooks, Teams tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterTeamsEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/teams
	//   Tekil işlem  : /iam/teams/with-id/:id
	//   Sayfalama    : /iam/teams/pagination
	//   Toplu işlem  : /iam/teams/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/teams/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/teams/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.TeamsForm)
	// 	// if created, ok := event.Data.(structures.TeamsForm); ok { ... }
	// })
	// em.OnEvent("/iam/teams", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/teams/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/teams/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.TeamsForm)
	// 	// if created, ok := event.Data.([]structures.TeamsForm); ok { ... }
	// })
	// em.OnEvent("/iam/teams/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.TeamsBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.TeamsBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/teams/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/teams/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/teams/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/teams/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterUserRolesEventHooks, UserRoles tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterUserRolesEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/user-roles
	//   Tekil işlem  : /iam/user-roles/with-id/:id
	//   Sayfalama    : /iam/user-roles/pagination
	//   Toplu işlem  : /iam/user-roles/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/user-roles/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/user-roles/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.UserRolesForm)
	// 	// if created, ok := event.Data.(structures.UserRolesForm); ok { ... }
	// })
	// em.OnEvent("/iam/user-roles", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/user-roles/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/user-roles/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.UserRolesForm)
	// 	// if created, ok := event.Data.([]structures.UserRolesForm); ok { ... }
	// })
	// em.OnEvent("/iam/user-roles/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.UserRolesBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.UserRolesBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/user-roles/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/user-roles/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/user-roles/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/user-roles/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}

// RegisterUsersEventHooks, Users tablosu için event hook'larını kaydeder.
// Aktif etmek için:
//  1. RegisterAllIamEventHandlers içinde bu fonksiyonun çağrısını uncomment edin.
//  2. Aşağıdaki handler bloklarından istediklerinizin başındaki "//" 'yi kaldırın.
func RegisterUsersEventHooks(em *events.EventManager) {
	// Rotalar:
	//   Liste/oluştur : /iam/users
	//   Tekil işlem  : /iam/users/with-id/:id
	//   Sayfalama    : /iam/users/pagination
	//   Toplu işlem  : /iam/users/bulk

	// ── Sayfalama ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/users/pagination", events.PaginationQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users/pagination", events.PaginationQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/users/pagination", events.PaginationQueryError, func(event events.Event) {})

	// ── Tekil okuma ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/users/with-id/:id", events.FindQueryRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users/with-id/:id", events.FindQuerySuccess, func(event events.Event) {})
	// em.OnEvent("/iam/users/with-id/:id", events.FindQueryError, func(event events.Event) {})

	// ── Oluşturma ──────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/users", events.CreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users", events.CreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıt (structures.UsersForm)
	// 	// if created, ok := event.Data.(structures.UsersForm); ok { ... }
	// })
	// em.OnEvent("/iam/users", events.CreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Güncelleme ─────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/users/with-id/:id", events.UpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users/with-id/:id", events.UpdateSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/users/with-id/:id", events.UpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Silme ──────────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/users/with-id/:id", events.DeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users/with-id/:id", events.DeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/users/with-id/:id", events.DeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu oluşturma ────────────────────────────────────────────────────────
	// em.OnRequest("/iam/users/bulk", events.BatchCreationRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users/bulk", events.BatchCreationSuccess, func(event events.Event) {
	// 	// event.Data → oluşturulan kayıtlar ([]structures.UsersForm)
	// 	// if created, ok := event.Data.([]structures.UsersForm); ok { ... }
	// })
	// em.OnEvent("/iam/users/bulk", events.BatchCreationError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu güncelleme ───────────────────────────────────────────────────────
	// em.OnRequest("/iam/users/bulk", events.BatchUpdateRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users/bulk", events.BatchUpdateSuccess, func(event events.Event) {
	// 	// event.Data → işlenen batch payload ([]structures.UsersBatchUpdate)
	// 	// if batch, ok := event.Data.([]structures.UsersBatchUpdate); ok { ... }
	// })
	// em.OnEvent("/iam/users/bulk", events.BatchUpdateError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })

	// ── Toplu silme ────────────────────────────────────────────────────────────
	// em.OnRequest("/iam/users/bulk", events.BatchDeletionRequest, func(event events.Event) (bool, error) {
	// 	return true, nil
	// })
	// em.OnEvent("/iam/users/bulk", events.BatchDeletionSuccess, func(event events.Event) {})
	// em.OnEvent("/iam/users/bulk", events.BatchDeletionError, func(event events.Event) {
	// 	// event.Err → servis hatası
	// })
}
