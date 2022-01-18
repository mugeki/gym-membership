package admins

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/helper/encrypt"
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

func (uc *adminUsecase) Register(adminData *Domain) (Domain, error) {
	hashedPassword, _ := encrypt.Hash(adminData.Password)
	adminData.Password = hashedPassword
	adminDomain, err := uc.adminRepository.Register(adminData)
	if err != nil {
		return Domain{}, business.ErrDuplicateData
	}
	return adminDomain, nil
}

func (uc *adminUsecase) Login(username, password string) (Domain, error) {
	adminDomain, err := uc.adminRepository.GetByUsername(username)
	if err != nil {
		return Domain{}, business.ErrInvalidLoginInfo
	}

	if !encrypt.ValidateHash(password, adminDomain.Password) {
		return Domain{}, business.ErrInvalidLoginInfo
	}

	adminDomain.Token = uc.jwtAuth.GenerateToken(int(adminDomain.ID), false, true, adminDomain.IsSuperAdmin)
	return adminDomain, nil
}

func (uc *adminUsecase) Update(id uint, adminData *Domain) (Domain, error) {
	adminDomain, err := uc.adminRepository.Update(id, adminData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return adminDomain, nil
}
