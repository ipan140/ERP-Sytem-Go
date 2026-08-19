package social_marketing

import "ERP-System/config"

// 7. Social Marketing Manager
type SocialPost struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Message         string `gorm:"type:text" json:"message"`
	PostToFacebook  bool   `gorm:"default:false" json:"post_to_facebook"`
	PostToTwitter   bool   `gorm:"default:false" json:"post_to_twitter"`
	PostToInstagram bool   `gorm:"default:false" json:"post_to_instagram"`
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &SocialPost{}) }
