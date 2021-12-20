package classifications

import (
	"gym-membership/business/classification"

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

func (mysqlRepo *mysqlClassificationRepo) GetClassificationID(classification string) (uint, error) {
	rec := Classification{}
	err := mysqlRepo.Conn.First(&rec, "name = ?", classification).Error
	if err != nil {
		return 0, err
	}
	return rec.ID, nil
}

func (mysqlRepo *mysqlClassificationRepo) Insert(classificationData *classification.Domain) (classification.Domain, error) {
	rec := fromDomain(*classificationData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return classification.Domain{}, err
	}
	return rec.toDomain(), nil
}
