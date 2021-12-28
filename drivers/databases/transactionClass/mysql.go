package transactionClass

import (
	"gym-membership/business/class"
	// _classRepo "gym-membership/drivers/databases/class"
	"gym-membership/business/transactionClass"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlTransactionClassRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) transactionClass.Repository {
	return &mysqlTransactionClassRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlTransactionClassRepo) Insert(transactionClassData *transactionClass.Domain) (transactionClass.Domain, error) {
	domain := transactionClass.Domain{}
	recTransaction := TransactionClass{}
	copier.Copy(&recTransaction, &transactionClassData)
	err := mysqlRepo.Conn.Create(&recTransaction).Error
	mysqlRepo.Conn.Joins("Trainers").Find(&recTransaction)
	if err != nil {
		return transactionClass.Domain{}, err
	}
	copier.Copy(&domain, &recTransaction)
	domain.Nominal = recTransaction.Class.Price
	domain.Location = recTransaction.Class.Location

	return domain, nil
}

func (mysqlRepo *mysqlTransactionClassRepo) GetAll(offset, limit int) ([]transactionClass.Domain, int64, error) {
	var totalData int64
	domain := []transactionClass.Domain{}
	rec := []TransactionClass{}

	err := mysqlRepo.Conn.Find(&rec).Count(&totalData).Error
	// err := mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").
	// 	Joins("Trainers").Find(&rec).Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	return domain, totalData, nil
}

func (mysqlRepo *mysqlTransactionClassRepo) GetActiveClass(idUser uint) ([]class.Domain, error) {
	domainResult := []class.Domain{}
	domainClass := class.Domain{}
	rec := []TransactionClass{}
	status := "accepted"
	err := mysqlRepo.Conn.Find(&rec, "user_id = ? AND status = ?", idUser, status).Error
	if err != nil {
		return []class.Domain{}, err
	}
	mysqlRepo.Conn.Order("updated_at desc").Joins("Class").Find(&rec)
	for i := 0; i < len(rec); i++ {
		copier.Copy(&domainClass, &rec[i].Class)
		domainResult = append(domainResult, domainClass)
	}

	return domainResult, nil
}

func (mysqlRepo *mysqlTransactionClassRepo) UpdateStatus(id uint, status string) (transactionClass.Domain, error) {
	rec := TransactionClass{}
	domain := transactionClass.Domain{}
	// println("update to false")
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", id).Update("status", status).Error
	if errUpdate != nil {
		return transactionClass.Domain{}, errUpdate
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}
