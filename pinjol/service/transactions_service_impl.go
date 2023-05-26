package service

import (
	"github.com/go-playground/validator"
	"github.com/sferawann/test2/pinjol/model"
	"github.com/sferawann/test2/pinjol/repository"
)

type TransactionsServiceImpl struct {
	TransactionsRepository repository.TransactionsRepository
	Validate               *validator.Validate
}

// Delete implements BorrowerService
func (s *TransactionsServiceImpl) Delete(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *TransactionsServiceImpl) FindAll() ([]model.Transactions, error) {
	return s.TransactionsRepository.FindAll()
}

// FindById implements BorrowerService
func (s *TransactionsServiceImpl) FindById(id int64) (model.Transactions, error) {
	return s.TransactionsRepository.FindById(id)
}

// Save implements BorrowerService
func (s *TransactionsServiceImpl) Save(newTransactions model.Transactions) (model.Transactions, error) {

	newTra := model.Transactions{
		Requirements:      newTransactions.Requirements,
		Lenders:           newTransactions.Lenders,
		Users:             newTransactions.Users,
		Amount:            newTransactions.Amount,
		Transactions_Date: newTransactions.Transactions_Date,
		Due_Date:          newTransactions.Due_Date,
	}
	return s.TransactionsRepository.Save(newTra)

}

// Update implements BorrowerService
func (s *TransactionsServiceImpl) Update(updateTransactions model.Transactions) (model.Transactions, error) {

	var tra model.Transactions
	transaction_date := tra.Transactions_Date

	newTra := model.Transactions{
		Id:                updateTransactions.Id,
		Requirements:      updateTransactions.Requirements,
		Lenders:           updateTransactions.Lenders,
		Users:             updateTransactions.Users,
		Amount:            updateTransactions.Amount,
		Transactions_Date: transaction_date,
		Due_Date:          updateTransactions.Due_Date,
	}

	return s.TransactionsRepository.Update(newTra)
}

func NewTransactionsServiceImpl(transactionsRepository repository.TransactionsRepository, validate *validator.Validate) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionsRepository,
		Validate:               validate,
	}
}
