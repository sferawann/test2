package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sferawann/test2/config"
	"github.com/sferawann/test2/data/request"
	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/repository"
	"github.com/sferawann/test2/utils"
)

type AuthServiceImpl struct {
	BorrowerRepository repository.BorrowerRepository
	Validate           *validator.Validate
}

func NewAuthServiceImpl(borrowerRepository repository.BorrowerRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		BorrowerRepository: borrowerRepository,
		Validate:           validate,
	}
}

// Login implements AuthenticationService
func (s *AuthServiceImpl) Login(borrower request.LoginRequest) (string, error) {
	// Find username in database
	new_borrower, err := s.BorrowerRepository.FindByUsername(borrower.Username)
	if err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_borrower.Password, borrower.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_borrower.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (s *AuthServiceImpl) Register(borrower request.CreateBorrowerRequest) {

	hashedPassword, err := utils.HashPassword(borrower.Password)
	helper.ErrorPanic(err)
	var bor model.Borrower
	create_at := bor.Created_At

	newBorrower := model.Borrower{
		Username:     borrower.Username,
		Password:     hashedPassword,
		Name:         borrower.Name,
		Alamat:       borrower.Alamat,
		Phone_Number: borrower.Phone_Number,
		Created_At:   create_at,
	}
	s.BorrowerRepository.Save(newBorrower)
}
