package transactioner

import (
	"context"
	"gorm.io/gorm"
)

type GormTransactioner interface {
	Begin(context context.Context) *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB)
}
