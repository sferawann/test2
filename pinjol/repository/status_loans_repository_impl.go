package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/pinjol/model"
	"gorm.io/gorm"
)

type StatusLoansRepositoryImpl struct {
	Db *gorm.DB
}

func NewStatusLoansRepositoryImpl(Db *gorm.DB) StatusLoansRepository {
	return &StatusLoansRepositoryImpl{Db: Db}
}

// Delete implements StatusLoansRepository
func (r *StatusLoansRepositoryImpl) Delete(id int64) (model.StatusLoans, error) {
	var bor model.StatusLoans
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements StatusLoansRepository
func (r *StatusLoansRepositoryImpl) FindAll() ([]model.StatusLoans, error) {
	var bor []model.StatusLoans
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements StatusLoansRepository
func (r *StatusLoansRepositoryImpl) FindById(id int64) (model.StatusLoans, error) {
	var bor model.StatusLoans
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("StatusLoans is not found")
	}
	return bor, nil
}

// Save implements StatusLoansRepository
func (r *StatusLoansRepositoryImpl) Save(newStatusLoans model.StatusLoans) (model.StatusLoans, error) {
	currentTime := time.Now()
	newStatusLoans.Created_At = currentTime
	result := r.Db.Create(&newStatusLoans)
	helper.ErrorPanic(result.Error)
	return newStatusLoans, nil
}

// Update implements StatusLoansRepository
func (r *StatusLoansRepositoryImpl) Update(updatedStatusLoans model.StatusLoans) (model.StatusLoans, error) {
	result := r.Db.Model(&model.StatusLoans{}).Where("id = ?", updatedStatusLoans.Id).Updates(updatedStatusLoans)
	helper.ErrorPanic(result.Error)
	return updatedStatusLoans, nil
}
