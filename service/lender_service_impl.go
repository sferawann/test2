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

// Delete implements BorrowerService
func (s *LenderServiceImpl) Delete(id int64) (model.Lender, error) {
	return s.LenderRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *LenderServiceImpl) FindAll() []model.Lender {
	return s.LenderRepository.FindAll()
}

// FindById implements BorrowerService
func (s *LenderServiceImpl) FindById(id int64) (model.Lender, error) {
	return s.LenderRepository.FindById(id)
}

// Save implements BorrowerService
func (s *LenderServiceImpl) Save(newLender model.Lender) {
	newBor := model.Lender{
		Name:       newLender.Name,
		Created_At: newLender.Created_At,
	}
	s.LenderRepository.Save(newBor)
}

// Update implements BorrowerService
func (s *LenderServiceImpl) Update(updatedLender model.Lender) {
	var len model.Lender
	create_at := len.Created_At

	newLen := model.Lender{
		Name:       updatedLender.Name,
		Created_At: create_at,
	}
	s.LenderRepository.Update(newLen)
}

func NewLenderServiceImpl(lenderRepository repository.LenderRepository, validate *validator.Validate) LenderService {
	return &LenderServiceImpl{
		LenderRepository: lenderRepository,
		Validate:         validate,
	}
}
