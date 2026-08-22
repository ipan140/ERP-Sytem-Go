package website

import (
	"log"
	"ERP-System/pkg/rabbitmq"
)

func PublishActivityTracker(userID string, page string) {
	if rabbitmq.Channel == nil { return }
	log.Printf("🌐 [Website] Memasukkan aktivitas user %s di halaman %s ke RabbitMQ", userID, page)
	// (Simulasi Publish)
}
