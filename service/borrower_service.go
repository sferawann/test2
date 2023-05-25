package service

import "github.com/sferawann/test2/model"

type BorrowerService interface {
	Save(newBorrower model.Borrower)
	Update(updatedBorrower model.Borrower)
	Delete(id int64) (model.Borrower, error)
	FindById(id int64) (model.Borrower, error)
	FindAll() []model.Borrower
	FindByUsername(username string) (model.Borrower, error)
}
