/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Helper struct {
	SqlCon *DB
	DB     *gorm.DB
}

func (c *Helper) GetMariaDb() *DB {
	if c.SqlCon == nil {
		maria := NewStorage(c.GetSqlConnection())
		c.SqlCon = maria
	}
	return c.SqlCon
}

func (c *Helper) GetSqlConnection() *gorm.DB {
	if c.DB != nil {
		dsn := "root:JAN3zLLVTrhtD9@tcp(195.201.32.138:40000)/paygoal_app?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("could not connect to database")
		} else {
			log.Print("Success connect to database")
		}
		c.DB = db
	}
	return c.DB
}
