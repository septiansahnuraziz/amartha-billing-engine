package repository

import (
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/utils"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type LoanScheduleRepository struct {
	db *gorm.DB
}

func NewLoanScheduleRepository(db *gorm.DB) entity.ILoanScheduleRepository {
	return &LoanScheduleRepository{
		db: db,
	}
}

func (r *LoanScheduleRepository) BulkCreate(c context.Context, tx *gorm.DB, loanSchedule []entity.LoanSchedule) error {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx":          utils.DumpIncomingContext(c),
		"loanSchedule": utils.Dump(loanSchedule),
	})

	if err := tx.WithContext(c).Debug().Create(&loanSchedule).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *LoanScheduleRepository) GetByLoanIDAndWeekNumberAndStatusPaid(c context.Context, loanID uint, weekNumber int) (*entity.LoanSchedule, error) {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx":        utils.DumpIncomingContext(c),
		"loanID":     loanID,
		"weekNumber": weekNumber,
	})

	var loanSchedule entity.LoanSchedule
	if err := r.db.Where("loan_id = ? AND week_number = ? AND status = 'pending'", loanID, weekNumber).First(&loanSchedule).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		logger.Error(err)

		return nil, err
	}

	return &loanSchedule, nil
}

func (r *LoanScheduleRepository) GetLoanSchedulesByLoanID(c context.Context, loanID uint) ([]entity.LoanSchedule, error) {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx":    utils.DumpIncomingContext(c),
		"loanID": loanID,
	})

	var loanSchedule []entity.LoanSchedule
	if err := r.db.Where("loan_id = ?", loanID).Order("id ASC").Find(&loanSchedule).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return loanSchedule, nil
}

func (r *LoanScheduleRepository) GetHasOverDuePayment(c context.Context, loanID uint) bool {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx":    utils.DumpIncomingContext(c),
		"loanID": loanID,
	})

	var overdueCount int64
	if err := r.db.Model(&entity.LoanSchedule{}).Where("loan_id = ? AND due_date < ? AND status='pending'", loanID, time.Now()).Count(&overdueCount).Error; err != nil {
		logger.Error(err)
		return true
	}

	if overdueCount > 0 {

		return true
	}
	return false
}

func (r *LoanScheduleRepository) GetDelinquent(c context.Context, loanID uint) ([]uint, error) {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx":    utils.DumpIncomingContext(c),
		"loanID": loanID,
	})

	var ids []uint
	if err := r.db.Model(&entity.LoanSchedule{}).Where("loan_id = ? AND due_date < ?", loanID, time.Now()).Limit(2).Pluck("id", &ids).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return ids, nil
}

func (r *LoanScheduleRepository) UpdateToPaidByIDs(c context.Context, tx *gorm.DB, ids []uint) error {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(c),
		"ids": ids,
	})

	if err := tx.WithContext(c).Model(entity.LoanSchedule{}).Where("id IN ?", ids).UpdateColumn("status", "paid").Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *LoanScheduleRepository) UpdateToPaidByLoanIDAndWeekNumber(c context.Context, tx *gorm.DB, loanID uint, weekNumber int) error {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"ctx":        utils.DumpIncomingContext(c),
		"loanID":     loanID,
		"weekNumber": weekNumber,
	})

	if err := tx.WithContext(c).Model(entity.LoanSchedule{}).Where("loan_id = ? AND week_number = ?", loanID, weekNumber).UpdateColumn("status", "paid").Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
