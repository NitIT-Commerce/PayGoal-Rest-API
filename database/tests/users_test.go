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
	queryUser := "SELECT ID FROM users"

	cdrRows := mock.NewRows([]string{
		"ID",
	})
	cdrRows.AddRow([]driver.Value{
		"1",
	}...)

	mock.ExpectQuery(regexp.QuoteMeta(queryUser)).WillReturnRows(cdrRows)

	_, err := sqlDb.GetUserByID("1")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
