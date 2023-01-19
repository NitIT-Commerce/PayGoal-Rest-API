/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package utils

import (
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *DB {
	return &DB{db: db}
}
