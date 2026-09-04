package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"

	"ERP-System/config"
	"github.com/brianvoe/gofakeit/v6"

	// Module Imports for struct usage
	"ERP-System/app/modules/auth"
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/finance/accounting"
	"ERP-System/app/modules/finance/invoicing"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/app/modules/hr/time_off"
	"ERP-System/app/modules/marketing/events"
	"ERP-System/app/modules/marketing/marketing_automation"
	"ERP-System/app/modules/marketing/mass_mailing"
	"ERP-System/app/modules/sales/crm"
	"ERP-System/app/modules/sales/sales_core"
	"ERP-System/app/modules/services/project"
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/app/modules/supply_chain/purchase"

	// Blank imports to trigger init() for AutoMigrate (63 Modules)
	_ "ERP-System/app/modules/core/artificial_intelligence"
	_ "ERP-System/app/modules/core/dashboards"
	_ "ERP-System/app/modules/core/discuss"
	_ "ERP-System/app/modules/core/documents"
	_ "ERP-System/app/modules/core/iot"
	_ "ERP-System/app/modules/core/knowledge"
	_ "ERP-System/app/modules/core/mailer"
	_ "ERP-System/app/modules/core/permissions"
	_ "ERP-System/app/modules/core/report"
	_ "ERP-System/app/modules/core/storage"
	_ "ERP-System/app/modules/core/user_roles"
	_ "ERP-System/app/modules/core/voip"
	_ "ERP-System/app/modules/core/whatsapp"
	_ "ERP-System/app/modules/finance/approvals"
	_ "ERP-System/app/modules/finance/consolidation"
	_ "ERP-System/app/modules/finance/documents"
	_ "ERP-System/app/modules/finance/expenses"
	_ "ERP-System/app/modules/finance/sign"
	_ "ERP-System/app/modules/finance/spreadsheet_bi"
	_ "ERP-System/app/modules/hr/appraisals"
	_ "ERP-System/app/modules/hr/attendances"
	_ "ERP-System/app/modules/hr/fleet"
	_ "ERP-System/app/modules/hr/lunch"
	_ "ERP-System/app/modules/hr/payroll"
	_ "ERP-System/app/modules/hr/recruitment"
	_ "ERP-System/app/modules/hr/referrals"
	_ "ERP-System/app/modules/marketing/sms_marketing"
	_ "ERP-System/app/modules/marketing/social_marketing"
	_ "ERP-System/app/modules/marketing/surveys"
	_ "ERP-System/app/modules/sales/point_of_sale"
	_ "ERP-System/app/modules/sales/rental"
	_ "ERP-System/app/modules/sales/subscriptions"
	_ "ERP-System/app/modules/services/appointments"
	_ "ERP-System/app/modules/services/field_service"
	_ "ERP-System/app/modules/services/helpdesk"
	_ "ERP-System/app/modules/services/planning"
	_ "ERP-System/app/modules/services/repairs"
	_ "ERP-System/app/modules/services/timesheets"
	_ "ERP-System/app/modules/supply_chain/barcode"
	_ "ERP-System/app/modules/supply_chain/maintenance"
	_ "ERP-System/app/modules/supply_chain/manufacturing"
	_ "ERP-System/app/modules/supply_chain/plm"
	_ "ERP-System/app/modules/supply_chain/quality"
	_ "ERP-System/app/modules/website/blog"
	_ "ERP-System/app/modules/website/ecommerce"
	_ "ERP-System/app/modules/website/elearning"
	_ "ERP-System/app/modules/website/forum"
	_ "ERP-System/app/modules/website/live_chat"
	_ "ERP-System/app/modules/website/website_builder"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// Matikan aturan Foreign Key agar tidak ada error constraint saat seeding
	config.DB.Exec("SET session_replication_role = 'replica';")

	gofakeit.Seed(0)
	rand.Seed(time.Now().UnixNano())

	for _, model := range config.ModelsToMigrate {
		seedModelSafe(model)
	}

	log.Println("=== 2. MEMULAI ERP SEEDER TERSTRUKTUR (FASE 1 - 8) ===")
	
	log.Println("[1/5] Menyiapkan Master Data...")
	
	company := auth.Company{Name: "PT. ERP Maju Bersama"}
	config.DB.Create(&company)

	var users []auth.User
	var emps []employees.Employee
	for i := 0; i < 10; i++ {
		u := auth.User{Name: gofakeit.Name(), Email: gofakeit.Email(), Password: "password123", CompanyID: company.ID}
		config.DB.Create(&u)
		users = append(users, u)

		emp := employees.Employee{Name: u.Name, UserID: &u.ID, WorkEmail: u.Email}
		config.DB.Create(&emp)
		emps = append(emps, emp)
	}

	var partners []base.Partner
	for i := 0; i < 20; i++ {
		p := base.Partner{Name: gofakeit.Company(), Email: gofakeit.Email(), IsCustomer: i%2 == 0, IsVendor: i%2 != 0}
		config.DB.Create(&p)
		partners = append(partners, p)
	}

	cat := inventory.ProductCategory{Name: "Elektronik"}
	config.DB.Create(&cat)

	var prods []inventory.ProductTemplate
	for i := 0; i < 20; i++ {
		tmpl := inventory.ProductTemplate{Name: gofakeit.ProductName(), CategoryID: cat.ID, ListPrice: float64(gofakeit.Price(100, 1000))}
		config.DB.Create(&tmpl)
		
		prod := inventory.Product{ProductTemplateID: tmpl.ID, DefaultCode: gofakeit.UUID()}
		config.DB.Create(&prod)
		prods = append(prods, tmpl)
	}

	// =========================================================================
	// FASE 3: CRM & SALES
	// =========================================================================
	log.Println("[2/5] Menyiapkan Data Penjualan (CRM & Sales)...")
	stage := crm.Stage{Name: "Qualified", Sequence: 1}
	config.DB.Create(&stage)

	for i := 0; i < 15; i++ {
		lead := crm.Lead{Name: "Peluang Bisnis " + gofakeit.Word(), PartnerID: &partners[i].ID, StageID: stage.ID}
		config.DB.Create(&lead)

		if partners[i].IsCustomer {
			so := sales_core.SaleOrder{Name: fmt.Sprintf("SO/2026/%04d", i+1), PartnerID: partners[i].ID, State: "sale"}
			config.DB.Create(&so)

			sol := sales_core.SaleOrderLine{OrderID: so.ID, ProductID: prods[i].ID, Quantity: 2, UnitPrice: prods[i].ListPrice}
			config.DB.Create(&sol)
		}
	}

	// =========================================================================
	// FASE 4: FINANCE INVOICING
	// =========================================================================
	log.Println("[3/5] Menyiapkan Data Keuangan (Invoicing)...")
	tax := invoicing.Tax{Name: "PPN 11%", Rate: 11.00}
	config.DB.Create(&tax)

	for i := 0; i < 10; i++ {
		if partners[i].IsCustomer {
			inv := invoicing.Invoice{Name: fmt.Sprintf("INV/2026/%04d", i+1), PartnerID: partners[i].ID, State: "posted"}
			config.DB.Create(&inv)

			line := invoicing.InvoiceLine{InvoiceID: inv.ID, Description: "Tagihan Produk", Quantity: 1, UnitPrice: 500000}
			config.DB.Create(&line)
		}
	}

	// =========================================================================
	// FASE 5: HR TIME OFF & SERVICES
	// =========================================================================
	log.Println("[4/5] Menyiapkan Data HR Cuti & Manajemen Proyek...")
	leaveType := time_off.LeaveType{Name: "Cuti Tahunan"}
	config.DB.Create(&leaveType)

	for i := 0; i < 5; i++ {
		req := time_off.LeaveRequest{EmployeeID: emps[i].ID, LeaveTypeID: leaveType.ID, NumberOfDays: 2, Status: "approved"}
		config.DB.Create(&req)

		proj := project.Project{Name: "Proyek " + gofakeit.Word(), ManagerID: emps[i].ID, CustomerID: partners[i].ID}
		config.DB.Create(&proj)

		task := project.Task{ProjectID: proj.ID, Name: "Desain UI", AssigneeID: users[i].ID, Stage: "todo"}
		config.DB.Create(&task)
	}

	// =========================================================================
	// FASE 6, 7, 8: PURCHASE, EVENTS, ACCOUNTING JOURNALS
	// =========================================================================
	log.Println("[5/5] Menyiapkan Purchase Orders, Tiket Event, & Jurnal...")
	for i := 0; i < 10; i++ {
		if partners[i].IsVendor {
			po := purchase.PurchaseOrder{Name: fmt.Sprintf("PO/2026/%04d", i+1), PartnerID: partners[i].ID, State: "purchase"}
			config.DB.Create(&po)
		}
	}

	event := events.Event{EventName: "Seminar Transformasi Digital & ERP", Location: "Grand Ballroom Hotel Indonesia", MaxCapacity: 100}
	config.DB.Create(&event)
	for i := 0; i < 20; i++ {
		pID := partners[i].ID
		ticket := events.EventTicket{
			EventID:       event.ID,
			CustomerID:    &pID,
			AttendeeName:  partners[i].Name,
			AttendeeEmail: partners[i].Email,
			Barcode:       gofakeit.UUID(),
		}
		config.DB.Create(&ticket)
	}

	journal := accounting.Journal{Code: "MISC", Name: "Operasional", Type: "general"}
	config.DB.Create(&journal)

	var account accounting.Account
	config.DB.First(&account)
	if account.ID != 0 {
		for i := 0; i < 20; i++ {
			entry := accounting.JournalEntry{Name: fmt.Sprintf("JRNL/%04d", i+1), JournalID: journal.ID, State: "posted"}
			config.DB.Create(&entry)

			item1 := accounting.JournalItem{EntryID: entry.ID, AccountID: account.ID, Debit: 100000, Credit: 0}
			config.DB.Create(&item1)
			
			item2 := accounting.JournalItem{EntryID: entry.ID, AccountID: account.ID, Debit: 0, Credit: 100000}
			config.DB.Create(&item2)
		}
	}

	// [FASE MARKETING] Seed Realistis Email Massal & UTM Tracker
	log.Println("[Marketing] Menyiapkan kampanye email realistis & UTM attribution...")
	config.DB.Exec("DELETE FROM marketing.utm_trackers;")
	config.DB.Exec("DELETE FROM marketing.mailing_campaigns;")

	c1 := mass_mailing.MailingCampaign{
		Name:           "Newsletter Bulanan Enterprise Tech Q3",
		Subject:        "Update Fitur ERP & Tips Efisiensi Operasional",
		TargetAudience: "Semua Kontak Pelanggan",
		Status:         "Sent",
		SentCount:      1250,
		OpenedCount:    750,
		ClickedCount:   225,
		BouncedCount:   12,
	}
	c2 := mass_mailing.MailingCampaign{
		Name:           "Promo Early Bird Renewal Kontrak Tahunan",
		Subject:        "Dapatkan Cashback 25% untuk Perpanjangan Layanan",
		TargetAudience: "Klien Prioritas VIP",
		Status:         "Sent",
		SentCount:      450,
		OpenedCount:    320,
		ClickedCount:   145,
		BouncedCount:   3,
	}
	c3 := mass_mailing.MailingCampaign{
		Name:           "Follow-Up Webinar Integrasi Supply Chain & Finance",
		Subject:        "Materi Presentasi & Akses Demo Gratis ERP Eksklusif",
		TargetAudience: "Prospek CRM Belum Closing",
		Status:         "Sent",
		SentCount:      680,
		OpenedCount:    390,
		ClickedCount:   118,
		BouncedCount:   5,
	}
	c4 := mass_mailing.MailingCampaign{
		Name:           "Program Re-Engagement Klien Pasif (Winback)",
		Subject:        "Klaim Sesi Konsultasi Bisnis Gratis Bersama Tim Ahli",
		TargetAudience: "Pelanggan Tidak Aktif (Winback)",
		Status:         "Scheduled",
		SentCount:      0,
		OpenedCount:    0,
		ClickedCount:   0,
		BouncedCount:   0,
	}
	c5 := mass_mailing.MailingCampaign{
		Name:           "Undangan Eksklusif Peluncuran Fitur AI Analytics",
		Subject:        "Daftar Lebih Awal: Akses Beta Modul AI Forecasting & BI",
		TargetAudience: "Klien Prioritas VIP",
		Status:         "Draft",
		SentCount:      0,
		OpenedCount:    0,
		ClickedCount:   0,
		BouncedCount:   0,
	}

	config.DB.Create(&c1)
	config.DB.Create(&c2)
	config.DB.Create(&c3)
	config.DB.Create(&c4)
	config.DB.Create(&c5)

	u1 := mass_mailing.UtmTracker{CampaignID: c1.ID, UtmSource: "newsletter", UtmMedium: "email", GeneratedRevenue: 37500000}
	u2 := mass_mailing.UtmTracker{CampaignID: c2.ID, UtmSource: "vip_promo", UtmMedium: "email_blast", GeneratedRevenue: 85000000}
	u3 := mass_mailing.UtmTracker{CampaignID: c3.ID, UtmSource: "webinar_crm", UtmMedium: "email", GeneratedRevenue: 42000000}
	u4 := mass_mailing.UtmTracker{CampaignID: c2.ID, UtmSource: "account_manager", UtmMedium: "direct_email", GeneratedRevenue: 28500000}
	config.DB.Create(&u1)
	config.DB.Create(&u2)
	config.DB.Create(&u3)
	config.DB.Create(&u4)

	// [FASE MARKETING AUTOMATION] Bersihkan dan Seed Alur Kerja Drip Nyata
	log.Println("[Marketing Automation] Menyiapkan skenario alur kerja drip otomatis...")
	config.DB.Exec("DELETE FROM marketing.workflow_activities;")
	config.DB.Exec("DELETE FROM marketing.automation_campaigns;")

	auto1 := marketing_automation.AutomationCampaign{
		Name:        "Onboarding Prospek Baru Website",
		TriggerType: "Ketika Lead Baru Masuk dari Form Web",
		Status:      "Active",
		TargetModel: "Sales Leads",
	}
	auto2 := marketing_automation.AutomationCampaign{
		Name:        "Retensi & Layanan Prioritas Klien VIP",
		TriggerType: "Ketika Faktur Invoice Lunas (Won Deal)",
		Status:      "Active",
		TargetModel: "Pelanggan VIP",
	}
	auto3 := marketing_automation.AutomationCampaign{
		Name:        "Drip Seri Edukasi Paska Webinar",
		TriggerType: "Ketika Tiket Acara Terdaftar",
		Status:      "Active",
		TargetModel: "Peserta Acara",
	}
	auto4 := marketing_automation.AutomationCampaign{
		Name:        "Follow-Up Survei Kepuasan Net Promoter (NPS)",
		TriggerType: "Ketika Responden Mengisi Survei",
		Status:      "Paused",
		TargetModel: "Pelanggan VIP",
	}
	auto5 := marketing_automation.AutomationCampaign{
		Name:        "Nurturing Prospek Dingin Q4",
		TriggerType: "Ketika Lead Baru Masuk dari Form Web",
		Status:      "Draft",
		TargetModel: "Sales Leads",
	}

	config.DB.Create(&auto1)
	config.DB.Create(&auto2)
	config.DB.Create(&auto3)
	config.DB.Create(&auto4)
	config.DB.Create(&auto5)

	// Activities untuk auto1 (Onboarding Prospek Baru Website)
	act1 := marketing_automation.WorkflowActivity{
		CampaignID:   auto1.ID,
		ActivityName: "Kirim Email Sambutan & Katalog Produk Enterprise",
		ActionType:   "Email",
		DelayHours:   1,
		Condition:    "Always",
	}
	act2 := marketing_automation.WorkflowActivity{
		CampaignID:   auto1.ID,
		ActivityName: "Kirim WhatsApp Penawaran Demo & Jadwal Konsultasi",
		ActionType:   "SMS",
		DelayHours:   24,
		Condition:    "Opened",
	}
	act3 := marketing_automation.WorkflowActivity{
		CampaignID:   auto1.ID,
		ActivityName: "Tugaskan Sales Representative untuk Telepon Follow-up",
		ActionType:   "Notification",
		DelayHours:   48,
		Condition:    "Always",
	}

	// Activities untuk auto2 (Retensi Klien VIP)
	act4 := marketing_automation.WorkflowActivity{
		CampaignID:   auto2.ID,
		ActivityName: "Kirim Ucapan Terima Kasih & Akses Portal VIP",
		ActionType:   "Email",
		DelayHours:   1,
		Condition:    "Always",
	}
	act5 := marketing_automation.WorkflowActivity{
		CampaignID:   auto2.ID,
		ActivityName: "Kirim Voucher Diskon Maintenance Kontrak Tahunan",
		ActionType:   "Email",
		DelayHours:   72,
		Condition:    "Always",
	}

	// Activities untuk auto3 (Drip Seri Edukasi)
	act6 := marketing_automation.WorkflowActivity{
		CampaignID:   auto3.ID,
		ActivityName: "Kirim Rekaman Webinar & E-Book PDF Solusi",
		ActionType:   "Email",
		DelayHours:   2,
		Condition:    "Always",
	}

	config.DB.Create(&act1)
	config.DB.Create(&act2)
	config.DB.Create(&act3)
	config.DB.Create(&act4)
	config.DB.Create(&act5)
	config.DB.Create(&act6)

	// Nyalakan kembali aturan Foreign Key
	config.DB.Exec("SET session_replication_role = 'origin';")

	log.Println("=== SEEDING TERSTRUKTUR SELESAI DENGAN SUKSES! ===")
}

func seedModelSafe(model interface{}) {
	defer func() {
		if r := recover(); r != nil {
			// Abaikan panic
		}
	}()

	modelType := reflect.TypeOf(model).Elem()
	successCount := 0

	for i := 0; i < 5; i++ {
		newObj := reflect.New(modelType).Interface()
		val := reflect.ValueOf(newObj).Elem()

		for j := 0; j < val.NumField(); j++ {
			field := val.Field(j)
			fieldType := modelType.Field(j)

			if !field.CanSet() {
				continue
			}

			// Lewati ID agar GORM yang mengatur auto-increment
			if fieldType.Name == "ID" {
				continue
			}

			// Isi data acak hanya pada tipe data dasar (Hindari pointer/struct untuk cegah infinite loop)
			switch field.Kind() {
			case reflect.String:
				field.SetString(gofakeit.Word())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				field.SetInt(int64(gofakeit.Number(1, 100)))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				field.SetUint(uint64(gofakeit.Number(1, 100)))
			case reflect.Float32, reflect.Float64:
				field.SetFloat(gofakeit.Float64Range(10.0, 1000.0))
			case reflect.Bool:
				field.SetBool(gofakeit.Bool())
			case reflect.Struct:
				if field.Type().String() == "time.Time" {
					field.Set(reflect.ValueOf(gofakeit.Date()))
				}
			}
		}

		if err := config.DB.Create(newObj).Error; err == nil {
			successCount++
		}
	}

	if successCount > 0 {
		log.Printf("-> Terisi %d data acak di tabel: %s\n", successCount, modelType.Name())
	}
}
