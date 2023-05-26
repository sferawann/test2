package service

import "github.com/sferawann/test2/pinjol/model"

type StatusLoansService interface {
	Save(newStatusLoans model.StatusLoans) (model.StatusLoans, error)
	Update(updateStatusLoans model.StatusLoans) (model.StatusLoans, error)
	Delete(id int64) (model.StatusLoans, error)
	FindById(id int64) (model.StatusLoans, error)
	FindAll() ([]model.StatusLoans, error)
}
