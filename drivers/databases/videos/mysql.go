package videos

import (
	"gym-membership/business/videos"
	"gym-membership/drivers/databases/classifications"

	"github.com/jinzhu/copier"
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
	domain := []videos.Domain{}
	rec := []Videos{}
	err := mysqlRepo.Conn.Joins("Classification").Find(&rec).Error
	if err != nil {
		return nil, err
	}
	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].ClassificationName = rec[i].Classification.Name
	}
	return domain, nil

}

func (mysqlRepo *mysqlVideosRepo) GetClassificationID(classification string) (int, error){
	rec := classifications.Classifications{}
	err := mysqlRepo.Conn.First(&rec, "name = ?", classification).Error
	if err != nil { 
		return -1, err
	}
	return rec.ID, nil
}

func (mysqlRepo *mysqlVideosRepo) Insert(videoData *videos.Domain) (videos.Domain, error){
	domain := videos.Domain{}
	rec := Videos{}
	copier.Copy(&rec, videoData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return videos.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlVideosRepo) UpdateByID(id uint, videoData *videos.Domain) (videos.Domain, error){
	domain := videos.Domain{}
	rec := Videos{}
	recData := Videos{}
	copier.Copy(&recData, videoData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).
							Update("member_only",recData.MemberOnly).Error
	if err != nil {
		return videos.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}
