package helpdesk

import (
	"fmt"
	"strings"
)

// 4. Helpdesk SLA (Robot Pengawas Waktu)
func CheckSLAViolations() {
	// Di dunia nyata, ini dipasang di CRON JOB (berjalan setiap 5 menit)
	// Mengecek apakah (Waktu_Sekarang - CreatedAt) > TargetHours dari tabel HelpdeskSLA
	fmt.Println("[ENGINE SLA] Memindai tiket yang melewati batas waktu respon... Auto-Escalate!")
}

// 10. Helpdesk Canned Responses (Menjawab otomatis komplain pelanggan dengan AI/Keyword)
func GetAutoResponse(customerIssue string) string {
	customerIssue = strings.ToLower(customerIssue)
	if strings.Contains(customerIssue, "mati") {
		return "Canned Response: Mohon maaf atas gangguan ini. Tim lapangan kami segera meluncur!"
	}
	return "CS kami akan segera membalas tiket Anda."
}
