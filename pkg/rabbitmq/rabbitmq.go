package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func ConnectRabbitMQ() {
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	conn, err := NewRabbitMQConnection(url)
	if err != nil {
		log.Fatalf("❌ Gagal menyambung ke RabbitMQ setelah 15 percobaan: %v", err)
	}

	Conn = conn
	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("❌ Gagal membuka channel RabbitMQ: %v", err)
	}
}

func NewRabbitMQConnection(url string) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error
	for i := 1; i <= 15; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			log.Println("✅ RabbitMQ Connection berhasil dibuat!")
			return conn, nil
		}
		log.Printf("⏳ RabbitMQ belum siap (Percobaan %d/15). Menunggu 2 detik... Error: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	return nil, err
}

func PublishEvent(ch *amqp.Channel, queueName string, body []byte) error {
	_, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func PublishDelayedEvent(ch *amqp.Channel, targetQueue string, body []byte, ttlMS int) error {
	waitQueue := targetQueue + "_wait"
	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": targetQueue,
		"x-message-ttl":             int32(ttlMS),
	}

	_, err := ch.QueueDeclare(waitQueue, true, false, false, false, args)
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(targetQueue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		waitQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func Close() {
	if Channel != nil {
		Channel.Close()
	}
	if Conn != nil {
		Conn.Close()
	}
}
