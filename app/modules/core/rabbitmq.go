package core

import (
	"encoding/json"
	"log"
	"time"

	"ERP-System/pkg/rabbitmq"
)

const (
	QueueAuditLog = "core_audit_log"
)

type AuditEvent struct {
	Action    string    `json:"action"`     // CREATE, UPDATE, DELETE
	Module    string    `json:"module"`     // e.g., "sales", "hr"
	RecordID  string    `json:"record_id"`
	UserID    string    `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

// PublishAuditLog pushes an audit trail to RabbitMQ
func PublishAuditLog(action, module, recordID, userID string) error {
	if rabbitmq.Channel == nil {
		return nil
	}

	event := AuditEvent{
		Action:    action,
		Module:    module,
		RecordID:  recordID,
		UserID:    userID,
		Timestamp: time.Now(),
	}

	bodyBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = rabbitmq.PublishEvent(rabbitmq.Channel, QueueAuditLog, bodyBytes)
	if err != nil {
		log.Printf("❌ Failed to publish audit log: %v", err)
		return err
	}
	return nil
}

// StartAuditWorker starts a background goroutine to save audit logs to database
func StartAuditWorker() {
	if rabbitmq.Channel == nil {
		return
	}

	msgs, err := rabbitmq.Channel.Consume(
		QueueAuditLog, 
		"audit_worker",
		true,          
		false,         
		false,         
		false,         
		nil,           
	)
	if err != nil {
		log.Printf("❌ Failed to register audit consumer: %v", err)
		return
	}

	go func() {
		for d := range msgs {
			var event AuditEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				continue
			}

			log.Printf("📝 [Worker Audit] %s record %s in module %s by user %s", event.Action, event.RecordID, event.Module, event.UserID)
			// TODO: Save this event to the Audit Trail table in the DB
		}
	}()
}

// StartCoreMiscWorker (Phase 3 & Phase 5)
func StartCoreMiscWorker() {
	if rabbitmq.Channel == nil { return }
	
	// Excel Import
	msgsExcel, _ := rabbitmq.Channel.Consume("core_excel_import", "core_excel_worker", true, false, false, false, nil)
	go func() {
		for d := range msgsExcel {
			log.Printf("📑 [Worker Core] Memproses baris Excel: %s", string(d.Body))
		}
	}()

	// Outbound Webhooks
	msgsWebhook, _ := rabbitmq.Channel.Consume("core_outbound_webhook", "core_webhook_worker", true, false, false, false, nil)
	go func() {
		for d := range msgsWebhook {
			log.Printf("📡 [Worker Core] Mengirim Webhook ke Pihak Ketiga: %s", string(d.Body))
		}
	}()
}
