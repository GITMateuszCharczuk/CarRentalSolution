package order_management

import (
	"rental-api/Domain/constants"
	domestic_models "rental-api/Domain/models/domestic"
	"time"
)

func (c *OrderStatusChecker) shouldArchiveOrder(order domestic_models.CarOrderModel) bool {
	endDate, _ := time.Parse(time.RFC3339, order.EndDate)
	isOldEnough := time.Since(endDate) > 30*24*time.Hour
	status := constants.CarOrderStatus(order.Status)
	return isOldEnough && (status == constants.OrderStatusCompleted || status == constants.OrderStatusCancelled)
}
