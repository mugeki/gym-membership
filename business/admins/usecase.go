package admins

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/helper/encrypt"

	"github.com/google/uuid"
)

type adminUsecase struct {
	adminRepository Repository
	jwtAuth         *middleware.ConfigJWT
}

func NewAdminUsecase(adminRepo Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &adminUsecase{
		adminRepository: adminRepo,
		jwtAuth:         jwtauth,
	}
}

func (uc *adminUsecase) Register(userData *Domain) (string, error) {
	hashedPassword, _ := encrypt.Hash(userData.Password)
	userData.Password = hashedPassword
	userData.UUID = uuid.New()
	_, err := uc.adminRepository.Register(userData)
	if err != nil {
		return "", business.ErrDuplicateData
	}
	return "", nil
}

func (uc *adminUsecase) Login(username, password string) (string, error) {
	userDomain, err := uc.adminRepository.GetByUsername(username)
	if err != nil {
		return "", business.ErrInvalidLoginInfo
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", business.ErrInvalidLoginInfo
	}

	token := uc.jwtAuth.GenerateToken(int(userDomain.ID))
	return token, nil
}
