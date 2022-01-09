package transactionClass

import (

	// _classRepo "gym-membership/drivers/databases/class"
	"gym-membership/business/class"
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
	mysqlRepo.Conn.Joins("Class").Find(&recTransaction)
	if err != nil {
		return transactionClass.Domain{}, err
	}
	copier.Copy(&domain, &recTransaction)
	domain.Nominal = recTransaction.Class.Price

	return domain, nil
}

func (mysqlRepo *mysqlTransactionClassRepo) GetAll(status string, idUser uint, offset, limit int) ([]transactionClass.Domain, int64, error) {
	var totalData int64
	domain := []transactionClass.Domain{}
	rec := []TransactionClass{}
	var err error
	if status != "" || idUser != 0 {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("Class").
			Find(&rec, "status = ? OR user_id = ?", status, idUser).Count(&totalData).Error
	} else {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("Class").
			Find(&rec).Count(&totalData).Error
	}

	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].Nominal = rec[i].Class.Price
	}
	return domain, totalData, nil
}

func (mysqlRepo *mysqlTransactionClassRepo) UpdateStatus(idTransactionClass, idAdmin uint, status string) (transactionClass.Domain, error) {
	rec := TransactionClass{}
	domain := transactionClass.Domain{}
	println("id trans : ", idTransactionClass, "idAdmin : ", idAdmin)
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idTransactionClass).Updates(map[string]interface{}{"status": status, "admin_id": idAdmin}).Error
	if errUpdate != nil {
		return transactionClass.Domain{}, errUpdate
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlTransactionClassRepo) GetActiveClass(idUser uint) ([]class.Domain, error) {
	rec := []TransactionClass{}
	domain := []transactionClass.Domain{}
	domainArrClass := []class.Domain{}

	err := mysqlRepo.Conn.Order("updated_at desc").Joins("Class").Find(&rec, "user_id = ? AND status = ?", idUser, "accepted").Error
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
