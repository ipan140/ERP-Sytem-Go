package discuss

import (
	"ERP-System/config"
	"time"
)

type Channel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(100);default:''" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Type      string    `gorm:"type:varchar(50);default:'PUBLIC'" json:"type"` // PUBLIC, DIRECT
	Topic     string    `gorm:"type:varchar(255);default:''" json:"topic"`
	Members   string    `gorm:"type:varchar(100);default:'248 Anggota'" json:"members"`
	Unread    int       `gorm:"default:0" json:"unread"`
	Online    bool      `gorm:"default:false" json:"online"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Channel) TableName() string {
	return "setting.channels"
}

type DiscussMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ChannelID string    `gorm:"type:varchar(100);not null;index" json:"channel_id"`
	Sender    string    `gorm:"type:varchar(255);not null" json:"sender"`
	Text      string    `gorm:"type:text;not null" json:"text"`
	IsMe      bool      `gorm:"default:false" json:"is_me"`
	CreatedAt time.Time `json:"created_at"`
}

func (DiscussMessage) TableName() string {
	return "setting.discuss_messages"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Channel{}, &DiscussMessage{})
}
