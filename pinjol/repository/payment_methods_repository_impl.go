package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/pinjol/model"
	"gorm.io/gorm"
)

type PaymentMethodsRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentMethodsRepositoryImpl(Db *gorm.DB) PaymentMethodsRepository {
	return &PaymentMethodsRepositoryImpl{Db: Db}
}

// Delete implements PaymentMethodsRepository
func (r *PaymentMethodsRepositoryImpl) Delete(id int64) (model.PaymentMethods, error) {
	var bor model.PaymentMethods
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements PaymentMethodsRepository
func (r *PaymentMethodsRepositoryImpl) FindAll() ([]model.PaymentMethods, error) {
	var bor []model.PaymentMethods
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements PaymentMethodsRepository
func (r *PaymentMethodsRepositoryImpl) FindById(id int64) (model.PaymentMethods, error) {
	var bor model.PaymentMethods
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("PaymentMethods is not found")
	}
	return bor, nil
}

// Save implements PaymentMethodsRepository
func (r *PaymentMethodsRepositoryImpl) Save(newPaymentMethods model.PaymentMethods) (model.PaymentMethods, error) {
	currentTime := time.Now()
	newPaymentMethods.Created_At = currentTime
	result := r.Db.Create(&newPaymentMethods)
	helper.ErrorPanic(result.Error)
	return newPaymentMethods, nil
}

// Update implements PaymentMethodsRepository
func (r *PaymentMethodsRepositoryImpl) Update(updatedPaymentMethods model.PaymentMethods) (model.PaymentMethods, error) {
	result := r.Db.Model(&model.PaymentMethods{}).Where("id = ?", updatedPaymentMethods.Id).Updates(updatedPaymentMethods)
	helper.ErrorPanic(result.Error)
	return updatedPaymentMethods, nil
}
