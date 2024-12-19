package order_management

import (
	"context"
	"fmt"
	"log"
	"rental-api/Domain/constants"
	domestic_models "rental-api/Domain/models/domestic"
	"rental-api/Domain/pagination"
	"time"
)

func (c *OrderStatusChecker) archiveOldOrders(ctx context.Context) error {
	monthAgo := time.Now().Add(-30 * 24 * time.Hour)
	orders, err := c.getOrdersByDateRange(
		monthAgo.Add(-6*30*24*time.Hour),
		monthAgo,
		"",
		"END_BETWEEN",
	)
	if err != nil {
		return err
	}

	for _, order := range orders.Items {
		if c.shouldArchiveOrder(order) {
			order.Status = string(constants.OrderStatusArchived)
			if err := c.orderCommandRepository.UpdateCarOrder(ctx, &order); err != nil {
				log.Printf("Error archiving order %s: %v", order.Id, err)
			}
		}
	}
	return nil
}

func (c *OrderStatusChecker) getOrdersByDateRange(start, end time.Time, status string, filterType string) (*pagination.PaginatedResult[domestic_models.CarOrderModel], error) {
	return c.orderQueryRepository.GetCarOrders(
		nil,
		nil,
		start.Format(time.RFC3339),
		end.Format(time.RFC3339),
		"",
		"",
		status,
		filterType,
	)
}

func (c *OrderStatusChecker) notifyAndUpdateOrder(ctx context.Context, order domestic_models.CarOrderModel, newStatus constants.CarOrderStatus) error {
	carOffer, err := c.offerQueryRepository.GetCarOfferByID(order.CarOfferId)
	if err != nil || carOffer == nil {
		return fmt.Errorf("error fetching car offer: %v", err)
	}

	notification := c.prepareNotification(order, *carOffer, newStatus)
	if err := c.microserviceConnector.SendEmail(notification); err != nil {
		return fmt.Errorf("error sending notification: %v", err)
	}

	order.Status = string(newStatus)
	return c.orderCommandRepository.UpdateCarOrder(ctx, &order)
}
