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

func IsValidCarOrderStatus(status string) bool {
	switch CarOrderStatus(status) {
	case OrderStatusPending, OrderStatusPreparing, OrderStatusReady,
		OrderStatusActive, OrderStatusCompleted, OrderStatusCancelled,
		OrderStatusArchived:
		return true
	}
	return false
}

func ToCarOrderStatus(status string) CarOrderStatus {
	return CarOrderStatus(status)
}
