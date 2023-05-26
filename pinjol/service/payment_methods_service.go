package service

import "github.com/sferawann/test2/pinjol/model"

type PaymentMethodsService interface {
	Save(newPaymentMethods model.PaymentMethods) (model.PaymentMethods, error)
	Update(updatePaymentMethods model.PaymentMethods) (model.PaymentMethods, error)
	Delete(id int64) (model.PaymentMethods, error)
	FindById(id int64) (model.PaymentMethods, error)
	FindAll() ([]model.PaymentMethods, error)
}
