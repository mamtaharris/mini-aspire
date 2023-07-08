package entities

import (
	"time"
)

type Loans struct {
	ID        int     `gorm:"column:loan_id;primaryKey;autoIncrement"`
	PID       string  `gorm:"column:loan_pid;unique;type:varchar(64)"`
	Amount    float64 `gorm:"column:amount;"`
	Term      int     `gorm:"column:term;"`
	Status    string  `gorm:"column:status;type:varchar(64)"`
	UsersID   int     `gorm:"column:users_id;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Loans) TableName() string {
	return "loans"
}
