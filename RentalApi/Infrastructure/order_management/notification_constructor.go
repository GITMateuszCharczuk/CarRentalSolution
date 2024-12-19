package order_management

import (
	"fmt"
	"rental-api/Domain/constants"
	domestic_models "rental-api/Domain/models/domestic"
	external_models "rental-api/Domain/models/external"
)

func (c *OrderStatusChecker) prepareNotification(order domestic_models.CarOrderModel, carOffer domestic_models.CarOfferModel, status constants.CarOrderStatus) external_models.Email {
	var subject, body string
	switch status {
	case constants.OrderStatusPreparing:
		subject = "Car Preparation Needed"
		body = c.preparePendingNotificationMessage(order, carOffer)
	case constants.OrderStatusCompleted:
		subject = "Overdue Car Rental"
		body = c.prepareOverdueNotificationMessage(order, carOffer)
	default:
		subject = "Car Rental Status Update"
		body = c.prepareDefaultNotificationMessage(order, carOffer, status)
	}

	return external_models.Email{
		To:      carOffer.CustodianEmail,
		Subject: subject,
		Body:    body,
	}
}

func (c *OrderStatusChecker) preparePendingNotificationMessage(order domestic_models.CarOrderModel, carOffer domestic_models.CarOfferModel) string {
	return fmt.Sprintf(
		"Car preparation needed for tomorrow's rental:\n"+
			"Car: %s\n"+
			"Order ID: %s\n"+
			"Start Date: %s\n"+
			"End Date: %s\n"+
			"Delivery Location: %s\n",
		carOffer.Heading,
		order.Id,
		order.StartDate,
		order.EndDate,
		order.DeliveryLocation,
	)
}

func (c *OrderStatusChecker) prepareOverdueNotificationMessage(order domestic_models.CarOrderModel, carOffer domestic_models.CarOfferModel) string {
	return fmt.Sprintf(
		"Overdue car rental notification:\n"+
			"Car: %s\n"+
			"Order ID: %s\n"+
			"End Date: %s\n"+
			"Return Location: %s\n"+
			"Please check the status of this rental immediately.",
		carOffer.Heading,
		order.Id,
		order.EndDate,
		order.ReturnLocation,
	)
}

func (c *OrderStatusChecker) prepareDefaultNotificationMessage(order domestic_models.CarOrderModel, carOffer domestic_models.CarOfferModel, status constants.CarOrderStatus) string {
	return fmt.Sprintf(
		"Car rental status update:\n"+
			"Car: %s\n"+
			"Order ID: %s\n"+
			"New Status: %s\n"+
			"Please take appropriate action.",
		carOffer.Heading,
		order.Id,
		status,
	)
}
