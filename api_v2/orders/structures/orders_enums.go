package orders_api_structure

// Bu dosya, veritabanındaki ENUM tiplerinden otomatik olarak üretilmiştir.

// OrdersOrderStatus represents the 'order_status' enum type from the 'orders' schema.
type OrdersOrderStatus string

const (
	OrdersOrderStatusPending    OrdersOrderStatus = "pending"
	OrdersOrderStatusConfirmed  OrdersOrderStatus = "confirmed"
	OrdersOrderStatusProcessing OrdersOrderStatus = "processing"
	OrdersOrderStatusShipped    OrdersOrderStatus = "shipped"
	OrdersOrderStatusDelivered  OrdersOrderStatus = "delivered"
	OrdersOrderStatusCancelled  OrdersOrderStatus = "cancelled"
	OrdersOrderStatusRefunded   OrdersOrderStatus = "refunded"
)

// OrdersPaymentMethod represents the 'payment_method' enum type from the 'orders' schema.
type OrdersPaymentMethod string

const (
	OrdersPaymentMethodCreditCard   OrdersPaymentMethod = "credit_card"
	OrdersPaymentMethodDebitCard    OrdersPaymentMethod = "debit_card"
	OrdersPaymentMethodBankTransfer OrdersPaymentMethod = "bank_transfer"
	OrdersPaymentMethodWallet       OrdersPaymentMethod = "wallet"
	OrdersPaymentMethodCrypto       OrdersPaymentMethod = "crypto"
)

// OrdersPaymentStatus represents the 'payment_status' enum type from the 'orders' schema.
type OrdersPaymentStatus string

const (
	OrdersPaymentStatusPending    OrdersPaymentStatus = "pending"
	OrdersPaymentStatusAuthorized OrdersPaymentStatus = "authorized"
	OrdersPaymentStatusCaptured   OrdersPaymentStatus = "captured"
	OrdersPaymentStatusFailed     OrdersPaymentStatus = "failed"
	OrdersPaymentStatusRefunded   OrdersPaymentStatus = "refunded"
)
