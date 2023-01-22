/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package database

import (
	"gorm.io/gorm"
	"log"
)

type DB struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *DB {
	log.Println("Somethig went wrong3")

	return &DB{db: db}
}
