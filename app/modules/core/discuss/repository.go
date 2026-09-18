package discuss

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateChannel(data *Channel) error {
	return config.DB.Create(data).Error
}

func GetAllChannel() ([]Channel, error) {
	var list []Channel
	err := config.DB.Preload(clause.Associations).Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultChannels()
		_ = config.DB.Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultChannels() error {
	defaults := []Channel{
		{Code: "ch_general", Name: "general", Type: "PUBLIC", Members: "248 Anggota", Topic: "Pengumuman umum seluruh karyawan perusahaan", Unread: 0},
		{Code: "ch_hr", Name: "pengumuman-hrd", Type: "PUBLIC", Members: "248 Anggota", Topic: "Informasi cuti bersama, benefit BPJS, dan slip gaji", Unread: 2},
		{Code: "ch_finance", Name: "finance-budgeting", Type: "PUBLIC", Members: "18 Anggota", Topic: "Koordinasi penutupan buku akhir bulan & perpajakan", Unread: 0},
		{Code: "ch_warehouse", Name: "warehouse-cikarang", Type: "PUBLIC", Members: "34 Anggota", Topic: "Laporan barang masuk inbound & mutasi rak", Unread: 0},
		{Code: "ch_sales", Name: "sales-champions", Type: "PUBLIC", Members: "52 Anggota", Topic: "Target closing kuartal III & diskon approval", Unread: 1},
		{Code: "dm_hr_lead", Name: "Siti Aminah (HR Specialist)", Type: "DIRECT", Members: "Direct Message", Topic: "Obrolan pribadi", Online: true},
		{Code: "dm_finance_head", Name: "Budi Santoso (Finance Manager)", Type: "DIRECT", Members: "Direct Message", Topic: "Obrolan pribadi", Online: true},
		{Code: "dm_director", Name: "Ahmad Fauzi (Direktur)", Type: "DIRECT", Members: "Direct Message", Topic: "Obrolan pribadi", Online: false},
	}
	for _, ch := range defaults {
		var existing Channel
		if err := config.DB.Where("code = ?", ch.Code).First(&existing).Error; err != nil {
			_ = config.DB.Create(&ch).Error
		}
	}
	return nil
}

func GetMessagesByChannel(channelID string) ([]DiscussMessage, error) {
	var list []DiscussMessage
	err := config.DB.Where("channel_id = ?", channelID).Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 && channelID == "ch_general" {
		_ = SeedDefaultMessages()
		_ = config.DB.Where("channel_id = ?", channelID).Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultMessages() error {
	defaults := []DiscussMessage{
		{ChannelID: "ch_general", Sender: "Ahmad Fauzi (Direktur)", Text: "Selamat pagi rekan-rekan. Terima kasih atas kerja keras seluruh tim pada kuartal ini.", IsMe: false},
		{ChannelID: "ch_general", Sender: "HR Specialist", Text: "Pengingat: Batas akhir klaim pengobatan bulan September adalah besok pukul 17:00 WIB.", IsMe: false},
		{ChannelID: "ch_general", Sender: "Saya (Superadmin)", Text: "Server sistem ERP telah diperbarui ke versi Enterprise v2.5.0. Semua modul normal.", IsMe: true},
		{ChannelID: "ch_hr", Sender: "HR Specialist", Text: "Batch slip gaji periode September 2026 telah diposting dan dikirimkan via email/WhatsApp.", IsMe: false},
		{ChannelID: "ch_finance", Sender: "Finance Manager", Text: "Rekonsiliasi bank BCA dan Mandiri sudah sinkron 100%. File e-Faktur PPN siap lapor.", IsMe: false},
	}
	for _, m := range defaults {
		_ = config.DB.Create(&m).Error
	}
	return nil
}

func CreateDiscussMessage(data *DiscussMessage) error {
	return config.DB.Create(data).Error
}

func GetChannelByID(id uint) (*Channel, error) {
	var data Channel
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateChannel(data *Channel) error {
	return config.DB.Save(data).Error
}

func DeleteChannel(id uint) error {
	return config.DB.Delete(&Channel{}, id).Error
}
