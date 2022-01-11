package users

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/business/members"
	"gym-membership/helper/encrypt"
	"time"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepository 	Repository
	memberRepository members.Repository
	jwtAuth			*middleware.ConfigJWT
}

func NewUserUsecase(userRepo Repository, memberRepo members.Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &userUsecase{
		userRepository: userRepo,
		memberRepository: memberRepo,
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

	isMember := false
	memberDomain, err := uc.memberRepository.GetByUserID(userDomain.ID)
	if err == nil && memberDomain.ExpireDate.Unix() >= time.Now().Unix() {
		isMember = true
	}
	userDomain.Token = uc.jwtAuth.GenerateToken(int(userDomain.ID), isMember, false, false)
	return userDomain, nil
}

func (uc *userUsecase) Update(id uint, userData *Domain) (Domain, error) {
	userDomain, err := uc.userRepository.Update(id, userData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return userDomain, nil
}