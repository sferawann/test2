package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/repository"
	"github.com/sferawann/test2/utils"
)

type BorrowerServiceImpl struct {
	BorrowerRepository repository.BorrowerRepository
	Validate           *validator.Validate
}

// Delete implements BorrowerService
func (s *BorrowerServiceImpl) Delete(deletedBorrower model.Borrower) error {
	return s.BorrowerRepository.Delete(deletedBorrower)
}

// FindAll implements BorrowerService
func (s *BorrowerServiceImpl) FindAll() []model.Borrower {
	return s.BorrowerRepository.FindAll()
}

// FindById implements BorrowerService
func (s *BorrowerServiceImpl) FindById(id int64) (model.Borrower, error) {
	return s.BorrowerRepository.FindById(id)
}

// FindByUsername implements BorrowerService
func (s *BorrowerServiceImpl) FindByUsername(username string) (model.Borrower, error) {
	return s.BorrowerRepository.FindByUsername(username)
}

// Save implements BorrowerService
func (s *BorrowerServiceImpl) Save(newBorrower model.Borrower) {
	hashedPassword, err := utils.HashPassword(newBorrower.Password)
	helper.ErrorPanic(err)
	var bor model.Borrower
	create_at := bor.Created_At

	newBor := model.Borrower{
		Username:     newBorrower.Username,
		Password:     hashedPassword,
		Name:         newBorrower.Name,
		Alamat:       newBorrower.Alamat,
		Phone_Number: newBorrower.Phone_Number,
		Created_At:   create_at,
	}
	s.BorrowerRepository.Save(newBor)
}

// Update implements BorrowerService
func (s *BorrowerServiceImpl) Update(updatedBorrower model.Borrower) {
	hashedPassword, err := utils.HashPassword(updatedBorrower.Password)
	helper.ErrorPanic(err)
	var bor model.Borrower
	create_at := bor.Created_At

	newBor := model.Borrower{
		Username:     updatedBorrower.Username,
		Password:     hashedPassword,
		Name:         updatedBorrower.Name,
		Alamat:       updatedBorrower.Alamat,
		Phone_Number: updatedBorrower.Phone_Number,
		Created_At:   create_at,
	}
	s.BorrowerRepository.Update(newBor)
}

func NewBorrowerServiceImpl(borrowerRepository repository.BorrowerRepository, validate *validator.Validate) BorrowerService {
	return &BorrowerServiceImpl{
		BorrowerRepository: borrowerRepository,
		Validate:           validate,
	}
}
