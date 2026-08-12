package mailer

import (
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
)

type EmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

// QueueEmail digunakan oleh modul bisnis (misal: Sales) untuk mengirim email notifikasi.
// Email tidak dikirim secara sinkron, melainkan dimasukkan ke dalam antrean RabbitMQ.
func QueueEmail(recipient, subject, body string) error {
	payload := EmailPayload{
		Recipient: recipient,
		Subject:   subject,
		Body:      body,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Jika RabbitMQ belum terhubung (misalnya saat testing lokal tanpa RabbitMQ server)
	if rabbitmq.Channel == nil {
		log.Println("⚠️ Peringatan: RabbitMQ Channel kosong. Menyimulasikan pengiriman email ke", recipient)
		return nil
	}

	// Memasukkan email ke antrean (Queue)
	err = rabbitmq.PublishEvent(rabbitmq.Channel, "email_queue", data)
	if err != nil {
		log.Printf("❌ Gagal memasukkan email ke queue: %v", err)
		return err
	}

	log.Println("📧 Email berhasil ditambahkan ke antrean (RabbitMQ)")
	return nil
}
