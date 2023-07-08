package entities

import "time"

type Users struct {
	ID int `gorm:"column:user_id;primaryKey;autoIncrement"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (h *Users) TableName() string {
	return "users"
}
