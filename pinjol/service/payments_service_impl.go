package service

import (
	"github.com/go-playground/validator"
	"github.com/sferawann/test2/pinjol/model"
	"github.com/sferawann/test2/pinjol/repository"
)

type PaymentsServiceImpl struct {
	PaymentsRepository repository.PaymentsRepository
	Validate           *validator.Validate
}

// Delete implements BorrowerService
func (s *PaymentsServiceImpl) Delete(id int64) (model.Payments, error) {
	return s.PaymentsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *PaymentsServiceImpl) FindAll() ([]model.Payments, error) {
	return s.PaymentsRepository.FindAll()
}

// FindById implements BorrowerService
func (s *PaymentsServiceImpl) FindById(id int64) (model.Payments, error) {
	return s.PaymentsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *PaymentsServiceImpl) Save(newPayments model.Payments) (model.Payments, error) {

	newPay := model.Payments{
		Transactions:   newPayments.Transactions,
		Payment_Method: newPayments.Payment_Method,
		Payment_Amount: newPayments.Payment_Amount,
		Payment_Date:   newPayments.Payment_Date,
	}
	return s.PaymentsRepository.Save(newPay)

}

// Update implements BorrowerService
func (s *PaymentsServiceImpl) Update(updatePayments model.Payments) (model.Payments, error) {

	var pay model.Payments
	payment_date := pay.Payment_Date

	newPay := model.Payments{
		Id:             updatePayments.Id,
		Transactions:   updatePayments.Transactions,
		Payment_Method: updatePayments.Payment_Method,
		Payment_Amount: updatePayments.Payment_Amount,
		Payment_Date:   payment_date,
	}

	return s.PaymentsRepository.Update(newPay)
}

func NewPaymentsServiceImpl(paymentsRepository repository.PaymentsRepository, validate *validator.Validate) PaymentsService {
	return &PaymentsServiceImpl{
		PaymentsRepository: paymentsRepository,
		Validate:           validate,
	}
}