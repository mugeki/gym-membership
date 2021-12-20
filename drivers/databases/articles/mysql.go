package articles

import (
	"gym-membership/business/articles"

	"gorm.io/gorm"
)

type mysqlArticlesRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) articles.Repository {
	return &mysqlArticlesRepo{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlArticlesRepo) GetAll() ([]articles.Domain, error) {
	rec := []Articles{}
	err := mysqlRepo.Conn.Find(&rec).Error
	if err != nil {
		return nil, err
	}
	return toDomainArray(rec), nil

}

func (mysqlRepo *mysqlArticlesRepo) Insert(videoData *articles.Domain) (articles.Domain, error) {
	rec := fromDomain(*videoData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return articles.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlArticlesRepo) UpdateByID(id uint, videoData *articles.Domain) (articles.Domain, error) {
	// println("cek id", id)
	rec := Articles{}
	recData := *fromDomain(*videoData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).Error
	if err != nil {
		return articles.Domain{}, err
	}
	return rec.toDomain(), nil
}
