package blog

import (
	"ERP-System/app/modules/auth"
	"ERP-System/config"
	"time"
)

type BlogPost struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Title         string     `gorm:"type:varchar(255)" json:"title"`
	Content       string     `gorm:"type:text" json:"content"`
	AuthorID      uint       `json:"author_id"`
	Author        *auth.User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	State         string     `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, published
	PublishedDate *time.Time `json:"published_date"`
	CreatedAt     time.Time  `json:"created_at"`
}

func (BlogPost) TableName() string {
	return "website_portal.blog_posts"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &BlogPost{})
}
