/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package tests

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"test/database"
	"testing"
)

func initCallMockRepo(t *testing.T) (*database.DB, sqlmock.Sqlmock) {
	sqlDB, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	repo := database.NewStorage(gormDB)
	return repo, mock
}

func TestMariaDb_GetUser_CorrectTest(t *testing.T) {
	sqlDb, mock := initCallMockRepo(t)
	queryUser := "SELECT paygoal_app.users.ID as ID," +
		" paygoal_app.users.user_login as user_login," +
		" paygoal_app.users.user_pass," +
		" paygoal_app.users.user_finapi_pass," +
		" paygoal_app.users.user_nicename," +
		" paygoal_app.users.user_email," +
		" paygoal_app.users.activation_code," +
		" paygoal_app.users.user_registered," +
		" paygoal_app.users.is_verified," +
		" paygoal_app.users.last_name," +
		" paygoal_app.users.first_name," +
		" paygoal_app.users.user_credentials" +
		" FROM `paygoal_app`.`users` WHERE paygoal_app.users.ID = 1"

	userRows := mock.NewRows([]string{
		"ID", "user_login", "user_pass", "user_finapi_pass", "user_nicename", "user_email", "activation_code", "user_registered", "is_verified", "last_name", "first_name", "user_credentials",
	})
	userRows.AddRow([]driver.Value{
		"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
	}...)

	mock.ExpectQuery(regexp.QuoteMeta(queryUser)).WillReturnRows(userRows)

	_, err := sqlDb.GetUserByID("1")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMariaDb_GetUser_EmptyTest(t *testing.T) {
	sqlDb, mock := initCallMockRepo(t)
	queryUser := "SELECT paygoal_app.users.ID as ID," +
		" paygoal_app.users.user_login as user_login," +
		" paygoal_app.users.user_pass," +
		" paygoal_app.users.user_finapi_pass," +
		" paygoal_app.users.user_nicename," +
		" paygoal_app.users.user_email," +
		" paygoal_app.users.activation_code," +
		" paygoal_app.users.user_registered," +
		" paygoal_app.users.is_verified," +
		" paygoal_app.users.last_name," +
		" paygoal_app.users.first_name," +
		" paygoal_app.users.user_credentials" +
		" FROM `paygoal_app`.`users` WHERE paygoal_app.users.ID = 1"

	userRows := mock.NewRows([]string{
		"ID", "user_login", "user_pass", "user_finapi_pass", "user_nicename", "user_email", "activation_code", "user_registered", "is_verified", "last_name", "first_name", "user_credentials",
	})
	userRows.AddRow([]driver.Value{
		"", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
	}...)

	mock.ExpectQuery(regexp.QuoteMeta(queryUser)).WillReturnRows(userRows)

	_, err := sqlDb.GetUserByID("1")
	assert.Empty(t, err)
	assert.Empty(t, mock.ExpectationsWereMet())
}

// TODO
func TestMariaDb_CreateUser_CorrectTest(t *testing.T) {

	sqlDb, mock := initCallMockRepo(t)
	queryUser := "INSERT INTO users ID, user_login, user_pass, user_fin_api_pass, user_nicename, user_email, activation_code, user_registered, is_verified, last_name, first_name, user_credentials VALUES ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?"

	mock.ExpectQuery(regexp.QuoteMeta(queryUser))

	_, err := sqlDb.CreateUser("test@gmail.com", "test", "muster", "muster", "mann")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestMariaDb_CreateUser_WrongTest(t *testing.T) {
	sqlDb, mock := initCallMockRepo(t)
	//TODO

	_, err := sqlDb.CreateUser("test@gmail.com", "test", "muster", "muster", "mann")
	assert.Empty(t, err)
	assert.Empty(t, mock.ExpectationsWereMet())
}
