package manufacturing

import (
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
	"time"
)

// Enterprise SCM Production Services (Fase 3)
func GetPaginatedProductionsService(offset, limit int, search, state string, productID uint) ([]MrpProduction, int64, error) {
	return GetPaginatedProductions(offset, limit, search, state, productID)
}

func GetMrpSummaryService() (MrpSummary, error) {
	return GetMrpSummary()
}

func CreateProductionWithSequenceService(req CreateMORequest) (*MrpProduction, error) {
	return CreateProductionWithSequence(req)
}

func ConfirmProductionService(id uint) (*MrpProduction, error) {
	return ConfirmProduction(id)
}

func StartProductionService(id uint) (*MrpProduction, error) {
	return StartProduction(id)
}

func FinishProductionService(id uint) (*MrpProduction, error) {
	return FinishProduction(id)
}

func CancelProductionService(id uint) (*MrpProduction, error) {
	return CancelProduction(id)
}

func GetPaginatedBomsService(offset, limit int, search string) ([]MrpBom, int64, error) {
	return GetPaginatedBoms(offset, limit, search)
}

func CreateBomWithLinesService(req CreateBomRequest) (*MrpBom, error) {
	return CreateBomWithLines(req)
}

func CreateMrpProductionService(data *MrpProduction) error {
	return CreateMrpProduction(data)
}

func GetAllMrpProductionService() ([]MrpProduction, error) {
	return GetAllMrpProduction()
}

func GetMrpProductionByIDService(id uint) (*MrpProduction, error) {
	return GetMrpProductionByID(id)
}

func UpdateMrpProductionService(data *MrpProduction) error {
	return UpdateMrpProduction(data)
}

func DeleteMrpProductionService(id uint) error {
	return DeleteMrpProduction(id)
}

func CreateMrpWorkcenterService(data *MrpWorkcenter) error        { return CreateMrpWorkcenter(data) }
func GetAllMrpWorkcenterService() ([]MrpWorkcenter, error)        { return GetAllMrpWorkcenter() }
func GetMrpWorkcenterByIDService(id uint) (*MrpWorkcenter, error) { return GetMrpWorkcenterByID(id) }
func UpdateMrpWorkcenterService(data *MrpWorkcenter) error        { return UpdateMrpWorkcenter(data) }
func DeleteMrpWorkcenterService(id uint) error                    { return DeleteMrpWorkcenter(id) }

func CreateMrpBomService(data *MrpBom) error        { return CreateMrpBom(data) }
func GetAllMrpBomService() ([]MrpBom, error)        { return GetAllMrpBom() }
func GetMrpBomByIDService(id uint) (*MrpBom, error) { return GetMrpBomByID(id) }
func UpdateMrpBomService(data *MrpBom) error        { return UpdateMrpBom(data) }
func DeleteMrpBomService(id uint) error             { return DeleteMrpBom(id) }

func CreateMrpBomLineService(data *MrpBomLine) error        { return CreateMrpBomLine(data) }
func GetAllMrpBomLineService() ([]MrpBomLine, error)        { return GetAllMrpBomLine() }
func GetMrpBomLineByIDService(id uint) (*MrpBomLine, error) { return GetMrpBomLineByID(id) }
func UpdateMrpBomLineService(data *MrpBomLine) error        { return UpdateMrpBomLine(data) }
func DeleteMrpBomLineService(id uint) error                 { return DeleteMrpBomLine(id) }

func CreateMrpBomByproductService(data *MrpBomByproduct) error { return CreateMrpBomByproduct(data) }
func GetAllMrpBomByproductService() ([]MrpBomByproduct, error) { return GetAllMrpBomByproduct() }
func GetMrpBomByproductByIDService(id uint) (*MrpBomByproduct, error) {
	return GetMrpBomByproductByID(id)
}
func UpdateMrpBomByproductService(data *MrpBomByproduct) error { return UpdateMrpBomByproduct(data) }
func DeleteMrpBomByproductService(id uint) error               { return DeleteMrpBomByproduct(id) }

func CreateMrpWorkorderService(data *MrpWorkorder) error        { return CreateMrpWorkorder(data) }
func GetAllMrpWorkorderService() ([]MrpWorkorder, error)        { return GetAllMrpWorkorder() }
func GetMrpWorkorderByIDService(id uint) (*MrpWorkorder, error) { return GetMrpWorkorderByID(id) }
func UpdateMrpWorkorderService(data *MrpWorkorder) error        { return UpdateMrpWorkorder(data) }
func DeleteMrpWorkorderService(id uint) error                   { return DeleteMrpWorkorder(id) }

func GenerateWorkOrderPDFService(woID uint, userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"document_type": "work_order", "document_id": woID, "format": "pdf", "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("🏭 Event RabbitMQ: Generate PDF Work Order %d dikirim ke antrean!", woID)
	}
	return nil
}

// FASE 9: MPS Service
func CreateMPSService(req CreateMPSRequest) (*MrpProductionSchedule, error) {
	parsedDate, err := time.Parse("2006-01-02", req.DatePlanned)
	if err != nil {
		return nil, err
	}
	mps := MrpProductionSchedule{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		DatePlanned: parsedDate,
		ForecastQty: req.ForecastQty,
	}
	if err := CreateMrpProductionSchedule(&mps); err != nil {
		return nil, err
	}
	return &mps, nil
}

func GenerateMOFromMPSService(mpsID uint) error {
	mps, err := GetMrpProductionScheduleByID(mpsID)
	if err != nil {
		return err
	}
	qtyToProduce := mps.ForecastQty - mps.ActualQty
	if qtyToProduce <= 0 {
		return nil
	}

	dateStr := mps.DatePlanned.Format(time.RFC3339)
	
	// Cari BOM default untuk product
	var boms []MrpBom
	boms, _ = GetAllMrpBom()
	var bomID *uint
	for _, b := range boms {
		if b.ProductID == mps.ProductID {
			bomID = &b.ID
			break
		}
	}

	req := CreateMORequest{
		ProductID:   mps.ProductID,
		BomID:       bomID,
		ProductQty:  qtyToProduce,
		WarehouseID: &mps.WarehouseID,
		DatePlanned: &dateStr,
		Notes:       "Generated from MPS",
	}

	if _, err := CreateProductionWithSequence(req); err != nil {
		return err
	}

	mps.ActualQty += qtyToProduce
	return UpdateMrpProductionSchedule(mps)
}

// FASE 9: OEE Service
func LogOEEService(req LogOEERequest) error {
	now := time.Now()
	oee := MrpWorkcenterProductivity{
		WorkcenterID: req.WorkcenterID,
		LossType:     req.LossType,
		Duration:     req.Duration,
		DateStart:    now,
	}
	return CreateMrpWorkcenterProductivity(&oee)
}

func GetWorkcenterOEEService(workcenterID uint) (OEESummary, error) {
	logs, _ := GetAllMrpWorkcenterProductivity()
	var totalTime, downtime float64
	for _, l := range logs {
		if l.WorkcenterID == workcenterID {
			totalTime += l.Duration
			if l.LossType != "productive" {
				downtime += l.Duration
			}
		}
	}

	avail := 100.0
	if totalTime > 0 {
		avail = ((totalTime - downtime) / totalTime) * 100
	}

	// Sederhanakan kalkulasi untuk OEE
	return OEESummary{
		Availability: avail,
		Performance:  100, // Dummy until we integrate cycle times
		Quality:      100, // Dummy until we integrate scrap
		OEE:          avail * 1.0 * 1.0,
	}, nil
}
