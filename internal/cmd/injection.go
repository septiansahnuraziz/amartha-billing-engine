package cmd

import (
	"amartha-billing-engine/internal/database/transactioner"
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/internal/repository"
	"amartha-billing-engine/internal/service"
	"gorm.io/gorm"
)

func InitLoan(db *gorm.DB, gormTransactioner transactioner.GormTransactioner) entity.ILoanService {
	loanRepository := repository.NewLoanRepository(db)
	loanScheduleRepository := repository.NewLoanScheduleRepository(db)
	return service.NewLoanService(loanRepository, loanScheduleRepository, gormTransactioner)
}
