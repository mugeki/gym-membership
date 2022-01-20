package videos

import (
	"gym-membership/business/videos"

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

func (mysqlRepo *mysqlVideosRepo) GetAll(title string, offset, limit int) ([]videos.Domain, int64, error){
	var totalData int64
	domain := []videos.Domain{}
	rec := []Videos{}

	mysqlRepo.Conn.Find(&rec, "title LIKE ?", "%"+title+"%").Count(&totalData)
	err := mysqlRepo.Conn.Limit(limit).Offset(offset).Order("updated_at desc").
						Joins("Classification").Find(&rec, "title LIKE ?", "%"+title+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].ClassificationName = rec[i].Classification.Name
	}
	return domain, totalData, nil
}

func (mysqlRepo *mysqlVideosRepo) GetByID(id uint) (videos.Domain, error) {
	domain := videos.Domain{}
	rec := Videos{}
	err := mysqlRepo.Conn.Joins("Classification").First(&rec, id).Error
	if err != nil {
		return videos.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	domain.ClassificationName = rec.Classification.Name
	return domain, nil
}

func (mysqlRepo *mysqlVideosRepo) Insert(videoData *videos.Domain) (videos.Domain, error){
	domain := videos.Domain{}
	rec := Videos{}
	copier.Copy(&rec, videoData)

	err := mysqlRepo.Conn.Create(&rec).Joins("Classification").First(&rec).Error
	if err != nil {
		return videos.Domain{}, err
	}

	copier.Copy(&domain, &rec)
	domain.ClassificationName = rec.Classification.Name
	return domain, nil
}

func (mysqlRepo *mysqlVideosRepo) UpdateByID(id uint, videoData *videos.Domain) (videos.Domain, error){
	domain := videos.Domain{}
	rec := Videos{}
	recData := Videos{}
	copier.Copy(&recData, videoData)

	err := mysqlRepo.Conn.Joins("Classification").First(&rec, "videos.id = ?", id).
		Updates(recData).Update("member_only",recData.MemberOnly).Error
	if err != nil {
		return videos.Domain{}, err
	}

	copier.Copy(&domain, &rec)
	domain.ClassificationName = rec.Classification.Name
	return domain, nil
}

func (mysqlRepo *mysqlVideosRepo) DeleteByID(id uint) error {
	rec := Videos{}
	err := mysqlRepo.Conn.First(&rec,id).Delete(&rec).Error
	if rec.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	if err != nil{
		return err
	}
	return nil
}