package seed

import (
	"github.com/jinzhu/gorm"
	"errors"
	"fmt"
)

type SeedFunc func(*gorm.DB)error;

var seeds = []SeedFunc{
	userseeds,
	//reviewseeds,
	courseseeds,
}

func RunTx(DB *gorm.DB) error {
	tx := DB.Begin()
	defer func() {
    if r := recover(); r != nil {
			tx.Rollback()
    }
  }()

  if tx.Error != nil {
		return errors.New("failed to start transactions")
	}

	for _, s := range seeds {
		if err := s(tx); err != nil {
			return errors.New(fmt.Sprintf("Transaction failed: %v", err))
		}
	}
	
	return tx.Commit().Error
}