package services

import (
	"log"
	"ERP-System/pkg/rabbitmq"
)

func StartServicesWorker() {
	if rabbitmq.Channel == nil { return }
	
	// Queue untuk SLA Escalation (Phase 5)
	msgsSLA, _ := rabbitmq.Channel.Consume("services_sla_escalation", "services_sla_worker", true, false, false, false, nil)
	go func() {
		for d := range msgsSLA {
			log.Printf("🚨 [Worker Services] TIKET SLA TERLAMBAT! Mengirim peringatan ke Manajer untuk Tiket: %s", string(d.Body))
		}
	}()
}
