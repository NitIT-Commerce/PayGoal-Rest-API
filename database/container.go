package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"test/services"
)

type Container struct {
	DB       *gorm.DB
	SqlCon   *DB
	Services struct {
		user *services.UserService
	}
}

func (c *Container) GetUserService() *services.UserService {
	if c.Services.user == nil {
		c.Services.user = services.NewUserService(c.GetMariaDb())
	}

	return c.Services.user
}

func (db *Container) GetMariaDb() *DB {
	if db.SqlCon == nil {
		maria := NewStorage(db.GetSqlConnection())
		db.SqlCon = maria
	}
	return db.SqlCon
}

func (db *Container) GetSqlConnection() *gorm.DB {
	if db.DB != nil {
		dsn := "root:123@tcp(localhost:6044)/paygoal_app?charset=utf8&parseTime=True&loc=Local"

		//dsn := "root:123@tcp(localhost:6044)/paygoal_app?charset=utf8&parseTime=True&loc=Local"
		gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("could not connect to database")
		}
		db.DB = gormdb
	}
	log.Println("Somethig went wrong5")

	return db.DB
}
