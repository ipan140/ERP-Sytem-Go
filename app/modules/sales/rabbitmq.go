package sales

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

const (
	EventOrderCompleted = "sales_order_completed"
)

// OrderCompletedEvent data structure
type OrderCompletedEvent struct {
	OrderID     string  `json:"order_id"`
	TotalAmount float64 `json:"total_amount"`
	CustomerID  string  `json:"customer_id"`
}

// PublishOrderCompletedEvent publishes an event when a sales order is completed
func PublishOrderCompletedEvent(orderID string, amount float64, customerID string) error {
	if rabbitmq.Channel == nil {
		log.Println("⚠️ RabbitMQ channel is nil, skipping order completed event")
		return nil
	}

	event := OrderCompletedEvent{
		OrderID:     orderID,
		TotalAmount: amount,
		CustomerID:  customerID,
	}

	bodyBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = rabbitmq.PublishEvent(rabbitmq.Channel, EventOrderCompleted, bodyBytes)
	if err != nil {
		log.Printf("❌ Failed to publish sales order event: %v", err)
		return err
	}

	log.Printf("🛍️ Order completed event published for OrderID %s", orderID)
	return nil
}

// PublishDelayedCancelOrderEvent sends a delayed message to cancel unpaid orders (Phase 5)
func PublishDelayedCancelOrderEvent(orderID string, delayMS int) error {
	if rabbitmq.Channel == nil {
		return nil
	}
	
	event := map[string]string{"order_id": orderID}
	bodyBytes, _ := json.Marshal(event)

	// Publish ke queue tertunda (misal: "sales_auto_cancel")
	err := rabbitmq.PublishDelayedEvent(rabbitmq.Channel, "sales_auto_cancel", bodyBytes, delayMS)
	if err != nil {
		log.Printf("❌ Failed to publish delayed cancel event: %v", err)
		return err
	}
	
	log.Printf("⏳ Delayed Cancel Order event published for OrderID %s (Delay: %d ms)", orderID, delayMS)
	return nil
}

// StartSalesDelayedWorker listens for the expired orders
func StartSalesDelayedWorker() {
	if rabbitmq.Channel == nil {
		return
	}

	msgs, err := rabbitmq.Channel.Consume(
		"sales_auto_cancel", // Target queue (setelah delay habis)
		"sales_cancel_worker",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	go func() {
		for d := range msgs {
			var event map[string]string
			if err := json.Unmarshal(d.Body, &event); err != nil {
				continue
			}
			log.Printf("⏰ [Worker Sales] Waktu habis! Membatalkan otomatis Order ID %s jika belum dibayar", event["order_id"])
			// TODO: Cek database apakah order status = "draft" atau "sent", jika ya, ubah ke "cancel"
		}
	}()
}
