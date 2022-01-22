package admins

import (
	"gym-membership/business/admins"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlAdminsRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) admins.Repository {
	return &mysqlAdminsRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlAdminsRepo) Register(adminData *admins.Domain) (admins.Domain, error) {
	domain := admins.Domain{}
	rec := Admins{}
	copier.Copy(&rec, &adminData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return admins.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlAdminsRepo) GetByUsername(username string) (admins.Domain, error) {
	domain := admins.Domain{}
	rec := Admins{}

	err := mysqlRepo.Conn.Where("username = ?", username).First(&rec).Error
	if err != nil {
		return admins.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlAdminsRepo) GetByID(id uint) (admins.Domain, error) {
	domain := admins.Domain{}
	rec := Admins{}

	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return admins.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlAdminsRepo) Update(id uint, adminData *admins.Domain) (admins.Domain, error) {
	domain := admins.Domain{}
	rec := Admins{}
	recData := Admins{}
	copier.Copy(&recData, &adminData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).
		Update("is_super_admin", adminData.IsSuperAdmin).Error
	if err != nil {
		return admins.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlAdminsRepo) GetAll(name string, offset, limit int) ([]admins.Domain, int64, error) {
	var totalData int64
	domain := []admins.Domain{}
	rec := []Admins{}
	mysqlRepo.Conn.Find(&rec, "full_name LIKE ?", "%"+name+"%").Count(&totalData)
	err := mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").
		Find(&rec, "full_name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, 0, err
	}
	
	copier.Copy(&domain, &rec)
	return domain, totalData, nil
}

func (mysqlRepo *mysqlAdminsRepo) DeleteByID(id uint) (error) {
	rec := Admins{}
	err := mysqlRepo.Conn.First(&rec, id).Delete(&rec).Error
	if rec.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}