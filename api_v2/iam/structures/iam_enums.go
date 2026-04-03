package iam_api_structure

// Bu dosya, veritabanındaki ENUM tiplerinden otomatik olarak üretilmiştir.

// IamAuthProvider represents the 'auth_provider' enum type from the 'iam' schema.
type IamAuthProvider string

const (
	IamAuthProviderLocal   IamAuthProvider = "local"
	IamAuthProviderGoogle  IamAuthProvider = "google"
	IamAuthProviderGithub  IamAuthProvider = "github"
	IamAuthProviderAzureAd IamAuthProvider = "azure_ad"
	IamAuthProviderSaml    IamAuthProvider = "saml"
)

// IamUserStatus represents the 'user_status' enum type from the 'iam' schema.
type IamUserStatus string

const (
	IamUserStatusActive              IamUserStatus = "active"
	IamUserStatusInactive            IamUserStatus = "inactive"
	IamUserStatusSuspended           IamUserStatus = "suspended"
	IamUserStatusPendingVerification IamUserStatus = "pending_verification"
)
