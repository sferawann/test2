package repository

import "github.com/sferawann/test2/model"

type BorrowerRepository interface {
	Save(newBorrower model.Borrower) (model.Borrower, error)
	Update(updatedBorrower model.Borrower) (model.Borrower, error)
	Delete(id int64) (model.Borrower, error)
	FindById(id int64) (model.Borrower, error)
	FindAll() ([]model.Borrower, error)
	FindByUsername(username string) (model.Borrower, error)
}
