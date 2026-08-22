package auth

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

const (
	QueueEmailNotification = "auth_email_notification"
)

// EmailEvent represents the data needed to send an email
type EmailEvent struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// PublishEmailEvent pushes an email task to RabbitMQ
func PublishEmailEvent(to, subject, body string) error {
	if rabbitmq.Channel == nil {
		log.Println("⚠️ RabbitMQ channel is nil, skipping email publish")
		return nil
	}

	event := EmailEvent{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	bodyBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = rabbitmq.PublishEvent(rabbitmq.Channel, QueueEmailNotification, bodyBytes)
	if err != nil {
		log.Printf("❌ Failed to publish email event: %v", err)
		return err
	}

	log.Printf("📩 Email event published for %s", to)
	return nil
}

// StartEmailWorker starts a background goroutine to consume email events
func StartEmailWorker() {
	if rabbitmq.Channel == nil {
		log.Println("⚠️ RabbitMQ channel is nil, cannot start auth email worker")
		return
	}

	msgs, err := rabbitmq.Channel.Consume(
		QueueEmailNotification, // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // args
	)
	if err != nil {
		log.Printf("❌ Failed to register email consumer: %v", err)
		return
	}

	go func() {
		for d := range msgs {
			var event EmailEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("❌ Error decoding email event: %v", err)
				continue
			}

			log.Printf("✅ [Worker Auth] Sending email to: %s, Subject: %s", event.To, event.Subject)
			// TODO: Implement actual SMTP sending logic here
		}
	}()
}

// StartAuthCleanupWorker (Phase 5)
func StartAuthCleanupWorker() {
	if rabbitmq.Channel == nil { return }
	msgs, _ := rabbitmq.Channel.Consume("auth_delayed_cleanup", "auth_cleanup_worker", true, false, false, false, nil)
	go func() {
		for d := range msgs {
			log.Printf("🧹 [Worker Auth] Menghapus otomatis akun unverified (Waktu Habis): %s", string(d.Body))
		}
	}()
}
