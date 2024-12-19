package constants

type CarOrderStatus string

const (
	OrderStatusPending   CarOrderStatus = "PENDING"
	OrderStatusPreparing CarOrderStatus = "PREPARING"
	OrderStatusReady     CarOrderStatus = "READY"
	OrderStatusActive    CarOrderStatus = "ACTIVE"
	OrderStatusCompleted CarOrderStatus = "COMPLETED"
	OrderStatusCancelled CarOrderStatus = "CANCELLED"
	OrderStatusArchived  CarOrderStatus = "ARCHIVED"
)
