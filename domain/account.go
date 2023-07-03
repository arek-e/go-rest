// Package domain for account
package domain

import (
	"time"

	"gorm.io/gorm"
)

type Currency string

const (
	Euro      Currency = "Euro (€)"
	US_Dollar Currency = "US Dollar ($)"
	SEK       Currency = "SEK (kr)"
)

type Account struct {
	gorm.Model

	ID                string     `gorm:"primary_key"`
	Name              string     `gorm:"not null"`
	Currency          Currency   `gorm:"not null;default:'SEK (kr)'"`
	AccountType       *string    `gorm:"default:null"`
	AccountNumber     *string    `gorm:"default:null"`
	StartBalance      *float64   `gorm:"type:decimal(10,2);default:0"`
	StartBalanceDate  *time.Time `gorm:"default:current_timestamp"`
	CurrentBalance    float64    `gorm:"default:0"`
	IsActive          bool       `gorm:"default:true"`
	LastActivity      time.Time  `gorm:"default:current_timestamp"`
	BalanceDifference float64    `gorm:"default:0"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	if a.StartBalance != nil {
		a.CurrentBalance = *a.StartBalance
	}

	return nil
}
