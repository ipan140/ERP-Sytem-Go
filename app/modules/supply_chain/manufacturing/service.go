package manufacturing

import (
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
)

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
