package global

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	CreateAt time.Time      `json:"create_at"`
	UpdateAt time.Time      `json:"update_at" gorm:"column:update_at;autoUpdateTime"`
	DeleteAt gorm.DeletedAt `json:"-" gorm:"index"`
}
