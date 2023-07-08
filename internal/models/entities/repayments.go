package entities

import (
	"time"
)

type Repayments struct {
	ID        int     `gorm:"column:repayment_id;primaryKey;autoIncrement"`
	LoanID    int     `gorm:"column:loan_id;"`
	Amount    float64 `gorm:"column:amount;"`
	Status    string  `gorm:"column:status;type:varchar(64)"`
	UsersID   int     `gorm:"column:users_id;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Repayments) TableName() string {
	return "repayments"
}
