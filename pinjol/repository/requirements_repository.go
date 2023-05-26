package repository

import "github.com/sferawann/test2/pinjol/model"

type RequirementsRepository interface {
	Save(newRequirements model.Requirements) (model.Requirements, error)
	Update(updateRequirements model.Requirements) (model.Requirements, error)
	Delete(id int64) (model.Requirements, error)
	FindById(id int64) (model.Requirements, error)
	FindAll() ([]model.Requirements, error)
	FindByName(name string) (model.Requirements, error)
}
