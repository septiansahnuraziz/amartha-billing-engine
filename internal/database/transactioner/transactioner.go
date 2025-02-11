package transactioner

import (
	"context"
	"gorm.io/gorm"
)

type gormTransactioner struct {
	db *gorm.DB
}

func NewGormTransactioner(db *gorm.DB) GormTransactioner {
	return &gormTransactioner{db: db}
}

func (gormTransactioner *gormTransactioner) Begin(context context.Context) *gorm.DB {
	return gormTransactioner.db.WithContext(context).Begin()
}

func (gormTransactioner *gormTransactioner) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (gormTransactioner *gormTransactioner) Rollback(tx *gorm.DB) {
	tx.Rollback()
}
