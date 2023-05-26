package service

import (
	"github.com/go-playground/validator"
	"github.com/sferawann/test2/pinjol/model"
	"github.com/sferawann/test2/pinjol/repository"
)

type AcceptStatusServiceImpl struct {
	AcceptStatusRepository repository.AcceptStatusRepository
	Validate               *validator.Validate
}

// Delete implements BorrowerService
func (s *AcceptStatusServiceImpl) Delete(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *AcceptStatusServiceImpl) FindAll() ([]model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindAll()
}

// FindById implements BorrowerService
func (s *AcceptStatusServiceImpl) FindById(id int64) (model.AcceptStatus, error) {
	return s.AcceptStatusRepository.FindById(id)
}

// Save implements BorrowerService
func (s *AcceptStatusServiceImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {

	newAS := model.AcceptStatus{
		Transactions: newAcceptStatus.Transactions,
		Status:       newAcceptStatus.Status,
		Created_At:   newAcceptStatus.Created_At,
	}
	return s.AcceptStatusRepository.Save(newAS)

}

// Update implements BorrowerService
func (s *AcceptStatusServiceImpl) Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {

	var ast model.AcceptStatus
	create_at := ast.Created_At

	newAS := model.AcceptStatus{
		Id:           updateAcceptStatus.Id,
		Transactions: updateAcceptStatus.Transactions,
		Status:       updateAcceptStatus.Status,
		Created_At:   create_at,
	}

	return s.AcceptStatusRepository.Update(newAS)
}

func NewAcceptStatusServiceImpl(acceptStatusRepository repository.AcceptStatusRepository, validate *validator.Validate) AcceptStatusService {
	return &AcceptStatusServiceImpl{
		AcceptStatusRepository: acceptStatusRepository,
		Validate:               validate,
	}
}
