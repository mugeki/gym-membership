package users_test

// import (
// 	"database/sql/driver"
// 	userBusiness "gym-membership/business/users"
// 	"gym-membership/drivers/databases/users"
// 	"regexp"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type AnyTime struct{}

// func (a AnyTime) Match(v driver.Value) bool {
// 	_, ok := v.(time.Time)
// 	return ok
// }

// var (
// 	userUUID = uuid.New()
// 	user = userBusiness.Domain{
// 		ID 			: 1,
// 		UUID 		: userUUID,
// 		Username	: "test123",
// 		Password	: "testpassword",
// 		Email		: "test@gmail.com",
// 		FullName 	: "Test Name",
// 		Gender 		: "Male",
// 		Telephone 	: "88888000102",
// 		Address 	: "Test Street",
// 	}
// )

// func TestInsert(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	gdb, _ := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	userRepo := users.NewMySQLRepo(gdb)
// 	defer db.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("INSERT INTO `users`").
// 		WithArgs(AnyTime{}, AnyTime{}, nil, user.UUID, user.Username, user.Password, user.Email,
// 			user.FullName, user.Gender, user.Telephone, user.Address, user.ID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	_, err = userRepo.Register(&user)
// 	require.NoError(t, err)
// }

// func TestGetByUsername(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	gdb, _ := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	userRepo := users.NewMySQLRepo(gdb)
// 	defer db.Close()

// 	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
// 		WithArgs(user.Username).
// 		WillReturnRows(
// 			sqlmock.NewRows([]string{
// 				"id",
// 				"created_at",
// 				"updated_at",
// 				"deleted_at",
// 				"uuid",
// 				"username",
// 				"password",
// 				"email",
// 				"full_name",
// 				"gender",
// 				"telephone",
// 				"address",
// 			}).AddRow(
// 				user.ID,
// 				time.Now(),
// 				time.Now(),
// 				nil,
// 				user.UUID,
// 				user.Username,
// 				user.Password,
// 				user.Email,
// 				user.FullName,
// 				user.Gender,
// 				user.Telephone,
// 				user.Address,
// 			))

// 	_, err = userRepo.GetByUsername(user.Username)
// 	require.NoError(t, err)
// }