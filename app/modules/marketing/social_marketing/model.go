package social_marketing

import (
	"time"

	"ERP-System/config"
)

// 7. Social Marketing Manager
type SocialPost struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	Message         string     `gorm:"type:text" json:"message"`
	ImageURL        string     `gorm:"type:text" json:"image_url"` // URL Gambar atau Base64 Upload Data
	Channels        string     `gorm:"type:text" json:"channels"`  // Format list kanal dinamis: "Facebook,Instagram,TikTok,Pinterest,WhatsApp Channel"
	PostToFacebook  bool       `gorm:"default:false" json:"post_to_facebook"`
	PostToTwitter   bool       `gorm:"default:false" json:"post_to_twitter"`
	PostToInstagram bool       `gorm:"default:false" json:"post_to_instagram"`
	PostToLinkedin  bool       `gorm:"default:false" json:"post_to_linkedin"`
	Status          string     `gorm:"type:varchar(50);default:'Draft'" json:"status"` // Draft, In-Review, Scheduled, Published
	ScheduledAt     *time.Time `json:"scheduled_at"`
	ReachCount      int        `gorm:"default:0" json:"reach_count"`
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &SocialPost{}) }


// ---- Auto-Generated TableName methods ----
func (SocialPost) TableName() string {
	return "marketing.social_posts"
}