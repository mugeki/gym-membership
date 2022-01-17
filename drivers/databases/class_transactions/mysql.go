package class_transactions

import (
	"fmt"
	"gym-membership/business/class"
	"gym-membership/business/class_transactions"

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

func (mysqlRepo *mysqlClassTransactionRepo) GetAll(status string, idUser uint, offset, limit int) ([]class_transactions.Domain, int64, error) {
	var totalData int64
	domain := []class_transactions.Domain{}
	rec := []ClassTransaction{}
	var err error
	if status != "" || idUser != 0 {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("Class").Joins("Payment").
			Find(&rec, "status = ? OR user_id = ?", status, idUser).Count(&totalData).Error
	} else {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("Class").Joins("Payment").
			Find(&rec).Count(&totalData).Error
	}

	if err != nil {
		return nil, 0, err
	}
	copier.Copy(&domain, &rec)
	fmt.Println("len", len(rec))
	for i := 0; i < len(rec); i++ {
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
	domain := []class_transactions.Domain{}
	domainArrClass := []class.Domain{}

	err := mysqlRepo.Conn.Order("updated_at desc").Joins("Class").
		Find(&rec, "user_id = ? AND status = ?", idUser, "accepted").Error
	if err != nil {
		return []class.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domainClass := class.Domain{}
		copier.Copy(&domainClass, &rec[i].Class)
		domainArrClass = append(domainArrClass, domainClass)
	}
	return domainArrClass, nil
}
