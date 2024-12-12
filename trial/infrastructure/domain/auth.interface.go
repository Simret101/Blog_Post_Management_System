package domain

import (
	"github.com/gin-gonic/gin"
	"practice/user_service/domain"
	
)

type AuthController interface {
	SignUp() gin.HandlerFunc
	LogIn() gin.HandlerFunc
	GoogleLogIn() gin.HandlerFunc
	GoogleCallBack() gin.HandlerFunc
	LogOut() gin.HandlerFunc
	Refresh() gin.HandlerFunc
}

type AuthUsecase interface {
	RegisterUser(domain.RegisterUser) (domain.User, error)
	LoginUser(string, string) (domain.User, string, string, error)
	GoogleLogin() (string, error)
	GoogleCallBack(string, string)(*domain.User, string, string, error)
	RefreshTokens(string) (string, string, error) 
}

type AuthRepository interface{
	SaveUser(*domain.User) error
	FindUserByEmail(string) (*domain.User, error)
}

type StateRepository interface{
	InsertState(State) error
	GetState(string) (*State, error)
	DeleteState(string) error
}