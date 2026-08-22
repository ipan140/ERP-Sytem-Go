package timesheets

import (
	"ERP-System/app/modules/services/project"
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
)

func CreateTimesheetService(data *Timesheet) error {
	if data.IsBillable {
		project.ProcessBillableTimesheet(data.ID)
	}

	return CreateTimesheet(data)
}

func GetAllTimesheetService() ([]Timesheet, error) {
	return GetAllTimesheet()
}

func GetTimesheetByIDService(id uint) (*Timesheet, error) {
	return GetTimesheetByID(id)
}

func UpdateTimesheetService(data *Timesheet) error {
	if data.IsBillable {
		project.ProcessBillableTimesheet(data.ID)
	}

	return UpdateTimesheet(data)
}

func DeleteTimesheetService(id uint) error {
	return DeleteTimesheet(id)
}

func GenerateTimesheetPDFService(projectID uint, userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"document_type": "timesheet", "document_id": projectID, "format": "pdf", "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("⏱️ Event RabbitMQ: Generate PDF Timesheet Project %d dikirim ke antrean!", projectID)
	}
	return nil
}
