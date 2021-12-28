package members

import (
	"gym-membership/business/members"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type mysqlMembersRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) members.Repository {
	return &mysqlMembersRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlMembersRepo) Insert(userData *members.Domain) (members.Domain, error) {
	println("repo members", userData.Name)
	domain := members.Domain{}
	recUser := Members{}
	copier.Copy(&recUser, &userData)
	err := mysqlRepo.Conn.Create(&recUser).Error
	if err != nil {
		return members.Domain{}, err
	}
	copier.Copy(domain, recUser)
	return domain, nil
}

func (mysqlRepo *mysqlMembersRepo) GetByUserID(idMembers int) (string, error) {
	// println("repo members")
	rec := Members{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", idMembers).Error
	if err != nil {
		return "data not found", err
	}
	GetByUserID := rec.Participant + 1
	errUpdate := mysqlRepo.Conn.First(&rec, "id = ?", idMembers).Update("participant", GetByUserID)
	if errUpdate != nil {
		return "errUpdated", err
	}
	if rec.Members == rec.Participant {
		mysqlRepo.UpdateStatus(idMembers)
	}
	return "succes", nil
}
