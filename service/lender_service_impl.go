package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/repository"
)

type LenderServiceImpl struct {
	LenderRepository repository.LenderRepository
	Validate         *validator.Validate
}

// Delete implements LenderService
func (s *LenderServiceImpl) Delete(id int64) (model.Lender, error) {
	return s.LenderRepository.Delete(id)
}

// FindAll implements LenderService
func (s *LenderServiceImpl) FindAll() ([]model.Lender, error) {
	return s.LenderRepository.FindAll()
}

// FindById implements LenderService
func (s *LenderServiceImpl) FindById(id int64) (model.Lender, error) {
	return s.LenderRepository.FindById(id)
}

// FindByName implements LenderService
func (s *LenderServiceImpl) FindByName(name string) (model.Lender, error) {
	return s.LenderRepository.FindByName(name)
}

// Save implements LenderService
func (s *LenderServiceImpl) Save(newLender model.Lender) (model.Lender, error) {
	newLen := model.Lender{
		Name:       newLender.Name,
		Created_At: newLender.Created_At,
	}
	return s.LenderRepository.Save(newLen)

}

// Update implements LenderService
func (s *LenderServiceImpl) Update(updatedLender model.Lender) (model.Lender, error) {
	var len model.Lender
	create_at := len.Created_At

	newLen := model.Lender{
		Name:       updatedLender.Name,
		Created_At: create_at,
	}

	return s.LenderRepository.Update(newLen)
}

func NewLenderServiceImpl(lenderRepository repository.LenderRepository, validate *validator.Validate) LenderService {
	return &LenderServiceImpl{
		LenderRepository: lenderRepository,
		Validate:         validate,
	}
}
