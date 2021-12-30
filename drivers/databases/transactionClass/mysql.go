package transactionClass

import (

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
	mysqlRepo.Conn.Joins("Class").Find(&recTransaction)
	if err != nil {
		return transactionClass.Domain{}, err
	}
	copier.Copy(&domain, &recTransaction)
	domain.Nominal = recTransaction.Class.Price
	// domain.Location = recTransaction.Class.Location

	return domain, nil
}

func (mysqlRepo *mysqlTransactionClassRepo) GetAll(status string, idUser uint, offset, limit int) ([]transactionClass.Domain, int64, error) {
	var totalData int64
	domain := []transactionClass.Domain{}
	rec := []TransactionClass{}
	var err error
	if status != "" || idUser != 0 {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).
			Find(&rec, "status = ? OR user_id = ?", status, idUser).Count(&totalData).Order("updated_at desc").Error
	} else {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).
			Find(&rec).Count(&totalData).Order("updated_at desc").Error
	}

	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	return domain, totalData, nil
}

// func (mysqlRepo *mysqlTransactionClassRepo) GetActiveClass(idUser uint) ([]class.Domain, error) {
// 	println("repo user id", idUser)
// 	domainResult := []class.Domain{}
// 	domainClass := class.Domain{}
// 	rec := []TransactionClass{}
// 	status := "accepted"
// 	err := mysqlRepo.Conn.Joins("Class").Find(&rec, "user_id = ? AND status = ?", idUser, status).Error
// 	if err != nil {
// 		return []class.Domain{}, err
// 	}
// 	println("repo interface", len(rec))
// 	// mysqlRepo.Conn.Order("updated_at desc").Joins("Class").Find(&rec)
// 	for i := 0; i < len(rec); i++ {
// 		copier.Copy(&domainClass, &rec[i].Class)
// 		domainResult = append(domainResult, domainClass)
// 	}

// 	return domainResult, nil
// }

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
