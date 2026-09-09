package timesheets

import (
	"ERP-System/app/modules/services/project"
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
	"time"
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

func GetPaginatedTimesheetService(offset int, limit int, search string, projectID uint, employeeID uint, billable string, status string, companyID uint) ([]Timesheet, int64, error) {
	return GetPaginatedTimesheets(offset, limit, search, projectID, employeeID, billable, status, companyID)
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

// Fase 2: Workflow Approval Services
func SubmitTimesheetService(id uint) error {
	updates := map[string]interface{}{
		"status": "submitted",
	}
	return UpdateTimesheetStatus(id, updates)
}

func ApproveTimesheetService(id uint, approverID uint) error {
	now := time.Now()
	updates := map[string]interface{}{
		"status":          "approved",
		"approved_by_id": approverID,
		"approved_at":     now,
		"rejection_reason": "",
	}
	return UpdateTimesheetStatus(id, updates)
}

func RejectTimesheetService(id uint, reason string) error {
	updates := map[string]interface{}{
		"status":          "rejected",
		"rejection_reason": reason,
	}
	return UpdateTimesheetStatus(id, updates)
}

func BulkApproveTimesheetsService(ids []uint, approverID uint) error {
	return BulkApproveTimesheets(ids, approverID)
}

func GenerateTimesheetPDFService(projectID uint, userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"document_type": "timesheet", "document_id": projectID, "format": "pdf", "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("⏱️ Event RabbitMQ: Generate PDF Timesheet Project %d dikirim ke antrean!", projectID)
	}
	return nil
}
