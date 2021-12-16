package articles

import (
	"gym-membership/business/articles"

	"gorm.io/gorm"
)

type mysqlVideosRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) articles.Repository {
	return &mysqlVideosRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlVideosRepo) GetAll() ([]articles.Domain, error) {
	rec := []Articles{}
	err := mysqlRepo.Conn.Joins("Classification").Find(&rec).Error
	if err != nil {
		return nil, err
	}
	return toDomainArray(rec), nil

}

func (mysqlRepo *mysqlVideosRepo) Insert(videoData *articles.Domain) (articles.Domain, error) {
	rec := fromDomain(*videoData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return articles.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlVideosRepo) UpdateByID(id uint, videoData *articles.Domain) (articles.Domain, error) {
	rec := Articles{}
	recData := *fromDomain(*videoData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).Error
	if err != nil {
		return articles.Domain{}, err
	}
	return rec.toDomain(), nil
}
