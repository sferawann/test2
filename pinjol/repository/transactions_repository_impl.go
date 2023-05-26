package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/pinjol/model"
	"gorm.io/gorm"
)

type TransactionsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransactionsRepositoryImpl(Db *gorm.DB) TransactionsRepository {
	return &TransactionsRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (r *TransactionsRepositoryImpl) Delete(id int64) (model.Transactions, error) {
	var bor model.Transactions
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements UsersRepository
func (r *TransactionsRepositoryImpl) FindAll() ([]model.Transactions, error) {
	var bor []model.Transactions
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements UsersRepository
func (r *TransactionsRepositoryImpl) FindById(id int64) (model.Transactions, error) {
	var bor model.Transactions
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("transactions is not found")
	}
	return bor, nil
}

// Save implements UsersRepository
func (r *TransactionsRepositoryImpl) Save(newTransactions model.Transactions) (model.Transactions, error) {
	currentTime := time.Now()
	newTransactions.Transactions_Date = currentTime
	result := r.Db.Create(&newTransactions)
	helper.ErrorPanic(result.Error)
	return newTransactions, nil
}

// Update implements UsersRepository
func (r *TransactionsRepositoryImpl) Update(updatedTransactions model.Transactions) (model.Transactions, error) {
	result := r.Db.Model(&model.Transactions{}).Where("id = ?", updatedTransactions.Id).Updates(updatedTransactions)
	helper.ErrorPanic(result.Error)
	return updatedTransactions, nil
}
