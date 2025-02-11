package repository

import (
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/utils"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) entity.ILoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (r *LoanRepository) Create(c context.Context, tx *gorm.DB, loan *entity.Loan) (*entity.Loan, error) {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx":  utils.DumpIncomingContext(c),
		"loan": utils.Dump(loan),
	})

	if err := tx.WithContext(c).Debug().Save(&loan).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return loan, nil
}

func (r *LoanRepository) GetByID(c context.Context, id uint) (*entity.Loan, error) {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(c),
		"id":  id,
	})

	var loan entity.Loan
	if err := r.db.Where("id = ?", id).First(&loan).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return &loan, nil
}

func (r *LoanRepository) DecreaseOutstandingAmount(c context.Context, tx *gorm.DB, id uint, amount int) error {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(c),
		"id":  id,
	})

	if err := tx.WithContext(c).Model(entity.Loan{}).Where("id = ?", id).UpdateColumn("outstanding_amount", amount).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
