package crm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ERP-System/pkg/rabbitmq"
	redisPkg "ERP-System/pkg/redis"
)

var ctx = context.Background()

func invalidateCRMLeadCache() {
	if redisPkg.Client != nil {
		redisPkg.Client.Del(ctx, "sales:crm:leads")
		fmt.Println("🧹 [Redis Clear] Invalidated Kanban Board Cache ('sales:crm:leads')")
	}
}

func CalculateLeadScore(data *Lead) {
	score := 0

	// 1. Kelengkapan Kontak (Profil Firmografis)
	if data.Email != "" {
		score += 15
	}
	if data.Phone != "" {
		score += 15
	}
	if data.PartnerID != nil && *data.PartnerID > 0 {
		score += 20 // Perusahaan/Customer terverifikasi
	}

	// 2. Nilai Transaksi & Ekspektasi Revenue
	if data.ExpectedRevenue >= 50000000 {
		score += 30 // Enterprise Tier (>50 Juta)
	} else if data.ExpectedRevenue >= 10000000 {
		score += 20 // Mid Tier (>10 Juta)
	} else if data.ExpectedRevenue > 0 {
		score += 10
	}

	// 3. Sumber & Saluran (UTM / Referensi)
	if data.ReferralCode != "" {
		score += 25 // Datang dari Rekomendasi/Afiliasi (Kualitas Tinggi)
	} else if data.UtmSource == "wa_blast" || data.UtmSource == "meta" {
		score += 15
	}

	// 4. Hitung Grade (Hot, Warm, Cold)
	data.LeadScore = score
	if score >= 75 {
		data.ScoreGrade = "Hot 🔥"
	} else if score >= 40 {
		data.ScoreGrade = "Warm ⚡"
	} else {
		data.ScoreGrade = "Cold ❄️"
	}
}

// Round-Robin Pools per Territory
var territoryReps = map[string][]string{
	"Jakarta":  {"Ahmad Dahlan", "Dewi Sartika", "Budi Santoso", "Rina Marlina"},
	"Surabaya": {"Bambang Wijaya", "Siti Rahmawati", "Surya Saputra"},
	"Medan":    {"Horas Simanjuntak", "Faisal Siregar"},
	"Bandung":  {"Cecep Hidayat", "Neneng Hasanah"},
	"Bali":     {"I Wayan Sudarta", "Ni Made Ayu"},
}
var roundRobinIndex = 0

func AssignLeadRoundRobin(data *Lead) {
	if data.AssignedSalespersonName != "" && data.AssignmentMethod == "Manual" {
		return
	}
	if data.Territory == "" {
		data.Territory = "Jakarta"
	}
	reps, exists := territoryReps[data.Territory]
	if !exists || len(reps) == 0 {
		reps = territoryReps["Jakarta"]
	}

	selectedRep := reps[roundRobinIndex%len(reps)]
	roundRobinIndex++

	data.AssignedSalespersonName = selectedRep
	data.AssignmentMethod = "Round-Robin (" + data.Territory + ")"
	log.Printf("🎯 [Lead Round-Robin] Prospek '%s' dialokasikan otomatis ke sales: %s (Area: %s)",
		data.Name, selectedRep, data.Territory)
}

func CreateLeadService(data *Lead) error {
	CalculateLeadScore(data)
	AssignLeadRoundRobin(data)
	err := CreateLead(data)
	if err == nil {
		invalidateCRMLeadCache()
	}
	return err
}

func GetAllLeadService() ([]Lead, error) {
	cacheKey := "sales:crm:leads"

	if redisPkg.Client != nil {
		redisRepo := NewCrmRedis(redisPkg.Client)
		cachedData, err := redisRepo.GetKanbanBoard(ctx, cacheKey)
		if err == nil && cachedData != "" {
			var leads []Lead
			_ = json.Unmarshal([]byte(cachedData), &leads)
			fmt.Println("🚀 [Redis Hit] Fetching Kanban Board (Leads) from Cache!")
			return leads, nil
		}
	}

	fmt.Println("🐢 [DB Hit] Fetching Kanban Board (Leads) from PostgreSQL...")
	leads, err := GetAllLead()

	if err == nil && redisPkg.Client != nil {
		bytes, _ := json.Marshal(leads)
		_ = redisPkg.Client.Set(ctx, cacheKey, bytes, 1*time.Hour).Err()
	}

	return leads, err
}

func GetPaginatedLeadService(offset int, limit int, search string, stageID uint, salespersonID uint, territory string) ([]Lead, int64, error) {
	return GetPaginatedLeads(offset, limit, search, stageID, salespersonID, territory)
}

func GetLeadByIDService(id uint) (*Lead, error) {
	return GetLeadByID(id)
}

func UpdateLeadService(data *Lead) error {
	CalculateLeadScore(data)
	err := UpdateLead(data)
	if err == nil {
		invalidateCRMLeadCache()
	}
	return err
}

func DeleteLeadService(id uint) error {
	err := DeleteLead(id)
	if err == nil {
		invalidateCRMLeadCache()
	}
	return err
}

func CreateSalesTeamService(data *SalesTeam) error        { return CreateSalesTeam(data) }
func GetAllSalesTeamService() ([]SalesTeam, error)        { return GetAllSalesTeam() }
func GetSalesTeamByIDService(id uint) (*SalesTeam, error) { return GetSalesTeamByID(id) }
func UpdateSalesTeamService(data *SalesTeam) error        { return UpdateSalesTeam(data) }
func DeleteSalesTeamService(id uint) error                { return DeleteSalesTeam(id) }

func CreateStageService(data *Stage) error        { return CreateStage(data) }
func GetAllStageService() ([]Stage, error)        { return GetAllStage() }
func GetStageByIDService(id uint) (*Stage, error) { return GetStageByID(id) }
func UpdateStageService(data *Stage) error        { return UpdateStage(data) }
func DeleteStageService(id uint) error            { return DeleteStage(id) }

func CreateActivityService(data *Activity) error        { return CreateActivity(data) }
func GetAllActivityService() ([]Activity, error)        { return GetAllActivity() }
func GetActivityByIDService(id uint) (*Activity, error) { return GetActivityByID(id) }
func UpdateActivityService(data *Activity) error        { return UpdateActivity(data) }
func DeleteActivityService(id uint) error               { return DeleteActivity(id) }

func CreateSalesCommissionService(data *SalesCommission) error { return CreateSalesCommission(data) }
func GetAllSalesCommissionService() ([]SalesCommission, error) { return GetAllSalesCommission() }
func GetSalesCommissionByIDService(id uint) (*SalesCommission, error) {
	return GetSalesCommissionByID(id)
}
func UpdateSalesCommissionService(data *SalesCommission) error { return UpdateSalesCommission(data) }
func DeleteSalesCommissionService(id uint) error               { return DeleteSalesCommission(id) }

func ExportCRMExcelService(userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"action": "export_crm", "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("📈 Event RabbitMQ: Export Excel CRM Pipeline dikirim ke antrean!")
	}
	return nil
}
