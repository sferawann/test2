package service

import "github.com/sferawann/test2/model"

type LenderService interface {
	Save(newLender model.Lender) (model.Lender, error)
	Update(updatedLender model.Lender) (model.Lender, error)
	Delete(id int64) (model.Lender, error)
	FindById(id int64) (model.Lender, error)
	FindAll() ([]model.Lender, error)
	FindByName(name string) (model.Lender, error)
}
