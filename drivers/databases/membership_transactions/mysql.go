package membership_transactions

import (
	"gym-membership/business/membership_transactions"

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
	mysqlRepo.Conn.Joins("MembershipProducts").Find(&recTransaction)
	if err != nil {
		return membership_transactions.Domain{}, err
	}
	copier.Copy(&domain, &recTransaction)
	domain.Nominal = recTransaction.MembershipProduct.Price
	return domain, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) GetAll(status string, idUser uint, offset, limit int) ([]membership_transactions.Domain, int64, error) {
	var totalData int64
	domain := []membership_transactions.Domain{}
	rec := []MembershipTransactions{}
	var err error
	if status != "" || idUser != 0 {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").Joins("MembershipProducts").
			Find(&rec, "status = ? OR user_id = ?", status, idUser).Count(&totalData).Error
	} else {
		err = mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").
			Joins("MembershipProducts").Find(&rec).Count(&totalData).Error
	}

	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].Nominal = rec[i].MembershipProduct.Price
	}
	return domain, totalData, nil
}

func (mysqlRepo *mysqlMembershipTransactionRepo) UpdateStatus(id uint, status string) (membership_transactions.Domain, error) {
	rec := MembershipTransactions{}
	domain := membership_transactions.Domain{}
	errUpdate := mysqlRepo.Conn.Joins("MembershipProducts").First(&rec, "id = ?", id).Update("status", status).Error
	if errUpdate != nil {
		return membership_transactions.Domain{}, errUpdate
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}
