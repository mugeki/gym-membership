package payment_accounts

import (
	paymentAccount "gym-membership/business/payment_accounts"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type mysqlPaymentAccountRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) paymentAccount.Repository {
	return &mysqlPaymentAccountRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlPaymentAccountRepo) GetAll() ([]paymentAccount.Domain, error) {
	domain := []paymentAccount.Domain{}
	rec := []PaymentAccount{}

	err := mysqlRepo.Conn.Find(&rec).Error
	if err != nil {
		return nil, err
	}

	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlPaymentAccountRepo) Insert(accountData *paymentAccount.Domain) (paymentAccount.Domain, error) {
	rec := PaymentAccount{}
	domain := paymentAccount.Domain{}
	copier.Copy(&rec, &accountData)

	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return paymentAccount.Domain{}, err
	}

	copier.Copy(&domain, &rec)
	return domain, nil
}
