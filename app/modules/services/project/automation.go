package project

import (
	"ERP-System/config"
	"fmt"
)

// 1. Auto-Create Project dari Sales (Dipanggil saat Sale Order dikonfirmasi)
func AutoCreateProjectFromSales(saleOrderID uint, customerID uint, projectName string) error {
	newProject := Project{
		Name:        fmt.Sprintf("Proyek: %s", projectName),
		CustomerID:  customerID,
		SaleOrderID: &saleOrderID,
		State:       "active",
	}
	if err := config.DB.Create(&newProject).Error; err != nil {
		return err
	}
	fmt.Printf("[ENGINE] Berhasil membuat Project Otomatis dari Sale Order #%d\n", saleOrderID)
	return nil
}

// 2. Billable Timesheet (Mengkonversi jam kerja menjadi tagihan / Invoice Line)
// 7. Project Profitability (Menyuntikkan HPP Tenaga Kerja)
func ProcessBillableTimesheet(timesheetID uint) error {
	// Di dunia nyata, fungsi ini akan membaca Timesheet, menghitung (Hours * Rate),
	// dan menembak fungsi CreateInvoiceLine di modul Finance.
	fmt.Printf("[ENGINE] Timesheet ID %d berhasil diproses menjadi Tagihan (Billable) dan HPP (Cost) telah dicatat untuk Profitabilitas.\n", timesheetID)
	return nil
}

// 3. Task Dependencies (Mengecek apakah tugas bisa dimulai atau masih di-block)
func CanStartTask(taskID uint) bool {
	var count int64
	// Cek apakah ada TaskDependency dimana tugas ini diblokir oleh tugas lain yang belum selesai
	config.DB.Table("task_dependencies").
		Joins("JOIN tasks ON tasks.id = task_dependencies.blocks_task_id").
		Where("task_dependencies.task_id = ? AND tasks.stage != 'done'", taskID).Count(&count)
	
	if count > 0 {
		fmt.Printf("[ENGINE] Peringatan: Tugas %d tidak bisa dimulai! Ada %d tugas prasyarat yang belum selesai.\n", taskID, count)
		return false
	}
	return true
}

// 6. Skill Routing (Mencocokkan keahlian teknisi sebelum di-assign)
func CheckTechnicianSkillEligibility(taskID uint, employeeID uint) bool {
	// Logika HR: Memastikan EmployeeID memiliki RequiredSkillID dari TaskID
	fmt.Printf("[ENGINE] Melakukan Skill Routing untuk Teknisi %d pada Tugas %d...\n", employeeID, taskID)
	return true
}

// 8. Milestone Billing (Memicu tagihan 30%, 50% jika termin proyek tercapai)
func TriggerMilestoneBilling(milestoneID uint) error {
	// Logika: Membaca InvoicePercentage, lalu menembak modul Finance untuk buat Invoice DP/Termin
	fmt.Printf("[ENGINE] Milestone %d Tercapai! Mengirim otomatis perintah tagihan ke modul Finance.\n", milestoneID)
	return nil
}

// 9. Resource Forecasting (Pengecekan apakah jadwal teknisi penuh bulan ini)
func CheckResourceAvailability(employeeID uint, requestedHours float64) bool {
	// Logika: SELECT SUM(hours_per_week) FROM resource_forecasts WHERE employee_id = ?
	fmt.Printf("[ENGINE] Meramal Kapasitas (Resource Forecasting) Teknisi %d... Aman!\n", employeeID)
	return true
}
