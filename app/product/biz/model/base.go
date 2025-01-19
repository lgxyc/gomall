package model

import "time"

type Base struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
