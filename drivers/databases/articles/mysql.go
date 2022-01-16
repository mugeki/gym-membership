package articles

import (
	"gym-membership/business/articles"

	"github.com/jinzhu/copier"
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

func (mysqlRepo *mysqlArticlesRepo) GetAll(title string, offset, limit int) ([]articles.Domain, int64, error) {
	var totalData int64
	domain := []articles.Domain{}
	rec := []Articles{}

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

func (mysqlRepo *mysqlArticlesRepo) GetByID(id uint) (articles.Domain, error) {
	domain := articles.Domain{}
	rec := Articles{}
	err := mysqlRepo.Conn.Joins("Classification").First(&rec, id).Error
	if err != nil {
		return articles.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	domain.ClassificationName = rec.Classification.Name
	return domain, nil
}

func (mysqlRepo *mysqlArticlesRepo) Insert(videoData *articles.Domain) (articles.Domain, error) {
	domain := articles.Domain{}
	rec := Articles{}

	copier.Copy(&rec, videoData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return articles.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return domain, nil
}

func (mysqlRepo *mysqlArticlesRepo) UpdateByID(id uint, videoData *articles.Domain) (articles.Domain, error) {
	domain := articles.Domain{}
	rec := Articles{}
	recData := Articles{}
	copier.Copy(recData, videoData)
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Updates(recData).
		Update("member_only",recData.MemberOnly).Error
	if err != nil {
		return articles.Domain{}, err
	}
	copier.Copy(domain, rec)
	return domain, nil
}

func (mysqlRepo *mysqlArticlesRepo) DeleteByID(id uint) error {
	rec := Articles{}
	err := mysqlRepo.Conn.First(&rec, id).Delete(&rec).Error
	if rec.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
