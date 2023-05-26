package service

import (
	"github.com/go-playground/validator"
	"github.com/sferawann/test2/pinjol/model"
	"github.com/sferawann/test2/pinjol/repository"
)

type StatusLoansServiceImpl struct {
	StatusLoansRepository repository.StatusLoansRepository
	Validate              *validator.Validate
}

// Delete implements BorrowerService
func (s *StatusLoansServiceImpl) Delete(id int64) (model.StatusLoans, error) {
	return s.StatusLoansRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *StatusLoansServiceImpl) FindAll() ([]model.StatusLoans, error) {
	return s.StatusLoansRepository.FindAll()
}

// FindById implements BorrowerService
func (s *StatusLoansServiceImpl) FindById(id int64) (model.StatusLoans, error) {
	return s.StatusLoansRepository.FindById(id)
}

// Save implements BorrowerService
func (s *StatusLoansServiceImpl) Save(newStatusLoans model.StatusLoans) (model.StatusLoans, error) {

	newSL := model.StatusLoans{
		Accept_Status: newStatusLoans.Accept_Status,
		Created_At:    newStatusLoans.Created_At,
	}
	return s.StatusLoansRepository.Save(newSL)

}

// Update implements BorrowerService
func (s *StatusLoansServiceImpl) Update(updateStatusLoans model.StatusLoans) (model.StatusLoans, error) {

	var sl model.StatusLoans
	create_at := sl.Created_At

	newSL := model.StatusLoans{
		Id:            updateStatusLoans.Id,
		Accept_Status: updateStatusLoans.Accept_Status,
		Created_At:    create_at,
	}

	return s.StatusLoansRepository.Update(newSL)
}

func NewStatusLoansServiceImpl(statusLoansRepository repository.StatusLoansRepository, validate *validator.Validate) StatusLoansService {
	return &StatusLoansServiceImpl{
		StatusLoansRepository: statusLoansRepository,
		Validate:              validate,
	}
}
