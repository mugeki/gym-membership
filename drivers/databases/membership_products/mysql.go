package membership_products

import (
	"gym-membership/business/membership_products"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type mysqlMembershipProductsRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) membership_products.Repository {
	return &mysqlMembershipProductsRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlMembershipProductsRepo) Insert(userData *membership_products.Domain) (membership_products.Domain, error) {
	domain := membership_products.Domain{}
	recUser := MembershipProducts{}
	copier.Copy(&recUser, &userData)
	err := mysqlRepo.Conn.Create(&recUser).Error
	if err != nil {
		return membership_products.Domain{}, err
	}
	copier.Copy(domain, recUser)
	return domain, nil
}

func (mysqlRepo *mysqlMembershipProductsRepo) GetByID(idMembershipProducts uint) (membership_products.Domain, error) {
	rec := MembershipProducts{}
	domain := membership_products.Domain{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", idMembershipProducts).Error
	if err != nil {
		return membership_products.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}
