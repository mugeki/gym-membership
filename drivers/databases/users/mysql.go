package users

import (
	"gym-membership/business/users"

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
	recUser := fromDomain(*userData)
	err := mysqlRepo.Conn.Create(&recUser).Error
	if err != nil {
		return users.Domain{}, err
	}
	return recUser.toDomain(), nil
}

func (mysqlRepo *mysqlUsersRepo) GetByUsername(username string) (users.Domain, error){
	rec := Users{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}