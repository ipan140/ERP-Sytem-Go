package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"ERP-System/config"
	"ERP-System/app/modules/auth"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/app/modules/hr/recruitment"
	"ERP-System/app/modules/website/website_builder"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// Bypass Foreign Key constraints during seeding
	config.DB.Exec("SET session_replication_role = 'replica';")
	rand.Seed(time.Now().UnixNano())

	log.Println("🚀 [1/6] Menyiapkan Perusahaan & Master Departemen...")

	var company auth.Company
	if err := config.DB.FirstOrCreate(&company, auth.Company{Name: "PT. Nusantara Prima Solusindo"}).Error; err != nil {
		log.Fatalf("Gagal inisialisasi company: %v", err)
	}

	deptNames := []string{
		"Direksi & Eksekutif",
		"Human Resources & GA",
		"Teknologi Informasi & Digital",
		"Keuangan, Akuntansi & Pajak",
		"Sales & Business Development",
		"Supply Chain & Operasional",
	}

	depts := make(map[string]*employees.Department)
	for _, name := range deptNames {
		var dept employees.Department
		config.DB.Where("name = ?", name).FirstOrCreate(&dept, employees.Department{
			Name:      name,
			CreatedAt: time.Now(),
		})
		depts[name] = &dept
	}

	log.Println("👔 [2/6] Menyiapkan Posisi Pekerjaan (Job Positions)...")

	jobsList := []struct {
		Dept string
		Name string
	}{
		{"Direksi & Eksekutif", "Chief Executive Officer (CEO)"},
		{"Direksi & Eksekutif", "Chief Technology Officer (CTO)"},
		{"Direksi & Eksekutif", "Chief Financial Officer (CFO)"},
		{"Human Resources & GA", "HR Director"},
		{"Human Resources & GA", "Talent Acquisition Lead"},
		{"Human Resources & GA", "Compensation & Benefit Specialist"},
		{"Human Resources & GA", "People Operations Staff"},
		{"Teknologi Informasi & Digital", "VP of Engineering"},
		{"Teknologi Informasi & Digital", "Principal Software Architect"},
		{"Teknologi Informasi & Digital", "Senior Backend Go Engineer"},
		{"Teknologi Informasi & Digital", "Senior Frontend Vue Engineer"},
		{"Teknologi Informasi & Digital", "DevOps & Cloud Engineer"},
		{"Teknologi Informasi & Digital", "QA Automation Engineer"},
		{"Teknologi Informasi & Digital", "IT Support Specialist"},
		{"Keuangan, Akuntansi & Pajak", "Finance Director"},
		{"Keuangan, Akuntansi & Pajak", "Accounting Manager"},
		{"Keuangan, Akuntansi & Pajak", "Senior Tax Specialist"},
		{"Keuangan, Akuntansi & Pajak", "Account Payable & Receivable"},
		{"Keuangan, Akuntansi & Pajak", "Treasury & Payroll Officer"},
		{"Sales & Business Development", "Chief Commercial Officer"},
		{"Sales & Business Development", "Enterprise Sales Director"},
		{"Sales & Business Development", "B2B Key Account Manager"},
		{"Sales & Business Development", "Business Development Executive"},
		{"Sales & Business Development", "Digital Marketing Specialist"},
		{"Supply Chain & Operasional", "Head of Supply Chain"},
		{"Supply Chain & Operasional", "Warehouse & Inventory Manager"},
		{"Supply Chain & Operasional", "Procurement Specialist"},
		{"Supply Chain & Operasional", "Logistics & Fleet Supervisor"},
		{"Supply Chain & Operasional", "Quality Control Inspector"},
	}

	var createdJobs []*employees.JobPosition
	for _, j := range jobsList {
		deptID := depts[j.Dept].ID
		var pos employees.JobPosition
		config.DB.Where("name = ?", j.Name).FirstOrCreate(&pos, employees.JobPosition{
			Name:         j.Name,
			DepartmentID: &deptID,
			State:        "recruit",
			CreatedAt:    time.Now(),
		})
		createdJobs = append(createdJobs, &pos)
	}

	log.Println("👥 [3/6] Memulai Seeding 100+ Karyawan Indonesia Skala Enterprise...")

	firstNames := []string{
		"Budi", "Siti", "Ahmad", "Dewi", "Rizky", "Rina", "Bayu", "Sri", "Dimas", "Tri",
		"Eko", "Maya", "Fajar", "Indah", "Hendra", "Ratna", "Agus", "Wulan", "Bambang", "Putri",
		"Doni", "Nur", "Wahyu", "Lestari", "Ilham", "Dian", "Reza", "Sari", "Aditya", "Fitri",
		"Gilang", "Anisa", "Arif", "Mega", "Yusuf", "Kartika", "Farhan", "Nadia", "Fikri", "Ayu",
		"Surya", "Intan", "Satria", "Widya", "Rangga", "Tia", "Danang", "Melati", "Galih", "Kusuma",
	}

	lastNames := []string{
		"Pratama", "Santoso", "Wijaya", "Kusuma", "Saputra", "Hidayat", "Nugroho", "Wibowo", "Siregar", "Nasution",
		"Purnomo", "Setiawan", "Utomo", "Suryono", "Handayani", "Wahyudi", "Gunawan", "Susanto", "Lestari", "Firmansyah",
		"Mahendra", "Hakim", "Zulkarnain", "Kurniawan", "Suharto", "Mulyadi", "Iskandar", "Hartono", "Syahputra", "Subagyo",
	}

	ptkpOptions := []string{"TK/0", "TK/1", "K/0", "K/1", "K/2", "K/3"}

	topLevels := []struct {
		Name     string
		JobIdx   int
		Email    string
		DeptName string
	}{
		{"Dr. Hendra Pratama", 0, "ceo@nusantara-erp.id", "Direksi & Eksekutif"},
		{"Bambang Wicaksono, M.Kom", 1, "cto@nusantara-erp.id", "Direksi & Eksekutif"},
		{"Dewi Sartika, S.E., Ak.", 2, "cfo@nusantara-erp.id", "Direksi & Eksekutif"},
		{"Rina Sulistyowati, S.Psi", 3, "hr.director@nusantara-erp.id", "Human Resources & GA"},
		{"Aris Nugroho, M.T.", 7, "vp.eng@nusantara-erp.id", "Teknologi Informasi & Digital"},
		{"Agus Salim, M.Ak", 14, "finance.director@nusantara-erp.id", "Keuangan, Akuntansi & Pajak"},
		{"Doni Firmansyah, MBA", 19, "sales.director@nusantara-erp.id", "Sales & Business Development"},
		{"Ir. Surya Darmawan", 24, "operations.head@nusantara-erp.id", "Supply Chain & Operasional"},
	}

	var managers []*employees.Employee
	for i, m := range topLevels {
		job := createdJobs[m.JobIdx]
		dept := depts[m.DeptName]
		joinDate := time.Now().AddDate(-5, -rand.Intn(10), -rand.Intn(20))

		emp := employees.Employee{
			Name:             m.Name,
			DepartmentID:     &dept.ID,
			JobPositionID:    &job.ID,
			WorkEmail:        m.Email,
			WorkPhone:        fmt.Sprintf("0812%08d", 10000000+i),
			EmergencyContact: "Keluarga Inti",
			EmergencyPhone:   fmt.Sprintf("0813%08d", 20000000+i),
			PTKPStatus:       "K/2",
			JoinDate:         &joinDate,
			IsActive:         true,
			CreatedAt:        time.Now(),
		}
		config.DB.Create(&emp)
		managers = append(managers, &emp)
	}

	currentCount := len(managers)
	for currentCount < 105 {
		fn := firstNames[rand.Intn(len(firstNames))]
		ln := lastNames[rand.Intn(len(lastNames))]
		fullName := fmt.Sprintf("%s %s", fn, ln)

		jobIdx := 3 + rand.Intn(len(createdJobs)-3)
		job := createdJobs[jobIdx]
		deptID := *job.DepartmentID

		mgr := managers[rand.Intn(len(managers))]

		joinYearsAgo := rand.Intn(4) + 1
		joinDate := time.Now().AddDate(-joinYearsAgo, -rand.Intn(11), -rand.Intn(25))
		ptkp := ptkpOptions[rand.Intn(len(ptkpOptions))]
		email := fmt.Sprintf("%s.%s%d@nusantara-erp.id", fn, ln, rand.Intn(99))

		emp := employees.Employee{
			Name:             fullName,
			DepartmentID:     &deptID,
			JobPositionID:    &job.ID,
			ManagerID:        &mgr.ID,
			WorkEmail:        email,
			WorkPhone:        fmt.Sprintf("08%d%08d", rand.Intn(3)+1, rand.Intn(90000000)+10000000),
			EmergencyContact: "Keluarga",
			EmergencyPhone:   fmt.Sprintf("081%d%07d", rand.Intn(8)+1, rand.Intn(9000000)+1000000),
			PTKPStatus:       ptkp,
			JoinDate:         &joinDate,
			IsActive:         true,
			CreatedAt:        time.Now(),
		}
		config.DB.Create(&emp)
		currentCount++
	}
	log.Printf("✅ Berhasil men-generate %d Karyawan Skala Korporat!", currentCount)

	log.Println("📋 [4/6] Menyiapkan Pipeline Rekrutmen & Pelamar Kerja...")
	stages := []string{
		"1. Seleksi Berkas (Screening)",
		"2. Wawancara HR (HR Interview)",
		"3. Tes Teknis & User",
		"4. Negosiasi & Offering",
		"5. Diterima (Hired)",
	}

	var stageIDs []uint
	for idx, sName := range stages {
		var st recruitment.Stage
		config.DB.Where("name = ?", sName).FirstOrCreate(&st, recruitment.Stage{
			Name:     sName,
			Sequence: (idx + 1) * 10,
		})
		stageIDs = append(stageIDs, st.ID)
	}

	applicantNames := []string{
		"Guruh Wicaksono", "Nadia Safitri", "Kevin Sanjaya", "Clarissa Amanda",
		"Rendi Septian", "Farah Diba", "Aldi Taher", "Tiara Andini", "Rizki Febian",
	}

	for i, aName := range applicantNames {
		job := createdJobs[rand.Intn(len(createdJobs))]
		stID := stageIDs[rand.Intn(len(stageIDs))]
		salary := float64((rand.Intn(15) + 6) * 1000000)

		applicant := recruitment.Applicant{
			Name:           aName,
			Email:          fmt.Sprintf("applicant.%d@gmail.com", i+1),
			Phone:          fmt.Sprintf("0857%08d", rand.Intn(90000000)+10000000),
			JobPositionID:  job.ID,
			StageID:        stID,
			ExpectedSalary: salary,
			State:          "in_progress",
			CreatedAt:      time.Now().AddDate(0, 0, -rand.Intn(30)),
		}
		config.DB.Create(&applicant)
	}

	log.Println("📢 [5/6] Menyiapkan Berita Korporat & Pengumuman Intranet...")
	announcements := []website_builder.Announcement{
		{
			Title:     "Town Hall Q3 2026: Pencapaian Kinerja & Roadmap Digitalisasi ERP",
			Category:  "Corporate",
			Content:   "Seluruh karyawan diundang untuk menghadiri pertemuan Town Hall pada hari Jumat pukul 14.00 WIB secara hybrid di Auditorium Utama dan Zoom.",
			Author:    "Direksi Perusahaan",
			IsPinned:  true,
			CreatedAt: time.Now().AddDate(0, 0, -2),
		},
		{
			Title:     "Pembaruan Kebijakan Klaim Medis & BPJS Ketenagakerjaan 2026",
			Category:  "HR Policy",
			Content:   "Mulai periode September 2026, integrasi sistem klaim asuransi rawat jalan dapat langsung diajukan melalui Portal Karyawan ESS.",
			Author:    "Divisi HR & People Operations",
			IsPinned:  true,
			CreatedAt: time.Now().AddDate(0, 0, -5),
		},
		{
			Title:     "Pelatihan Wajib: Keamanan Siber & Perlindungan Data Pribadi (UU PDP)",
			Category:  "Training & Compliance",
			Content:   "Diberitahukan kepada seluruh karyawan untuk menyelesaikan modul pelatihan Cyber Security Awareness di menu Website e-Learning sebelum akhir bulan ini.",
			Author:    "IT Security & Compliance",
			IsPinned:  false,
			CreatedAt: time.Now().AddDate(0, 0, -10),
		},
		{
			Title:     "Pengumuman Pemenang Employee of the Month Agustus 2026",
			Category:  "Recognition",
			Content:   "Selamat kepada para rekan kerja berprestasi yang telah memberikan dedikasi terbaik bagi pertumbuhan perusahaan bulan lalu.",
			Author:    "HR Directorate",
			IsPinned:  false,
			CreatedAt: time.Now().AddDate(0, 0, -14),
		},
	}

	for _, ann := range announcements {
		config.DB.Create(&ann)
	}

	// Re-enable Foreign Key constraints
	config.DB.Exec("SET session_replication_role = 'origin';")

	log.Println("🎉 [6/6] ENTERPRISE SEEDING SELESAI DENGAN SUKSES!")
	log.Println("Struktur 100+ karyawan, 6 departemen, lowongan karir, dan pengumuman korporat telah siap.")
}
