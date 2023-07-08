package entities

import "time"

type Users struct {
	ID        int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username  string `gorm:"column:username;unique;not null;type:varchar(36);default:null"`
	Password  string `gorm:"column:password;unique;not null;type:varchar(36);default:null"`
	Role      string `gorm:"column:role;unique;not null;type:enum('USER', 'ADMIN')"`
	IsActive  bool   `gorm:"column:is_active;not null;default:true;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (h *Users) TableName() string {
	return "users"
}
