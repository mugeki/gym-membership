package class_transactions

import (
	"gym-membership/business/class"
	"gym-membership/business/class_transactions"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlClassTransactionRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) class_transactions.Repository {
	return &mysqlClassTransactionRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlClassTransactionRepo) Insert(transactionClassData *class_transactions.Domain) (class_transactions.Domain, error) {
	domain := class_transactions.Domain{}
	recTransaction := ClassTransaction{}
	copier.Copy(&recTransaction, &transactionClassData)
	err := mysqlRepo.Conn.Create(&recTransaction).Error
	mysqlRepo.Conn.Joins("Class").Find(&recTransaction)
	if err != nil {
		return class_transactions.Domain{}, err
	}
	copier.Copy(&domain, &recTransaction)
	domain.Nominal = recTransaction.Class.Price

	return domain, nil
}

func (mysqlRepo *mysqlClassTransactionRepo) GetAll(date time.Time, status string, idUser uint, offset, limit int) ([]class_transactions.Domain, int64, error) {
	var totalData int64
	domain := []class_transactions.Domain{}
	rec := []ClassTransaction{}
	var err error
	if status != "" || idUser != 0 {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("Class").Joins("Payment").
			Joins("User").Where("class_transactions.created_at <= ?", date).Find(&rec, "status = ? OR user_id = ?", status, idUser).
			Count(&totalData).Error
	} else {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("Class").Joins("Payment").
			Joins("User").Where("class_transactions.created_at <= ?", date).Find(&rec).Count(&totalData).Error
	}

	if err != nil {
		return nil, 0, err
	}
	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].UserName = rec[i].User.FullName
		domain[i].ClassName = rec[i].Class.Name
		domain[i].Nominal = rec[i].Class.Price
	}
	return domain, totalData, nil
}

func (mysqlRepo *mysqlClassTransactionRepo) UpdateStatus(idClassTransaction, idAdmin uint, status string) (class_transactions.Domain, error) {
	rec := ClassTransaction{}
	domain := class_transactions.Domain{}
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idClassTransaction).
		Updates(map[string]interface{}{"status": status, "admin_id": idAdmin}).Error
	if errUpdate != nil {
		return class_transactions.Domain{}, errUpdate
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlClassTransactionRepo) UpdateReceipt(idClassTransaction uint, urlImage string) (class_transactions.Domain, error) {
	rec := ClassTransaction{}
	domain := class_transactions.Domain{}
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idClassTransaction).
		Updates(map[string]interface{}{"url_image_of_receipt": urlImage}).Error
	if errUpdate != nil {
		return class_transactions.Domain{}, errUpdate
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlClassTransactionRepo) GetActiveClass(idUser uint) ([]class.Domain, error) {
	rec := []ClassTransaction{}
	domain := []class.Domain{}

	err := mysqlRepo.Conn.Order("updated_at desc").Joins("Class").
		Find(&rec, "user_id = ? AND status = ?", idUser, "accepted").Error
	if err != nil {
		return []class.Domain{}, err
	}

	for i := 0; i < len(rec); i++ {
		copier.Copy(&domain, &rec[i].Class)
	}

	return domain, nil
}

func (mysqlRepo *mysqlClassTransactionRepo) GetByID(idTransaction uint) (class_transactions.Domain, error) {
	rec := ClassTransaction{}
	domain := class_transactions.Domain{}

	err := mysqlRepo.Conn.Order("updated_at desc").Joins("Class").Joins("Payment").Joins("User").
		First(&rec, idTransaction).Error
	if err != nil {
		return class_transactions.Domain{}, err
	}

	copier.Copy(&domain, &rec)
	domain.UserName = rec.User.FullName
	domain.ClassName = rec.Class.Name
	domain.Nominal = rec.Class.Price

	return domain, nil
}

func (mysqlRepo *mysqlClassTransactionRepo) GetAllByUser(idUser uint) ([]class_transactions.Domain, error) {
	domain := []class_transactions.Domain{}
	rec := []ClassTransaction{}
	err := mysqlRepo.Conn.Order("updated_at desc").Joins("Class").Joins("Payment").Joins("User").
		Find(&rec, "user_id = ?", idUser).Error
	if err != nil {
		return []class_transactions.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].UserName = rec[i].User.FullName
		domain[i].Nominal = rec[i].Class.Price
		domain[i].ClassName = rec[i].Class.Name
	}
	return domain, nil
}
