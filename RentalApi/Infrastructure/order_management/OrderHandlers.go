package order_management

import (
	"context"
	"log"
	"rental-api/Domain/constants"
	domestic_models "rental-api/Domain/models/domestic"
	"time"
)

func (c *OrderStatusChecker) handleOverdueOrder(ctx context.Context, order domestic_models.CarOrderModel) error {
	endDate, _ := time.Parse(time.RFC3339, order.EndDate)
	if time.Since(endDate) > 24*time.Hour {
		return c.notifyAndUpdateOrder(ctx, order, constants.OrderStatusCompleted)
	}
	return nil
}

func (c *OrderStatusChecker) handleActiveOrders(ctx context.Context) error {
	yesterday := time.Now().Add(-24 * time.Hour)
	orders, err := c.getOrdersByDateRange(
		yesterday.Add(-7*24*time.Hour), // Get orders from last week
		yesterday,
		string(constants.OrderStatusActive),
		"END_BETWEEN",
	)
	if err != nil {
		return err
	}

	for _, order := range orders.Items {
		if err := c.handleOverdueOrder(ctx, order); err != nil {
			log.Printf("Error handling overdue order %s: %v", order.Id, err)
		}
	}
	return nil
}

func (c *OrderStatusChecker) handlePendingOrders(ctx context.Context) error {
	tomorrow := time.Now().Add(24 * time.Hour)
	orders, err := c.getOrdersByDateRange(
		tomorrow,
		tomorrow.Add(24*time.Hour),
		string(constants.OrderStatusPending),
		"START_BETWEEN",
	)
	if err != nil {
		return err
	}

	for _, order := range orders.Items {
		if err := c.notifyAndUpdateOrder(ctx, order, constants.OrderStatusPreparing); err != nil {
			log.Printf("Error processing pending order %s: %v", order.Id, err)
		}
	}
	return nil
}
