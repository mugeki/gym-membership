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

func (mysqlRepo *mysqlMembershipProductRepo) UpdateByID(idMembershipProducts uint, membership_products.Domain, error) {
	// println("cek id", id)
	domain := membership_products.Domain{}
	rec := membership_products{}
	// domainData := articles.Domain{}
	recData := membership_products{}
	copier.Copy(idMembershipProducts)
	err := mysqlRepo.Conn.First(&rec, "id = ?", idMembershipProducts).Updates(recData).Error
	if err != nil {
		return articles.Domain{}, err
	}
	copier.Copy(domain, rec)
	return domain, nil
}

func (mysqlRepo *mysqlMembershipProductRepo) DeleteByID(idMembershipProducts uint) error {
	rec := membership_products{}
	err := mysqlRepo.Conn.First(&rec, idMembershipProducts).Delete(&rec).Error
	if rec.ID == 0 {
		// println("not found", rec.Title)
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}