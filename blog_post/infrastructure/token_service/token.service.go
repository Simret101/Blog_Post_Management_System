package tokenservice

import (
	"errors"
	"time"

	users"practice/user_service/domain"
	infra"practice/infrastructure/domain"
	"github.com/dgrijalva/jwt-go"
)

type TokenService_imp struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

func NewTokenService(accessSecret, refreshSecret string) *TokenService_imp {
	return &TokenService_imp{
		AccessTokenSecret:  accessSecret,
		RefreshTokenSecret: refreshSecret,
	}
}

func (t *TokenService_imp) GenerateAccessToken(user users.User) (string, error) {
	claims := infra.UserClaims{
		ID:      user.ID,
		Name:    user.UserName,
		Avatar:  user.ProfilePicture,
		Email:   user.Email,
		IsAdmin: user.Is_Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.AccessTokenSecret))
}

func (t *TokenService_imp) GenerateRefreshToken(user users.User) (string, error) {
	claims := infra.UserClaims{
		ID:      user.ID,
		Name:    user.UserName,
		Avatar:  user.ProfilePicture,
		Email:   user.Email,
		IsAdmin: user.Is_Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.RefreshTokenSecret))
}

func (t *TokenService_imp) ValidateAccessToken(tokenStr string) (*users.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &infra.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	claims, ok := token.Claims.(*infra.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return &users.User{
		ID:             claims.ID,
		UserName:       claims.Name,
		ProfilePicture: claims.Avatar,
		Email:          claims.Email,
		Is_Admin:       claims.IsAdmin,
	}, nil
}

func (t *TokenService_imp) ValidateRefreshToken(tokenStr string) (*users.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &infra.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.RefreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*infra.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return &users.User{
		ID:             claims.ID,
		UserName:       claims.Name,
		ProfilePicture: claims.Avatar,
		Email:          claims.Email,
		Is_Admin:       claims.IsAdmin,
	}, nil
}
