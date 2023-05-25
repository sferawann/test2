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

// Delete implements LenderRepository
func (r *LenderRepositoryImpl) Delete(id int64) (model.Lender, error) {
	var len model.Lender
	result := r.Db.Where("id = ?", id).Delete(&len)
	helper.ErrorPanic(result.Error)
	return len, nil
}

// FindAll implements LenderRepository
func (r *LenderRepositoryImpl) FindAll() ([]model.Lender, error) {
	var len []model.Lender
	results := r.Db.Find(&len)
	helper.ErrorPanic(results.Error)
	return len, nil
}

// FindById implements LenderRepository
func (r *LenderRepositoryImpl) FindById(id int64) (model.Lender, error) {
	var len model.Lender
	result := r.Db.Find(&len, id)
	if result != nil {
		return len, nil
	} else {
		return len, errors.New("lender is not found")
	}
}

// FindByName implements LenderRepository
func (r *LenderRepositoryImpl) FindByName(name string) (model.Lender, error) {
	var len model.Lender
	result := r.Db.First(&len, "name = ?", name)

	if result.Error != nil {
		return len, errors.New("invalid username or Password")
	}
	return len, nil
}

// Save implements LenderRepository
func (r *LenderRepositoryImpl) Save(newLender model.Lender) (model.Lender, error) {
	currentTime := time.Now()
	newLender.Created_At = currentTime
	result := r.Db.Create(&newLender)
	helper.ErrorPanic(result.Error)
	return newLender, nil
}

// Update implements LenderRepository
func (r *LenderRepositoryImpl) Update(updatedLender model.Lender) (model.Lender, error) {
	result := r.Db.Model(&model.Lender{}).Where("id = ?", updatedLender.Id).Updates(updatedLender)
	helper.ErrorPanic(result.Error)
	return updatedLender, nil
}
