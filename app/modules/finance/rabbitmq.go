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

// Payload struktur untuk menerima data laporan dinamis
type ReportRequestPayload struct {
	ReportType string `json:"report_type"` // e.g., "balance_sheet", "profit_loss"
	Format     string `json:"format"`      // "pdf" or "excel"
	DateStart  string `json:"date_start"`
	DateEnd    string `json:"date_end"`
	UserID     uint   `json:"user_id"`
}

// StartFinanceReportWorker (Phase 3)
func StartFinanceReportWorker() {
	if rabbitmq.Channel == nil { return }
	msgs, _ := rabbitmq.Channel.Consume("finance_report_generator", "fin_report_worker", true, false, false, false, nil)
	go func() {
		for d := range msgs {
			var req ReportRequestPayload
			if err := json.Unmarshal(d.Body, &req); err != nil {
				continue
			}

			log.Printf("📊 [Worker Finance] Memulai proses Generate %s (Format: %s) dari %s s/d %s", 
				req.ReportType, req.Format, req.DateStart, req.DateEnd)
			
			// SIMULASI PROSES BERAT (Misalnya: generate ribuan baris excel/pdf memakan waktu)
			time.Sleep(3 * time.Second) 

			fileName := req.ReportType + "_" + req.DateStart + "." + req.Format
			log.Printf("✅ [Worker Finance] Laporan %s selesai dibuat! Tersimpan di /uploads/reports/%s", req.ReportType, fileName)
			
			// SIMULASI MENGIRIM WEBSOCKET POP-UP KE USER YANG ME-REQUEST
			log.Printf("🔔 [Redis/Websocket] PUSH Pop-up ke Layar UserID %d: 'File %s siap diunduh!'", req.UserID, fileName)
		}
	}()
}
