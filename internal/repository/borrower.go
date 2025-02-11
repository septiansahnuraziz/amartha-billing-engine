package repository

import "gorm.io/gorm"

type BorrowerRepository struct {
	db *gorm.DB
}

func NewBorrowerRepository(db *gorm.DB) *BorrowerRepository {
	return &BorrowerRepository{
		db: db,
	}
}
