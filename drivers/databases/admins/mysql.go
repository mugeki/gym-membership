package admins

import (
	"gym-membership/business/admins"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlAdminsRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) admins.Repository {
	return &mysqlAdminsRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlAdminsRepo) Register(adminData *admins.Domain) (admins.Domain, error) {
	domain := admins.Domain{}
	rec := Admins{}
	copier.Copy(&rec, &adminData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return admins.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlAdminsRepo) GetByUsername(username string) (admins.Domain, error) {
	domain := admins.Domain{}
	rec := Admins{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).Error
	if err != nil {
		return admins.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlAdminsRepo) Update(id uint, adminData *admins.Domain) (admins.Domain, error) {
	domain := admins.Domain{}
	rec := Admins{}
	recData := Admins{}
	copier.Copy(&recData, &adminData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).
		Update("is_super_admin",adminData.IsSuperAdmin).Error
	if err != nil {
		return admins.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}
