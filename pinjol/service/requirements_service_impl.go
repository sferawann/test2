package service

import (
	"github.com/go-playground/validator"
	"github.com/sferawann/test2/pinjol/model"
	"github.com/sferawann/test2/pinjol/repository"
)

type RequirementsServiceImpl struct {
	RequirementsRepository repository.RequirementsRepository
	Validate               *validator.Validate
}

// Delete implements BorrowerService
func (s *RequirementsServiceImpl) Delete(id int64) (model.Requirements, error) {
	return s.RequirementsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *RequirementsServiceImpl) FindAll() ([]model.Requirements, error) {
	return s.RequirementsRepository.FindAll()
}

// FindById implements BorrowerService
func (s *RequirementsServiceImpl) FindById(id int64) (model.Requirements, error) {
	return s.RequirementsRepository.FindById(id)
}

// FindByUsername implements BorrowerService
func (s *RequirementsServiceImpl) FindByName(name string) (model.Requirements, error) {
	return s.RequirementsRepository.FindByName(name)
}

// Save implements BorrowerService
func (s *RequirementsServiceImpl) Save(newRequirements model.Requirements) (model.Requirements, error) {

	newReq := model.Requirements{
		Nik:          newRequirements.Nik,
		Name:         newRequirements.Name,
		Alamat:       newRequirements.Alamat,
		Phone_Number: newRequirements.Phone_Number,
		Created_At:   newRequirements.Created_At,
	}
	return s.RequirementsRepository.Save(newReq)

}

// Update implements BorrowerService
func (s *RequirementsServiceImpl) Update(updateRequirements model.Requirements) (model.Requirements, error) {

	var req model.Requirements
	create_at := req.Created_At

	newReq := model.Requirements{
		Id:           updateRequirements.Id,
		Nik:          updateRequirements.Nik,
		Name:         updateRequirements.Name,
		Alamat:       updateRequirements.Alamat,
		Phone_Number: updateRequirements.Phone_Number,
		Created_At:   create_at,
	}

	return s.RequirementsRepository.Update(newReq)
}

func NewRequirementsServiceImpl(requirementsRepository repository.RequirementsRepository, validate *validator.Validate) RequirementsService {
	return &RequirementsServiceImpl{
		RequirementsRepository: requirementsRepository,
		Validate:               validate,
	}
}
