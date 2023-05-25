package service

import "github.com/sferawann/test2/model"

type LenderService interface {
	Save(newLender model.Lender)
	Update(updatedLender model.Lender)
	Delete(id int64) (model.Lender, error)
	FindById(id int64) (model.Lender, error)
	FindAll() []model.Lender
}
