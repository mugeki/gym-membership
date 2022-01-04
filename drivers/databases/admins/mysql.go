package admins

import (
	"gym-membership/business/admins"

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
	recAdmin := fromDomain(*adminData)
	err := mysqlRepo.Conn.Create(&recAdmin).Error
	if err != nil {
		return admins.Domain{}, err
	}
	return recAdmin.toDomain(), nil
}

func (mysqlRepo *mysqlAdminsRepo) GetByUsername(username string) (admins.Domain, error) {
	rec := Admins{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).Error
	if err != nil {
		return admins.Domain{}, err
	}
	return rec.toDomain(), nil
}
