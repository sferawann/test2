package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/data/request"
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
func (u *BorrowerRepositoryImpl) Delete(updatedBorrowers model.Borrower) error {
	var bor model.Borrower
	result := u.Db.Where("id = ?", updatedBorrowers).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return nil
}

// FindAll implements UsersRepository
func (u *BorrowerRepositoryImpl) FindAll() []model.Borrower {
	var bor []model.Borrower
	results := u.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor
}

// FindById implements UsersRepository
func (u *BorrowerRepositoryImpl) FindById(id int64) (model.Borrower, error) {
	var bor model.Borrower
	result := u.Db.Find(&bor, id)
	if result != nil {
		return bor, nil
	} else {
		return bor, errors.New("Borrower is not found")
	}
}

// Save implements UsersRepository
func (u *BorrowerRepositoryImpl) Save(newBorrower model.Borrower) (model.Borrower, error) {
	currentTime := time.Now()
	newBorrower.Created_At = currentTime
	result := u.Db.Create(&newBorrower)
	if result.Error != nil {
		return model.Borrower{}, result.Error
	}
	return newBorrower, nil
}

// Update implements UsersRepository
func (u *BorrowerRepositoryImpl) Update(updatedBorrowers model.Borrower) {
	var bor model.Borrower
	create_at := bor.Created_At
	var updateBorrower = request.UpdateBorrowerRequest{
		Id:           updatedBorrowers.Id,
		Username:     updatedBorrowers.Username,
		Password:     updatedBorrowers.Password,
		Name:         updatedBorrowers.Name,
		Alamat:       updatedBorrowers.Alamat,
		Phone_Number: updatedBorrowers.Phone_Number,
		Created_At:   create_at,
	}
	result := u.Db.Model(&updatedBorrowers).Updates(updateBorrower)
	helper.ErrorPanic(result.Error)
}

// FindByUsername implements UsersRepository
func (u *BorrowerRepositoryImpl) FindByUsername(username string) (model.Borrower, error) {
	var bor model.Borrower
	result := u.Db.First(&bor, "username = ?", username)

	if result.Error != nil {
		return bor, errors.New("invalid username or Password")
	}
	return bor, nil
}
