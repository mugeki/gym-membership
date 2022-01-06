package trainers

import (
	"gym-membership/business/trainers"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlTrainersRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) trainers.Repository {
	return &mysqlTrainersRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlTrainersRepo) GetAll() ([]trainers.Domain, error) {
	domain := []trainers.Domain{}
	rec := []Trainers{}

	err := mysqlRepo.Conn.Find(&rec).Error
	if err != nil {
		return nil, err
	}

	copier.Copy(&domain, &rec)

	return domain, nil
}
