package entities

import "time"

type Users struct {
	ID        int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username  string `gorm:"column:username;unique;type:varchar(36)"`
	Password  string `gorm:"column:password;type:varchar(36);"`
	Role      string `gorm:"column:role;type:varchar(36)"`
	IsActive  bool   `gorm:"column:is_active;default:false;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (h *Users) TableName() string {
	return "users"
}
