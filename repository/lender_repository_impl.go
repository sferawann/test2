package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/model"
	"gorm.io/gorm"
)

type LenderRepositoryImpl struct {
	Db *gorm.DB
}

func NewLenderRepositoryImpl(Db *gorm.DB) LenderRepository {
	return &LenderRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (u *LenderRepositoryImpl) Delete(id int64) (model.Lender, error) {
	var len model.Lender
	result := u.Db.Where("id = ?", id).Delete(&len)
	helper.ErrorPanic(result.Error)
	return len, nil
}

// FindAll implements UsersRepository
func (u *LenderRepositoryImpl) FindAll() []model.Lender {
	var len []model.Lender
	results := u.Db.Find(&len)
	helper.ErrorPanic(results.Error)
	return len
}

// FindById implements UsersRepository
func (u *LenderRepositoryImpl) FindById(id int64) (model.Lender, error) {
	var len model.Lender
	result := u.Db.Find(&len, id)
	if result != nil {
		return len, nil
	} else {
		return len, errors.New("Lender is not found")
	}
}

// Save implements UsersRepository
func (u *LenderRepositoryImpl) Save(newLender model.Lender) {
	currentTime := time.Now()
	newLender.Created_At = currentTime
	result := u.Db.Create(&newLender)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (u *LenderRepositoryImpl) Update(updatedLender model.Lender) {
	var len model.Lender
	create_at := len.Created_At
	var updateLender = model.Lender{
		Id:         updatedLender.Id,
		Name:       updatedLender.Name,
		Created_At: create_at,
	}
	result := u.Db.Model(&updatedLender).Updates(updateLender)
	helper.ErrorPanic(result.Error)
}
