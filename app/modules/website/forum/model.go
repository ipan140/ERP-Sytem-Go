package forum

import (
	"ERP-System/app/modules/core/base"
	"ERP-System/config"
	"time"
)

type ForumPost struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	ForumID   uint          `json:"forum_id"`
	AuthorID  uint          `json:"author_id"`
	Author    *base.Partner `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Name      string        `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time     `json:"created_at"`
}

func (ForumPost) TableName() string {
	return "website_portal.forum_posts"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ForumPost{})
}
