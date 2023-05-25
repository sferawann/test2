package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/repository"
)

type LoanProductServiceImpl struct {
	LoanProductRepository repository.LoanProductRepository
	Validate              *validator.Validate
}

// Delete implements LoanProductService
func (s *LoanProductServiceImpl) Delete(id int64) (model.LoanProduct, error) {
	return s.LoanProductRepository.Delete(id)
}

// FindAll implements LoanProductService
func (s *LoanProductServiceImpl) FindAll() ([]model.LoanProduct, error) {
	return s.LoanProductRepository.FindAll()
}

// FindById implements LoanProductService
func (s *LoanProductServiceImpl) FindById(id int64) (model.LoanProduct, error) {
	return s.LoanProductRepository.FindById(id)
}

// FindByName implements LoanProductService
func (s *LoanProductServiceImpl) FindByName(name string) (model.LoanProduct, error) {
	return s.LoanProductRepository.FindByName(name)
}

// Save implements LoanProductService
func (s *LoanProductServiceImpl) Save(newLoanProduct model.LoanProduct) (model.LoanProduct, error) {
	newLP := model.LoanProduct{
		Name:        newLoanProduct.Name,
		Description: newLoanProduct.Description,
		Persyaratan: newLoanProduct.Persyaratan,
		Created_At:  newLoanProduct.Created_At,
	}
	return s.LoanProductRepository.Save(newLP)
}

// Update implements LoanProductService
func (s *LoanProductServiceImpl) Update(updatedLoanProduct model.LoanProduct) (model.LoanProduct, error) {
	var lp model.LoanProduct
	create_at := lp.Created_At

	newLP := model.LoanProduct{
		Id:          updatedLoanProduct.Id,
		Name:        updatedLoanProduct.Name,
		Description: updatedLoanProduct.Description,
		Persyaratan: updatedLoanProduct.Persyaratan,
		Created_At:  create_at,
	}

	return s.LoanProductRepository.Update(newLP)
}

func NewLoanProductServiceImpl(loanProductRepository repository.LoanProductRepository, validate *validator.Validate) LoanProductService {
	return &LoanProductServiceImpl{
		LoanProductRepository: loanProductRepository,
		Validate:              validate,
	}
}
