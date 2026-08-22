package hr

import (
	"log"

	"ERP-System/pkg/rabbitmq"
)

func StartHRWorker() {
	if rabbitmq.Channel == nil { return }
	
	// Queue untuk Notifikasi Cuti (Phase 2)
	msgsLeave, _ := rabbitmq.Channel.Consume("hr_leave_notification", "hr_leave_worker", true, false, false, false, nil)
	go func() {
		for d := range msgsLeave {
			log.Printf("📩 [Worker HR] Mengirim email notifikasi cuti: %s", string(d.Body))
		}
	}()

	// Queue untuk Payroll Massal (Phase 3)
	msgsPayroll, _ := rabbitmq.Channel.Consume("hr_payroll_generate", "hr_payroll_worker", true, false, false, false, nil)
	go func() {
		for d := range msgsPayroll {
			log.Printf("💸 [Worker HR] Memproses generate PDF Slip Gaji Massal untuk Batch: %s", string(d.Body))
		}
	}()
}
