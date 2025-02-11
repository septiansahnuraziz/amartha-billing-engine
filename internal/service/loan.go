package service

import (
	"amartha-billing-engine/internal/database/transactioner"
	"amartha-billing-engine/internal/entity"
	"amartha-billing-engine/utils"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

type LoanService struct {
	LoanRepository  entity.ILoanRepository
	LoanSchedule    entity.ILoanScheduleRepository
	gormTransaction transactioner.GormTransactioner
}

func NewLoanService(LoanRepository entity.ILoanRepository, LoanSchedule entity.ILoanScheduleRepository, gormTransaction transactioner.GormTransactioner) entity.ILoanService {
	return &LoanService{
		LoanRepository:  LoanRepository,
		LoanSchedule:    LoanSchedule,
		gormTransaction: gormTransaction,
	}
}

func (s *LoanService) CreateLoan(c context.Context, request entity.RequestCreateLoan) error {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"context": utils.DumpIncomingContext(c),
		"request": utils.Dump(request),
	})

	// Begin transaction
	tx := s.gormTransaction.Begin(c)

	// Do create loan
	loan, err := s.LoanRepository.Create(c, tx, request.ToLoanModel())
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	// Do create loan schecule,
	var loanSchedules []entity.LoanSchedule
	//loop and mapping data loan schedule base on tenor
	for i := 0; i < request.Tenor; i++ {
		dueDate := time.Now().AddDate(0, 0, (i+1)*7)
		loanSchedules = append(loanSchedules, request.ToLoanSchedules(*loan, i+1, dueDate))
	}

	//create loan schedule
	if err := s.LoanSchedule.BulkCreate(c, tx, loanSchedules); err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	//commit transaction
	tx.Commit()

	logger.Info("Loan created")
	return nil
}

func (s *LoanService) GetLoanDetail(c context.Context, loanID uint) (entity.ResponseGetLoanDetail, error) {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"context": utils.DumpIncomingContext(c),
		"loanID":  loanID,
	})

	var response entity.ResponseGetLoanDetail

	//	get loan
	existLoan, err := s.LoanRepository.GetByID(c, loanID)
	if err != nil {
		logger.Error(err)
	}

	//validate if loan doesnt exist
	if existLoan == nil {
		return entity.ResponseGetLoanDetail{}, errors.New("loan not found")
	}

	// get loan schedule data
	existLoanSchedule, err := s.LoanSchedule.GetLoanSchedulesByLoanID(c, existLoan.ID)
	if err != nil {
		logger.Error(err)
		return entity.ResponseGetLoanDetail{}, err
	}

	//mapping data and return result
	return response.FromLoanModel(*existLoan, existLoanSchedule), nil
}

// this function for pay the loan
func (s *LoanService) PayLoan(c context.Context, request entity.RequestPayLoan) error {
	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"context": utils.DumpIncomingContext(c),
		"request": utils.Dump(request),
	})

	//get loan by id loan
	existLoan, err := s.LoanRepository.GetByID(c, request.LoanID)
	if err != nil {
		logger.Error(err)
		return err
	}

	if existLoan == nil {
		return errors.New("loan not found")
	}

	tx := s.gormTransaction.Begin(c)

	// validate loan if has overdue
	hasOverdue := s.LoanSchedule.GetHasOverDuePayment(c, existLoan.ID)
	if hasOverdue {
		//do makesure the laon is deliqunt or not
		getDelinqunt, err := s.LoanSchedule.GetDelinquent(c, existLoan.ID)
		if err != nil {
			logger.Error(err)
			return err
		}

		//validate if loan delinqunt, and then borrower must pay based on his arrears
		if getDelinqunt != nil && len(getDelinqunt) == 2 {
			// do update status to paid by ids, which is loan schedule where delinqunt
			if err := s.LoanSchedule.UpdateToPaidByIDs(c, tx, getDelinqunt); err != nil {
				logger.Error(err)
				return err
			}

			// do decrease outstanding amount on loan
			if err := s.LoanRepository.DecreaseOutstandingAmount(c, tx, existLoan.ID, int(existLoan.OutstandingAmount-request.Amount)); err != nil {
				logger.Error(err)
				return err
			}

			// commit transaction
			tx.Commit()

			return nil
		}

		return errors.New("loan has overdue")
	}

	//validate amount in accordance with weekly payment
	if request.Amount != existLoan.WeeklyPayment {
		return errors.New("invalid amount")
	}

	//	get current week billing
	weekNumber, _ := utils.GetCurrentWeekBilling(existLoan.CreatedAt, existLoan.TotalWeeks)

	//validate belling it was paid
	existLoanSchedule, err := s.LoanSchedule.GetByLoanIDAndWeekNumberAndStatusPaid(c, existLoan.ID, weekNumber)
	if err != nil {
		logger.Error(err)
		return err
	}

	if existLoanSchedule == nil {
		return errors.New("your billing was paid in this week")
	}

	// do update status to paid by loan ID and week number
	if err := s.LoanSchedule.UpdateToPaidByLoanIDAndWeekNumber(c, tx, existLoan.ID, weekNumber); err != nil {
		logger.Error(err)
		return err
	}

	// do decrease outstanding amount on loan
	if err := s.LoanRepository.DecreaseOutstandingAmount(c, tx, existLoan.ID, int(existLoan.OutstandingAmount-request.Amount)); err != nil {
		logger.Error(err)
		return err
	}

	//commit transaction
	tx.Commit()

	return nil
}
