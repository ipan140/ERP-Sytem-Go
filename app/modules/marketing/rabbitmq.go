package marketing

import (
	"log"

	"ERP-System/pkg/rabbitmq"
)

func StartMarketingWorker() {
	if rabbitmq.Channel == nil { return }
	
	// Queue untuk Broadcast Email (Phase 2)
	msgsBroadcast, _ := rabbitmq.Channel.Consume("marketing_broadcast", "mkt_broadcast_worker", true, false, false, false, nil)
	go func() {
		for d := range msgsBroadcast {
			log.Printf("📢 [Worker Marketing] Mengirim blast email massal: %s", string(d.Body))
		}
	}()

	// Queue untuk Follow-up Cart Abandonment (Phase 5)
	msgsFollowUp, _ := rabbitmq.Channel.Consume("marketing_cart_reminder", "mkt_reminder_worker", true, false, false, false, nil)
	go func() {
		for d := range msgsFollowUp {
			log.Printf("🛒 [Worker Marketing] Mengirim reminder keranjang belanja (Delayed) untuk pelanggan: %s", string(d.Body))
		}
	}()
}
