package class

import (
	"gym-membership/business/class"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type mysqlClassRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) class.Repository {
	return &mysqlClassRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlClassRepo) Insert(userData *class.Domain) (class.Domain, error) {
	println("repo classes", userData.Name)
	domain := class.Domain{}
	recUser := Class{}
	copier.Copy(&recUser, &userData)
	err := mysqlRepo.Conn.Create(&recUser).Error
	if err != nil {
		return class.Domain{}, err
	}
	copier.Copy(domain, recUser)
	return domain, nil
}

func (mysqlRepo *mysqlClassRepo) UpdateStatus(idClass int) (string, error) {
	rec := Class{}
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idClass).Update("available_status", true).Error
	if errUpdate != nil {
		return "data not found", errUpdate
	}
	return "succes", nil
}

func (mysqlRepo *mysqlClassRepo) UpdateKuota(idClass int) (string, error) {
	// println("repo classes")
	rec := Class{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", idClass).Error
	if err != nil {
		return "data not found", err
	}
	kuotaUpdated := rec.Participant + 1
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idClass).Update("participant", kuotaUpdated)
	if errUpdate != nil {
		return "errUpdated", err
	}
	if rec.Kuota == rec.Participant {
		mysqlRepo.UpdateStatus(idClass)
	}

	return "succes", nil
}
