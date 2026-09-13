package calendar

import (
	"ERP-System/app/modules/auth"
	"ERP-System/config"
)

func FindEvents(userID uint, resModel string, start string, end string) ([]CalendarEvent, error) {
	var events []CalendarEvent
	query := config.DB.Model(&CalendarEvent{}).Preload("User").Preload("Attendees")

	if resModel != "" {
		query = query.Where("res_model = ?", resModel)
	}

	query = query.Where(
		config.DB.Where("visibility = ?", "public").
			Or("user_id = ?", userID).
			Or("id IN (SELECT calendar_event_id FROM calendar_event_attendees WHERE user_id = ?)", userID),
	)

	if start != "" && end != "" {
		query = query.Where("start >= ? AND start <= ?", start, end)
	}

	err := query.Find(&events).Error
	return events, err
}

func CreateEventRepo(event *CalendarEvent) error {
	if len(event.AttendeeIDs) > 0 {
		var users []*auth.User
		config.DB.Where("id IN ?", event.AttendeeIDs).Find(&users)
		event.Attendees = users
	}
	return config.DB.Create(event).Error
}

func GetEventByID(id string) (CalendarEvent, error) {
	var event CalendarEvent
	err := config.DB.First(&event, id).Error
	return event, err
}

func UpdateEventRepo(event *CalendarEvent) error {
	if len(event.AttendeeIDs) > 0 {
		var users []*auth.User
		config.DB.Where("id IN ?", event.AttendeeIDs).Find(&users)
		config.DB.Model(event).Association("Attendees").Replace(users)
	}
	return config.DB.Save(event).Error
}

func DeleteEventRepo(id string) error {
	return config.DB.Delete(&CalendarEvent{}, id).Error
}

func FindCategories(module string) ([]CalendarCategory, error) {
	var categories []CalendarCategory
	query := config.DB.Model(&CalendarCategory{})
	if module != "" {
		query = query.Where("module = ?", module)
	}
	err := query.Order("id ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	// Auto-seed default categories if empty for this module
	if len(categories) == 0 && module != "" {
		defaults := getDefaultCategoriesForModule(module)
		for i := range defaults {
			_ = config.DB.Create(&defaults[i]).Error
		}
		_ = query.Order("id ASC").Find(&categories).Error
	}

	return categories, nil
}

func getDefaultCategoriesForModule(module string) []CalendarCategory {
	switch module {
	case "HR":
		return []CalendarCategory{
			{Module: "HR", Name: "Cuti Staf", Color: "primary", Icon: "🏖️", Description: "Pengajuan dan persetujuan cuti tahunan staf"},
			{Module: "HR", Name: "Cut-off Payroll", Color: "danger", Icon: "💰", Description: "Batas akhir rekap absensi & hitung gaji"},
			{Module: "HR", Name: "Wawancara", Color: "warning", Icon: "📋", Description: "Jadwal interview pelamar rekrutmen"},
			{Module: "HR", Name: "Ulang Tahun", Color: "success", Icon: "🎂", Description: "Peringatan hari ulang tahun karyawan"},
		}
	case "Finance":
		return []CalendarCategory{
			{Module: "Finance", Name: "Pajak PPh / PPN", Color: "danger", Icon: "🧾", Description: "Batas pembayaran & pelaporan SPT Masa DJP"},
			{Module: "Finance", Name: "Jatuh Tempo Piutang (AR)", Color: "primary", Icon: "🏦", Description: "Jatuh tempo pembayaran faktur dari customer"},
			{Module: "Finance", Name: "Hutang Vendor (AP)", Color: "warning", Icon: "💸", Description: "Jadwal pelunasan tagihan supplier vendor"},
			{Module: "Finance", Name: "Tutup Buku & Rekonsiliasi", Color: "purple", Icon: "📊", Description: "Penutupan buku bulanan & rekonsiliasi bank"},
		}
	case "SupplyChain":
		return []CalendarCategory{
			{Module: "SupplyChain", Name: "ETA Supplier (PO)", Color: "warning", Icon: "🚚", Description: "Estimasi waktu tiba kiriman bahan baku"},
			{Module: "SupplyChain", Name: "Pengiriman (DO)", Color: "primary", Icon: "📦", Description: "Jadwal dispatch armada pengiriman barang"},
			{Module: "SupplyChain", Name: "Stock Opname Gudang", Color: "danger", Icon: "🔍", Description: "Penghitungan fisik stok barang gudang"},
			{Module: "SupplyChain", Name: "Perawatan Mesin", Color: "success", Icon: "🛠️", Description: "Jadwal servis preventif mesin pabrik"},
		}
	case "Sales":
		return []CalendarCategory{
			{Module: "Sales", Name: "Demo Presentasi B2B", Color: "primary", Icon: "🤝", Description: "Presentasi produk ke calon klien korporasi"},
			{Module: "Sales", Name: "Follow-up Quotation", Color: "warning", Icon: "📞", Description: "Follow-up negosiasi penawaran harga"},
			{Module: "Sales", Name: "Signing Kontrak MOU", Color: "success", Icon: "✍️", Description: "Penandatanganan kerja sama & closing deal"},
			{Module: "Sales", Name: "Review Target Kuota", Color: "danger", Icon: "🎯", Description: "Rapat evaluasi kuota bulanan tim sales"},
		}
	case "Marketing":
		return []CalendarCategory{
			{Module: "Marketing", Name: "Email Newsletter Blast", Color: "primary", Icon: "📧", Description: "Siaran email buletin promosi massal"},
			{Module: "Marketing", Name: "Konten Media Sosial", Color: "purple", Icon: "📱", Description: "Jadwal posting konten feed / reels / video"},
			{Module: "Marketing", Name: "Webinar Edukasi", Color: "warning", Icon: "🎪", Description: "Penyelenggaraan workshop atau webinar live"},
			{Module: "Marketing", Name: "Promo & Diskon", Color: "danger", Icon: "🏷️", Description: "Peluncuran kampanye diskon & kupon voucher"},
		}
	case "Services":
		return []CalendarCategory{
			{Module: "Services", Name: "Kunjungan Lapangan", Color: "warning", Icon: "🛠️", Description: "Kunjungan teknisi ke lokasi customer on-site"},
			{Module: "Services", Name: "Sprint Planning", Color: "primary", Icon: "⏳", Description: "Sesi perencanaan sprint kerja tim proyek"},
			{Module: "Services", Name: "Milestone & Delivery", Color: "success", Icon: "🎯", Description: "Pencapaian target fase serah terima proyek"},
			{Module: "Services", Name: "Evaluasi Tiket SLA", Color: "danger", Icon: "🎫", Description: "Peninjauan tiket insiden helpdesk prioritas tinggi"},
		}
	case "Website":
		return []CalendarCategory{
			{Module: "Website", Name: "Artikel Blog SEO", Color: "primary", Icon: "📝", Description: "Jadwal terbit artikel blog dan optimasi SEO"},
			{Module: "Website", Name: "Mentoring Live LMS", Color: "purple", Icon: "🎓", Description: "Sesi mentoring dan kelas live interaktif LMS"},
			{Module: "Website", Name: "Demo Konsultasi B2B", Color: "success", Icon: "🤝", Description: "Sesi appointment konsultasi yang dibooking user"},
			{Module: "Website", Name: "Rilis Fitur Web", Color: "danger", Icon: "🚀", Description: "Deployment pembaruan portal web ke server"},
		}
	default:
		return []CalendarCategory{
			{Module: module, Name: "Umum", Color: "primary", Icon: "📌", Description: "Kategori umum"},
		}
	}
}

func CreateCategoryRepo(cat *CalendarCategory) error {
	return config.DB.Create(cat).Error
}

func GetCategoryByID(id string) (CalendarCategory, error) {
	var cat CalendarCategory
	err := config.DB.First(&cat, id).Error
	return cat, err
}

func UpdateCategoryRepo(cat *CalendarCategory) error {
	return config.DB.Save(cat).Error
}

func DeleteCategoryRepo(id string) error {
	return config.DB.Delete(&CalendarCategory{}, id).Error
}
