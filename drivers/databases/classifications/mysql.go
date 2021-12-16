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

func (mysqlRepo *mysqlClassificationRepo) GetNameByID(id uint) (classification.Domain, error) {
	recClassificationName := Classification{}
	err := mysqlRepo.Conn.First(&recClassificationName, "id = ?", id).Error
	if err != nil {
		return classification.Domain{}, err
	}
	return recClassificationName.toDomain(), nil
}
