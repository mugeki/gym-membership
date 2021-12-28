package membership_products

import (
	"gym-membership/business/membership_products"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
	"fmt"
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
	println("repo membership_products", userData.Name)
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

func (mysqlRepo *mysqlMembershipProductsRepo) GetByUserID(idMembershipProducts uint) (string, error) {
	fmt.Println("mysql ",idMembershipProducts)
	// println("repo membership_products")
	rec := MembershipProducts{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", idMembershipProducts).Error
	if err != nil {
		return "data not found", err
	}
	return "succes", nil
}
