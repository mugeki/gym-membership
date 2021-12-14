package videos

import (
	"gym-membership/business/videos"

	"gorm.io/gorm"
)

type mysqlVideosRepo struct {
	Conn *gorm.DB
}

func NewMySQLRepo(conn *gorm.DB) videos.Repository {
	return &mysqlVideosRepo {
		Conn: conn,
	}
}

func (mysqlRepo *mysqlVideosRepo) GetAll() ([]videos.Domain, error){
	rec := []Videos{}
	err := mysqlRepo.Conn.Joins("Classification").Find(&rec).Error
	if err != nil {
		return nil, err
	}
	return toDomainArray(rec), nil

}

func (mysqlRepo *mysqlVideosRepo) GetClassificationID(classification string) (int, error){
	rec := VideoClassifications{}
	err := mysqlRepo.Conn.First(&rec, "name = ?", classification).Error
	if err != nil { 
		return -1, err
	}
	return rec.ID, nil
}

func (mysqlRepo *mysqlVideosRepo) Insert(videoData *videos.Domain) (videos.Domain, error){
	rec := fromDomain(*videoData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return videos.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlVideosRepo) UpdateByID(id uint, videoData *videos.Domain) (videos.Domain, error){
	rec := Videos{}
	recData := *fromDomain(*videoData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).Error
	if err != nil {
		return videos.Domain{}, err
	}
	return rec.toDomain(), nil
}
