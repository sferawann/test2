package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/model"
	"gorm.io/gorm"
)

type LoanProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewLoanProductRepositoryImpl(Db *gorm.DB) LoanProductRepository {
	return &LoanProductRepositoryImpl{Db: Db}
}

// Delete implements LoanProductRepository
func (r *LoanProductRepositoryImpl) Delete(id int64) (model.LoanProduct, error) {
	var lp model.LoanProduct
	result := r.Db.Where("id = ?", id).Delete(&lp)
	helper.ErrorPanic(result.Error)
	return lp, nil
}

// FindAll implements LoanProductRepository
func (r *LoanProductRepositoryImpl) FindAll() ([]model.LoanProduct, error) {
	var lp []model.LoanProduct
	results := r.Db.Find(&lp)
	helper.ErrorPanic(results.Error)
	return lp, nil
}

// FindById implements LoanProductRepository
func (r *LoanProductRepositoryImpl) FindById(id int64) (model.LoanProduct, error) {
	var lp model.LoanProduct
	result := r.Db.Find(&lp, id)
	if result != nil {
		return lp, nil
	} else {
		return lp, errors.New("lender is not found")
	}
}

// FindByName implements LoanProductRepository
func (r *LoanProductRepositoryImpl) FindByName(name string) (model.LoanProduct, error) {
	var lp model.LoanProduct
	result := r.Db.Where("name LIKE ?", "%"+name+"%").First(&lp)

	if result.Error != nil {
		return lp, errors.New("invalid username or Password")
	}
	return lp, nil
}

// Save implements LoanProductRepository
func (r *LoanProductRepositoryImpl) Save(newLoanProduct model.LoanProduct) (model.LoanProduct, error) {
	currentTime := time.Now()
	newLoanProduct.Created_At = currentTime
	result := r.Db.Create(&newLoanProduct)
	helper.ErrorPanic(result.Error)
	return newLoanProduct, nil
}

// Update implements LoanProductRepository
func (r *LoanProductRepositoryImpl) Update(updatedLoanProduct model.LoanProduct) (model.LoanProduct, error) {
	result := r.Db.Model(&model.LoanProduct{}).Where("id = ?", updatedLoanProduct.Id).Updates(updatedLoanProduct)
	helper.ErrorPanic(result.Error)
	return updatedLoanProduct, nil
}
