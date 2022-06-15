package model

import (
	"github.com/google/uuid"
	"time"
)

type Test struct {
	ID          string     `gorm:"column:ID;type:char(36);primary_key"`
	CreatedAt   time.Time  ``
	UpdatedAt   time.Time  ``
	DeletedAt   *time.Time `sql:"index"`
	CreatedBy   uuid.UUID  `gorm:"column:created_by;type:char(36)"`
	CreatorName string     `gorm:"column:creator_name;type:varchar(40)"`
	Check       string     `gorm:"column:check;index"`
}

// TableName 테이블이름
func (Test) TableName() string {
	return "TEST"
}
