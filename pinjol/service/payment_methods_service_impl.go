package service

import (
	"github.com/go-playground/validator"
	"github.com/sferawann/test2/pinjol/model"
	"github.com/sferawann/test2/pinjol/repository"
)

type PaymentMethodsServiceImpl struct {
	PaymentMethodsRepository repository.PaymentMethodsRepository
	Validate                 *validator.Validate
}

// Delete implements BorrowerService
func (s *PaymentMethodsServiceImpl) Delete(id int64) (model.PaymentMethods, error) {
	return s.PaymentMethodsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *PaymentMethodsServiceImpl) FindAll() ([]model.PaymentMethods, error) {
	return s.PaymentMethodsRepository.FindAll()
}

// FindById implements BorrowerService
func (s *PaymentMethodsServiceImpl) FindById(id int64) (model.PaymentMethods, error) {
	return s.PaymentMethodsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *PaymentMethodsServiceImpl) Save(newPaymentMethods model.PaymentMethods) (model.PaymentMethods, error) {

	newPM := model.PaymentMethods{
		Name:       newPaymentMethods.Name,
		Created_At: newPaymentMethods.Created_At,
	}
	return s.PaymentMethodsRepository.Save(newPM)

}

// Update implements BorrowerService
func (s *PaymentMethodsServiceImpl) Update(updatePaymentMethods model.PaymentMethods) (model.PaymentMethods, error) {

	var pm model.PaymentMethods
	create_at := pm.Created_At

	newPM := model.PaymentMethods{
		Id:         updatePaymentMethods.Id,
		Name:       updatePaymentMethods.Name,
		Created_At: create_at,
	}

	return s.PaymentMethodsRepository.Update(newPM)
}

func NewPaymentMethodsServiceImpl(paymentMethodsRepository repository.PaymentMethodsRepository, validate *validator.Validate) PaymentMethodsService {
	return &PaymentMethodsServiceImpl{
		PaymentMethodsRepository: paymentMethodsRepository,
		Validate:                 validate,
	}
}
