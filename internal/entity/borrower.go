package entity

import "time"

type IBorrower interface {
	//GetByID(id string) (*Borrower, error)
	GetByEmail(email string) (*Borrower, error)
}

type Borrower struct {
	ID        uint64 `gorm:"primary_key;auto_increment"`
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}
