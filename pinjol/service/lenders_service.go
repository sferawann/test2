package service

import "github.com/sferawann/test2/pinjol/model"

type LendersService interface {
	Save(newLenders model.Lenders) (model.Lenders, error)
	Update(updateLenders model.Lenders) (model.Lenders, error)
	Delete(id int64) (model.Lenders, error)
	FindById(id int64) (model.Lenders, error)
	FindAll() ([]model.Lenders, error)
	FindByName(name string) (model.Lenders, error)
}
