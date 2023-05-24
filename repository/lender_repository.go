package repository

import "github.com/sferawann/test2/model"

type LenderRepository interface {
	Save(newLender model.Lender)
	Update(updatedLender model.Lender)
	Delete(id int64)
	FindById(id int64) (model.Lender, error)
	FindAll() []model.Lender
	FindByUsername(username string) (model.Lender, error)
}