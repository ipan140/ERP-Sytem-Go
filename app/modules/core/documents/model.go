package documents

import (
	"ERP-System/config"
	"time"
)

type Workspace struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}


func (Workspace) TableName() string {
	return "setting.workspaces"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Workspace{})
}
