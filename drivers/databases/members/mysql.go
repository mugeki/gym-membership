package members

import (
	"gym-membership/business/members"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlMemberRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) members.Repository {
	return &mysqlMemberRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlMemberRepo) Insert(data *members.Domain) error {
	rec := Members{}
	copier.Copy(&rec, &data)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return err
	}
	return nil
}

func (mysqlRepo *mysqlMemberRepo) GetByUserID(userID uint) (members.Domain, error) {
	rec := Members{}
	err := mysqlRepo.Conn.Order("expire_date desc").First(&rec, "user_id = ?", userID).Error
	if err != nil {
		return members.Domain{}, err
	}
	domain := members.Domain{}
	copier.Copy(&domain, &rec)
	return domain, nil
}
