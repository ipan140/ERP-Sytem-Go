package supply_chain

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

const (
	EventOrderCompleted = "sales_order_completed" // Listens to sales event
)

type OrderEvent struct {
	OrderID     string  `json:"order_id"`
	TotalAmount float64 `json:"total_amount"`
	CustomerID  string  `json:"customer_id"`
}

// StartInventoryWorker starts a background goroutine to consume sales events
func StartInventoryWorker() {
	if rabbitmq.Channel == nil {
		log.Println("⚠️ RabbitMQ channel is nil, cannot start inventory worker")
		return
	}

	_, err := rabbitmq.Channel.QueueDeclare(EventOrderCompleted, true, false, false, false, nil)
	if err != nil {
		log.Printf("❌ Failed to declare inventory consumer queue: %v", err)
	}

	msgs, err := rabbitmq.Channel.Consume(
		EventOrderCompleted, // queue
		"inventory_worker",  // consumer name
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		log.Printf("❌ Failed to register inventory consumer: %v", err)
		return
	}

	go func() {
		for d := range msgs {
			var event OrderEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("❌ Error decoding order event in inventory: %v", err)
				continue
			}

			log.Printf("📦 [Worker Supply Chain] Memotong stok otomatis untuk Order ID %s", event.OrderID)
			// TODO: Add logic to fetch Order Lines and decrease stock via inventory.UpdateProductService or StockQuant
		}
	}()
}
