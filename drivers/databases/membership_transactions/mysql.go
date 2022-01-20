package membership_transactions

import (
	"gym-membership/business/membership_transactions"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlMembershipTransactionRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) membership_transactions.Repository {
	return &mysqlMembershipTransactionRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlMembershipTransactionRepo) Insert(data *membership_transactions.Domain) (membership_transactions.Domain, error) {
	domain := membership_transactions.Domain{}
	recTransaction := MembershipTransactions{}
	copier.Copy(&recTransaction, &data)
	err := mysqlRepo.Conn.Create(&recTransaction).Error
	mysqlRepo.Conn.Joins("MembershipProduct").Find(&recTransaction)
	if err != nil {
		return membership_transactions.Domain{}, err
	}
	copier.Copy(&domain, &recTransaction)
	domain.Nominal = recTransaction.MembershipProduct.Price
	return domain, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) GetAll(date time.Time, status string, idUser uint, offset, limit int) ([]membership_transactions.Domain, int64, error) {
	var totalData int64
	domain := []membership_transactions.Domain{}
	rec := []MembershipTransactions{}
	var err error
	if status != "" || idUser != 0 {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("MembershipProduct").Joins("Payment").
			Joins("User").Where("membership_transactions.created_at <= ?",date).Find(&rec, "status = ? OR user_id = ?", status, idUser).
			Count(&totalData).Error
	} else {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("User").
			Joins("MembershipProduct").Where("membership_transactions.created_at <= ?", date).Find(&rec).Count(&totalData).Error
	}

	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].UserName = rec[i].User.FullName
		domain[i].ProductName = rec[i].MembershipProduct.Name
		domain[i].Nominal = rec[i].MembershipProduct.Price
	}
	return domain, totalData, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) GetAllByUser(idUser uint) ([]membership_transactions.Domain, error) {
	domain := []membership_transactions.Domain{}
	rec := []MembershipTransactions{}
	var err error
	err = mysqlRepo.Conn.Order("updated_at desc").Joins("MembershipProduct").Joins("Payment").
		Joins("User").Find(&rec, "user_id = ?", idUser).Error

	if err != nil {
		return nil, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].UserName = rec[i].User.FullName
		domain[i].ProductName = rec[i].MembershipProduct.Name
		domain[i].Nominal = rec[i].MembershipProduct.Price
	}
	return domain, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) GetByID(idTransaction uint) (membership_transactions.Domain, error) {
	rec := MembershipTransactions{}
	domain := membership_transactions.Domain{}
	err := mysqlRepo.Conn.Order("updated_at desc").Joins("MembershipProduct").Joins("Payment").Joins("User").
		First(&rec, idTransaction).Error
	if err != nil {
		return membership_transactions.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	domain.UserName = rec.User.FullName
	domain.ProductName = rec.MembershipProduct.Name
	domain.Nominal = rec.MembershipProduct.Price
	return domain, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) UpdateStatus(id, idAdmin uint, status string) (membership_transactions.Domain, error) {
	rec := MembershipTransactions{}
	domain := membership_transactions.Domain{}
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", id).
		Updates(map[string]interface{}{"status": status, "admin_id": idAdmin}).Error
	if errUpdate != nil {
		return membership_transactions.Domain{}, errUpdate
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) UpdateReceipt(id uint, urlImage string) (membership_transactions.Domain, error) {
	rec := MembershipTransactions{}
	domain := membership_transactions.Domain{}
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", id).
		Updates(map[string]interface{}{"status": urlImage}).Error
	if errUpdate != nil {
		return membership_transactions.Domain{}, errUpdate
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) GetAllByUser(idUser uint) ([]membership_transactions.Domain, error) {
	domain := []membership_transactions.Domain{}
	rec := []MembershipTransactions{}
	err := mysqlRepo.Conn.Order("updated_at desc").Joins("MembershipProduct").Joins("Payment").Joins("User").
		Find(&rec, "user_id = ?", idUser).Error
	if err != nil {
		return []membership_transactions.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].UserName = rec[i].User.FullName
		domain[i].Nominal = rec[i].MembershipProduct.Price
		domain[i].ProductName = rec[i].MembershipProduct.Name
	}
	return domain, nil
}