package service

import "github.com/sferawann/test2/data/request"

type AuthService interface {
	Login(borrower request.LoginRequest) (string, error)
	Register(borrower request.CreateBorrowerRequest)
}
