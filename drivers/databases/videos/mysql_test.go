package videos_test

import (
	"database/sql/driver"
	videoBusiness "gym-membership/business/videos"
	"gym-membership/drivers/databases/videos"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var (
	video = videoBusiness.Domain{
		Title              : "Test Video",
		ClassificationID   : 1,
		ClassificationName : "test classification",
		AdminID            : 1,
		MemberOnly         : false,
		Url                : "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	videoRepo := videos.NewMySQLRepo(gdb)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT `videos`")).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"title",
				"classification",
				"adminId",
				"memberOnly",
				"url",
			}).AddRow(
				1,
				video.Title,
				video.ClassificationName,
				video.AdminID,
				video.MemberOnly,
				video.Url,
			))

	_, err = videoRepo.GetAll("Test",0,10)
	require.NoError(t, err)
}

func TestGetClassificationID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	videoRepo := videos.NewMySQLRepo(gdb)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `classifications`")).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"name",
			}).AddRow(
				video.ClassificationID,
				video.ClassificationName,
			))

	_, err = videoRepo.GetClassificationID(video.ClassificationName)
	require.NoError(t, err)
}
func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	videoRepo := videos.NewMySQLRepo(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `videos`").
		WithArgs(AnyTime{}, AnyTime{}, nil, video.Title, video.ClassificationID, 
			video.AdminID, video.MemberOnly, video.Url).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = videoRepo.Insert(&video)
	require.NoError(t, err)
}

func TestUpdateByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	videoRepo := videos.NewMySQLRepo(gdb)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `videos`")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"title",
				"classification",
				"adminId",
				"memberOnly",
				"url",
			}).AddRow(
				1,
				video.Title,
				video.ClassificationName,
				video.AdminID,
				video.MemberOnly,
				video.Url,
			))

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `videos` SET").
		WithArgs(AnyTime{}, video.Title, video.ClassificationID, 
			video.AdminID, video.Url, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `videos` SET").
		WithArgs(video.MemberOnly, AnyTime{}, 1, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	

	_, err = videoRepo.UpdateByID(1, &video)
	require.NoError(t, err)
}