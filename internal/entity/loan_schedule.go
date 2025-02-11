package entity

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type ILoanScheduleRepository interface {
	BulkCreate(c context.Context, tx *gorm.DB, loanSchedule []LoanSchedule) error
	GetLoanSchedulesByLoanID(c context.Context, loanID uint) ([]LoanSchedule, error)
	GetByLoanIDAndWeekNumberAndStatusPaid(c context.Context, loanID uint, weekNumber int) (*LoanSchedule, error)
	GetHasOverDuePayment(c context.Context, loanID uint) bool
	GetDelinquent(c context.Context, loanID uint) ([]uint, error)
	UpdateToPaidByIDs(c context.Context, tx *gorm.DB, ids []uint) error
	UpdateToPaidByLoanIDAndWeekNumber(c context.Context, tx *gorm.DB, loanID uint, weekNumber int) error
}

type LoanSchedule struct {
	ID         uint      `gorm:"primary_key;auto_increment" json:"id"`
	LoanID     uint      `gorm:"primary_key;auto_increment" json:"loan_id"`
	WeekNumber int       `json:"week_number"`
	DueDate    time.Time `json:"due_date"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
