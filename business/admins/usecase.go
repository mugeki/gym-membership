package admins

import (
	"errors"
	"gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/helper/encrypt"

	"gorm.io/gorm"
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
	if adminData.Password != "" {
		hashedPassword, _ := encrypt.Hash(adminData.Password)
		adminData.Password = hashedPassword
	} else {
		data, err := uc.adminRepository.GetByID(id)
		if err != nil {
			return Domain{}, business.ErrInternalServer
		}
		adminData.Password = data.Password
	}
	
	adminDomain, err := uc.adminRepository.Update(id, adminData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return adminDomain, nil
}

func (uc *adminUsecase) GetAll(title string, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}
	res, totalData, err := uc.adminRepository.GetAll(title, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

func (uc *adminUsecase) DeleteByID(id uint) (error) {
	err := uc.adminRepository.DeleteByID(id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return business.ErrArticleNotFound
		} else {
			return business.ErrInternalServer
		}
	}
	return nil
}