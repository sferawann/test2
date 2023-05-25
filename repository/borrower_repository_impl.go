package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/model"
	"gorm.io/gorm"
)

type BorrowerRepositoryImpl struct {
	Db *gorm.DB
}

func NewBorrowerRepositoryImpl(Db *gorm.DB) BorrowerRepository {
	return &BorrowerRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (r *BorrowerRepositoryImpl) Delete(id int64) (model.Borrower, error) {
	var bor model.Borrower
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements UsersRepository
func (r *BorrowerRepositoryImpl) FindAll() ([]model.Borrower, error) {
	var bor []model.Borrower
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements UsersRepository
func (r *BorrowerRepositoryImpl) FindById(id int64) (model.Borrower, error) {
	var bor model.Borrower
	result := r.Db.Find(&bor, id)
	if result != nil {
		return bor, nil
	} else {
		return bor, errors.New("Borrower is not found")
	}
}

// Save implements UsersRepository
func (r *BorrowerRepositoryImpl) Save(newBorrower model.Borrower) (model.Borrower, error) {
	currentTime := time.Now()
	newBorrower.Created_At = currentTime
	result := r.Db.Create(&newBorrower)
	helper.ErrorPanic(result.Error)
	return newBorrower, nil
}

// Update implements UsersRepository
func (r *BorrowerRepositoryImpl) Update(updatedBorrowers model.Borrower) (model.Borrower, error) {
	result := r.Db.Model(&model.Borrower{}).Where("id = ?", updatedBorrowers.Id).Updates(updatedBorrowers)
	helper.ErrorPanic(result.Error)
	return updatedBorrowers, nil
}

// FindByUsername implements UsersRepository
func (r *BorrowerRepositoryImpl) FindByUsername(username string) (model.Borrower, error) {
	var bor model.Borrower
	result := r.Db.First(&bor, "username = ?", username)

	if result.Error != nil {
		return bor, errors.New("invalid username or Password")
	}
	return bor, nil
}
