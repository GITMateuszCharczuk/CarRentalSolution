package order_management

import (
	"context"
	"log"
	offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	order_repository "rental-api/Domain/repository_interfaces/car_order_repository"
	services "rental-api/Domain/service_interfaces"
	"time"
)

type OrderStatusChecker struct {
	orderQueryRepository   order_repository.CarOrderQueryRepository
	orderCommandRepository order_repository.CarOrderCommandRepository
	offerQueryRepository   offer_repository.CarOfferQueryRepository
	microserviceConnector  services.MicroserviceConnector
}

func NewOrderStatusChecker(
	orderQueryRepo order_repository.CarOrderQueryRepository,
	orderCommandRepo order_repository.CarOrderCommandRepository,
	offerQueryRepo offer_repository.CarOfferQueryRepository,
	microserviceConnector services.MicroserviceConnector,
) *OrderStatusChecker {
	return &OrderStatusChecker{
		orderQueryRepository:   orderQueryRepo,
		orderCommandRepository: orderCommandRepo,
		offerQueryRepository:   offerQueryRepo,
		microserviceConnector:  microserviceConnector,
	}
}

func (c *OrderStatusChecker) StartPeriodicCheck(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				if err := c.CheckOrderStatuses(ctx); err != nil {
					log.Printf("Error checking order statuses: %v", err)
				}
			}
		}
	}()
}

func (c *OrderStatusChecker) CheckOrderStatuses(ctx context.Context) error {
	if err := c.handlePendingOrders(ctx); err != nil {
		log.Printf("Error handling pending orders: %v", err)
	}

	if err := c.handleActiveOrders(ctx); err != nil {
		log.Printf("Error handling active orders: %v", err)
	}

	if err := c.archiveOldOrders(ctx); err != nil {
		log.Printf("Error archiving old orders: %v", err)
	}

	return nil
}
