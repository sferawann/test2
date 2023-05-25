package repository

import "github.com/sferawann/test2/model"

type LoanProductRepository interface {
	Save(newLoanProduct model.LoanProduct) (model.LoanProduct, error)
	Update(updatedLoanProduct model.LoanProduct) (model.LoanProduct, error)
	Delete(id int64) (model.LoanProduct, error)
	FindById(id int64) (model.LoanProduct, error)
	FindAll() ([]model.LoanProduct, error)
	FindByName(name string) (model.LoanProduct, error)
}
