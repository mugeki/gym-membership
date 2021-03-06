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

func (mysqlRepo *mysqlMembershipProductsRepo) Insert(newData *membership_products.Domain) (membership_products.Domain,error) {
	domain := membership_products.Domain{}
	rec := MembershipProducts{}
	copier.Copy(&rec, &newData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return membership_products.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlMembershipProductsRepo) GetAll(offset, limit int) ([]membership_products.Domain, int64, error) {
	var totalData int64
	domain := []membership_products.Domain{}
	rec := []MembershipProducts{}

	mysqlRepo.Conn.Find(&rec).Count(&totalData)
	err := mysqlRepo.Conn.Limit(limit).Offset(offset).
		Order("updated_at desc").Find(&rec).Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	return domain, totalData, nil
}


func (mysqlRepo *mysqlMembershipProductsRepo) GetByID(id uint) (membership_products.Domain, error) {
	rec := MembershipProducts{}
	domain := membership_products.Domain{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return membership_products.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlMembershipProductsRepo) UpdateByID(id uint, newData *membership_products.Domain)  (membership_products.Domain, error) {
	domain := membership_products.Domain{}
	rec := MembershipProducts{}
	recData := MembershipProducts{}
	copier.Copy(&recData, &newData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).Error
	if err != nil {
		return membership_products.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlMembershipProductsRepo) DeleteByID(id uint) error {
	rec := MembershipProducts{}
	err := mysqlRepo.Conn.First(&rec, id).Delete(&rec).Error
	if rec.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}