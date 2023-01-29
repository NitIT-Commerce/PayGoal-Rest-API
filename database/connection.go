/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db     *gorm.DB
	DB     *gorm.DB
	SqlCon *DB
}

func NewStorage(db *gorm.DB) *DB {
	return &DB{db: db}
}

func (db *DB) GetMariaDb() *DB {
	if db.SqlCon == nil {
		maria := NewStorage(db.GetConnection())
		db.SqlCon = maria
	}
	return db.SqlCon
}

func (db *DB) GetConnection() *gorm.DB {
	if db.DB == nil {
		dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "", "", "", "", "")
		gormDb, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dbUri,
		}), &gorm.Config{})
		if err != nil {
			panic("could not connect to database")
		}
		db.DB = gormDb
	}
	return db.DB
}
