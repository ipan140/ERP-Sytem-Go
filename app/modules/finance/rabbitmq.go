package finance

import (
	"encoding/json"
	"log"
	"time"

	"ERP-System/app/modules/finance/invoicing"
	"ERP-System/pkg/rabbitmq"
)

const (
	EventOrderCompleted = "sales_order_completed" // Listens to sales event
)

// OrderEvent structure to decode from Sales
type OrderEvent struct {
	OrderID     string  `json:"order_id"`
	TotalAmount float64 `json:"total_amount"`
	CustomerID  string  `json:"customer_id"`
}

// StartFinanceWorker starts a background goroutine to consume cross-module events
func StartFinanceWorker() {
	if rabbitmq.Channel == nil {
		log.Println("⚠️ RabbitMQ channel is nil, cannot start finance worker")
		return
	}

	// Make sure the queue is declared before consuming
	_, err := rabbitmq.Channel.QueueDeclare(EventOrderCompleted, true, false, false, false, nil)
	if err != nil {
		log.Printf("❌ Failed to declare finance consumer queue: %v", err)
	}

	msgs, err := rabbitmq.Channel.Consume(
		EventOrderCompleted, // queue
		"finance_worker",    // consumer name
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		log.Printf("❌ Failed to register finance consumer: %v", err)
		return
	}

	go func() {
		for d := range msgs {
			var event OrderEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("❌ Error decoding order event in finance: %v", err)
				continue
			}

			log.Printf("✅ [Worker Finance] Generating draft invoice for Order %s (Amount: %.2f)", event.OrderID, event.TotalAmount)
			
			// [RabbitMQ] - Fase 4: Eksekusi otomatis dari event
			invoice := &invoicing.Invoice{
				Name:          "INV/AUTO/" + event.OrderID,
				PartnerID:     1, // Seharusnya dari event.CustomerID setelah dikonversi
				InvoiceDate:   time.Now(),
				DueDate:       time.Now().AddDate(0, 0, 14),
				State:         "draft",
				AmountUntaxed: event.TotalAmount, // Sederhana
				AmountTax:     0,
				AmountTotal:   event.TotalAmount,
			}
			
			if err := invoicing.CreateInvoiceService(invoice); err != nil {
				log.Printf("❌ Failed to create invoice: %v", err)
			} else {
				log.Printf("✅ Success Auto-Create Invoice for Order: %s", event.OrderID)
			}
		}
	}()
}

// StartFinanceReportWorker (Phase 3)
func StartFinanceReportWorker() {
	if rabbitmq.Channel == nil { return }
	msgs, _ := rabbitmq.Channel.Consume("finance_report_generator", "fin_report_worker", true, false, false, false, nil)
	go func() {
		for d := range msgs {
			log.Printf("📊 [Worker Finance] Meng-generate PDF Laporan Berat untuk periode: %s", string(d.Body))
		}
	}()
}
