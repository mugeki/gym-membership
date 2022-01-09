package users

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/helper/encrypt"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepository 	Repository
	jwtAuth			*middleware.ConfigJWT
}

func NewUserUsecase(userRepo Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &userUsecase{
		userRepository: userRepo,
		jwtAuth: jwtauth,
	}
}

func (uc *userUsecase) Register(userData *Domain) (error)  {
	hashedPassword, _ := encrypt.Hash(userData.Password)
	userData.Password = hashedPassword
	userData.UUID = uuid.New()
	_, err := uc.userRepository.Register(userData)
	if err != nil {
		return business.ErrDuplicateData
	}
	return nil
}

func (uc *userUsecase) Login(username, password string) (Domain, error) {
	userDomain, err := uc.userRepository.GetByUsername(username)
	if err != nil {
		return Domain{}, business.ErrInvalidLoginInfo
	}

	if !encrypt.ValidateHash(password, userDomain.Password){
		return Domain{}, business.ErrInvalidLoginInfo
	}

	userDomain.Token = uc.jwtAuth.GenerateToken(int(userDomain.ID))
	return userDomain, nil
}