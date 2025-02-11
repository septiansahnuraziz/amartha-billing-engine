package entity

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type ILoanService interface {
	CreateLoan(ctx context.Context, request RequestCreateLoan) error
	PayLoan(c context.Context, request RequestPayLoan) error
	GetLoanDetail(c context.Context, loanID uint) (ResponseGetLoanDetail, error)
}

type ILoanRepository interface {
	Create(c context.Context, tx *gorm.DB, loan *Loan) (*Loan, error)
	GetByID(c context.Context, id uint) (*Loan, error)
	DecreaseOutstandingAmount(c context.Context, tx *gorm.DB, id uint, amount int) error
}

type Loan struct {
	ID                uint
	BorrowerId        uint
	LoanAmount        float64
	InterestRate      float64
	TotalAmount       float64
	TotalWeeks        int
	WeeklyPayment     float64
	OutstandingAmount float64
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type RequestCreateLoan struct {
	BorrowerID uint    `json:"borrowerId"`
	LoanAmount float64 `json:"loanAmount"`
	Tenor      int     `json:"tenor"`
	Interest   float64 `json:"interest"`
}

func (request *RequestCreateLoan) ToLoanModel() *Loan {
	var loan Loan

	if request.Tenor == 0 {
		request.Tenor = 50
	}

	interest := request.LoanAmount * (request.Interest / 100)
	totalAmount := request.LoanAmount + interest

	loan.BorrowerId = request.BorrowerID
	loan.LoanAmount = request.LoanAmount
	loan.InterestRate = request.Interest
	loan.WeeklyPayment = totalAmount / float64(request.Tenor)
	loan.TotalAmount = totalAmount
	loan.OutstandingAmount = totalAmount
	loan.TotalWeeks = request.Tenor
	loan.Status = "pending"

	return &loan
}

func (request *RequestCreateLoan) ToLoanSchedules(loan Loan, week int, dueDate time.Time) (loanSchedule LoanSchedule) {
	loanSchedule.LoanID = loan.ID
	loanSchedule.WeekNumber = week
	loanSchedule.Amount = loan.WeeklyPayment
	loanSchedule.DueDate = dueDate
	loanSchedule.Status = "pending"

	return loanSchedule
}

type RequestPayLoan struct {
	LoanID uint    `json:"loanId"`
	Amount float64 `json:"amount"`
}

type ResponseGetLoanDetail struct {
	ID           uint           `json:"id"`
	BorrowerName string         `json:"borrowerName"`
	LoanAmount   float64        `json:"loanAmount"`
	InterestRate float64        `json:"interestRate"`
	TotalAmount  float64        `json:"totalAmount"`
	Status       string         `json:"status"`
	LoanSchedule []LoanSchedule `json:"loanSchedule"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

func (response *ResponseGetLoanDetail) FromLoanModel(loan Loan, loanSchedule []LoanSchedule) ResponseGetLoanDetail {
	response.ID = loan.ID
	response.BorrowerName = "Agus"
	response.LoanAmount = loan.LoanAmount
	response.InterestRate = loan.InterestRate
	response.TotalAmount = loan.TotalAmount
	response.Status = loan.Status
	response.LoanSchedule = loanSchedule
	return *response
}
