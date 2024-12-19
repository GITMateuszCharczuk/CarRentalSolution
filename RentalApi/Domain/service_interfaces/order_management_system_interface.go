package service_interfaces

import "context"

type OrderManagementSystem interface {
	StartPeriodicCheck(ctx context.Context)
}
