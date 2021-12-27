package classifications

import (
	"gym-membership/business/classification"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type mysqlClassificationRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) classification.Repository {
	return &mysqlClassificationRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlClassificationRepo) GetAll() ([]classification.Domain, error) {
	domain := []classification.Domain{}
	rec := []Classification{}

	err := mysqlRepo.Conn.Find(&rec).Error
	if err != nil {
		return nil, err
	}

	copier.Copy(&domain, &rec)

	return domain, nil
}

func (mysqlRepo *mysqlClassificationRepo) Insert(classificationData *classification.Domain) (classification.Domain, error) {
	rec := Classification{}
	domain := classification.Domain{}
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return classification.Domain{}, err
	}

	copier.Copy(&domain, &rec)
	return domain, nil
}
