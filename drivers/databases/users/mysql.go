package users

import (
	"gym-membership/business/users"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlUsersRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepo {
		Conn: conn,
	}
}

func (mysqlRepo *mysqlUsersRepo) Register(userData *users.Domain) (users.Domain, error) {
	domain := users.Domain{}
	rec := Users{}
	copier.Copy(&rec, &userData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlUsersRepo) GetByUsername(username string) (users.Domain, error){
	domain := users.Domain{}
	rec := Users{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).
		Or(&rec, "email = ?", username).Error
	if err != nil {
		return users.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}