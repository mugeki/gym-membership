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
	copier.Copy(&domain, &recUser)
	return domain, nil
}

func (mysqlRepo *mysqlClassRepo) GetAll(name string, offset, limit int) ([]class.Domain, int64, error) {
	var totalData int64
	domain := []class.Domain{}
	rec := []Class{}

	mysqlRepo.Conn.Find(&rec, "name LIKE ?", "%"+name+"%").Count(&totalData)
	err := mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").
		Joins("Trainers").Find(&rec).Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].TrainerName = rec[i].Trainers.Fullname
		domain[i].TrainerImage = rec[i].Trainers.UrlImage
	}

	return domain, totalData, nil
}

func (mysqlRepo *mysqlClassRepo) UpdateClassByID(id uint, classData *class.Domain) (class.Domain, error) {
	// println("cek id", id)
	domain := class.Domain{}
	rec := Class{}
	recData := Class{}
	copier.Copy(&recData, &classData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).Error
	if err != nil {
		return class.Domain{}, err
	}
	copier.Copy(domain, rec)
	return domain, nil
}

func (mysqlRepo *mysqlClassRepo) UpdateStatus(idClass int, status bool) (string, error) {
	rec := Class{}
	// println("update to false")
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idClass).Update("available_status", status).Error
	if errUpdate != nil {
		return "data not found", errUpdate
	}
	return "succes", nil
}

func (mysqlRepo *mysqlClassRepo) IsExist(idClass int) (class.Domain, error) {
	rec := Class{}
	domain := class.Domain{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", idClass).Error
	if err != nil {
		return class.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlClassRepo) UpdateParticipant(idClass int) (string, error) {
	// println("repo classes")
	rec := Class{}
	data, err := mysqlRepo.IsExist(idClass)
	if err != nil {
		return "data not found", err
	}
	kuotaUpdated := data.Participant + 1
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idClass).Update("participant", kuotaUpdated).Error
	if errUpdate != nil {
		return "errUpdated", err
	}
	if rec.Kuota == rec.Participant {
		mysqlRepo.UpdateStatus(idClass, false)
	}

	return "succes", nil
}