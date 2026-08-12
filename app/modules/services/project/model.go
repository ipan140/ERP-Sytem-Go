package project

import (
	"ERP-System/config"
	"time"
)

type Project struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	ManagerID  uint      `json:"manager_id"`  // Employee ID
	CustomerID uint      `json:"customer_id"` // Partner ID
	State      string    `gorm:"type:varchar(50);default:'active'" json:"state"` // active, done, cancelled
	CreatedAt  time.Time `json:"created_at"`
}

type Task struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ProjectID  uint      `json:"project_id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	AssigneeID uint      `json:"assignee_id"` // Employee ID
	Deadline   time.Time `json:"deadline"`
	Stage      string    `gorm:"type:varchar(50);default:'todo'" json:"stage"` // todo, in_progress, done
	CreatedAt  time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Project{}, &Task{})
}
