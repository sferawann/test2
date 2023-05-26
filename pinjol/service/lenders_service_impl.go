package service

import (
	"github.com/go-playground/validator"
	"github.com/sferawann/test2/pinjol/model"
	"github.com/sferawann/test2/pinjol/repository"
)

type LendersServiceImpl struct {
	LendersRepository repository.LendersRepository
	Validate          *validator.Validate
}

// Delete implements BorrowerService
func (s *LendersServiceImpl) Delete(id int64) (model.Lenders, error) {
	return s.LendersRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *LendersServiceImpl) FindAll() ([]model.Lenders, error) {
	return s.LendersRepository.FindAll()
}

// FindById implements BorrowerService
func (s *LendersServiceImpl) FindById(id int64) (model.Lenders, error) {
	return s.LendersRepository.FindById(id)
}

// FindByUsername implements BorrowerService
func (s *LendersServiceImpl) FindByName(name string) (model.Lenders, error) {
	return s.LendersRepository.FindByName(name)
}

// Save implements BorrowerService
func (s *LendersServiceImpl) Save(newLenders model.Lenders) (model.Lenders, error) {

	newLen := model.Lenders{
		Name:       newLenders.Name,
		Limits:     newLenders.Limits,
		Bunga:      newLenders.Bunga,
		Created_At: newLenders.Created_At,
	}
	return s.LendersRepository.Save(newLen)

}

// Update implements BorrowerService
func (s *LendersServiceImpl) Update(updateLenders model.Lenders) (model.Lenders, error) {

	var len model.Lenders
	create_at := len.Created_At

	newLen := model.Lenders{
		Id:         updateLenders.Id,
		Name:       updateLenders.Name,
		Limits:     updateLenders.Limits,
		Bunga:      updateLenders.Bunga,
		Created_At: create_at,
	}

	return s.LendersRepository.Update(newLen)
}

func NewLendersServiceImpl(lendersRepository repository.LendersRepository, validate *validator.Validate) LendersService {
	return &LendersServiceImpl{
		LendersRepository: lendersRepository,
		Validate:          validate,
	}
}
