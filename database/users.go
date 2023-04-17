/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package database

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/aidarkhanov/nanoid"
	"golang.org/x/crypto/argon2"
	"log"
	"math/rand"
	"test/models"
	"time"
)

type Users struct {
	ID              string `gorm:"ID;PRIMARY_KEY;BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	UserLogin       string `gorm:"user_login;VARCHAR(60)"`
	UserPass        string `gorm:"user_pass;VARCHAR(255)"`
	UserFinApiPass  string `gorm:"user_finapi_pass;VARCHAR(500)"`
	UserNicename    string `gorm:"user_nicename;VARCHAR(50)"`
	UserEmail       string `gorm:"user_email;VARCHAR(100)"`
	ActivationCode  string `gorm:"activation_code;VARCHAR(255)"`
	UserRegistered  string `gorm:"user_registered;DATETIME"`
	IsVerified      int    `gorm:"is_verified;TINYINT(1)"`
	LastName        string `gorm:"last_name;VARCHAR(250)"`
	FirstName       string `gorm:"first_name;VARCHAR(250)"`
	UserCredentials string `gorm:"user_credentials;VARCHAR(250)"`
}

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type UserRepository interface {
	GetUsers() ([]models.Users, error)
	UpdateUserByID(userId string, field string, value string) error
	DeleteUserByID(userId string) error
	GetUserByID(userId string) ([]models.Users, error)
	CreateUser(
		userEmail string,
		userPass string,
		userNickname string,
		userName string,
		userLastName string,
	) ([]models.Users, error)
}

func (*Users) TableName() string {
	return "paygoal_app.users"
}

func (db *DB) CreateUser(userEmail string, userPass string, userNickname string, userName string, userLastName string) ([]models.Users, error) {
	var users []Users
	now := time.Now().UTC().Format("2006-01-02 03:04:05")
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	hash, passwordErr := GenerateFromPassword(userPass, p)
	if passwordErr != nil {
		log.Fatal(passwordErr)
	}

	// cipher key
	key := ""
	// plaintext
	alphabet := nanoid.DefaultAlphabet
	size := nanoid.DefaultSize

	id, nanoErr := nanoid.Generate(alphabet, size)
	if nanoErr != nil {
		panic(nanoErr)
	}

	encryptedPass := EncryptAES([]byte(key), id)

	err := db.db.
		Create(&models.Users{
			ID:              "",
			UserLogin:       userNickname,
			UserPass:        hash,
			UserFinApiPass:  encryptedPass,
			UserNicename:    userNickname,
			UserEmail:       userEmail,
			ActivationCode:  "",
			UserRegistered:  now,
			IsVerified:      0,
			LastName:        userLastName,
			FirstName:       userName,
			UserCredentials: "",
		}).Error
	if err != nil {
		return nil, err
	}

	var newUsers []models.Users
	for _, user := range users {
		newUsers = append(newUsers, models.Users{
			ID:              user.ID,
			UserLogin:       user.UserLogin,
			UserPass:        user.UserPass,
			UserFinApiPass:  user.UserFinApiPass,
			UserNicename:    user.UserNicename,
			UserEmail:       user.UserEmail,
			ActivationCode:  user.ActivationCode,
			UserRegistered:  user.UserRegistered,
			IsVerified:      user.IsVerified,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			UserCredentials: user.UserCredentials,
		})
	}

	return newUsers, err
}

func (db *DB) GetUsers() ([]models.Users, error) {
	var users []Users
	err := db.db.
		Select("paygoal_app.users.ID as ID, " +
			"paygoal_app.users.user_login as user_login, " +
			"paygoal_app.users.user_pass, " +
			"paygoal_app.users.user_finapi_pass, " +
			"paygoal_app.users.user_nicename, " +
			"paygoal_app.users.user_email, " +
			"paygoal_app.users.activation_code, " +
			"paygoal_app.users.user_registered, " +
			"paygoal_app.users.is_verified, " +
			"paygoal_app.users.last_name, " +
			"paygoal_app.users.first_name, " +
			"paygoal_app.users.user_credentials ").
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	var newUsers []models.Users
	for _, user := range users {
		newUsers = append(newUsers, models.Users{
			ID:              user.ID,
			UserLogin:       user.UserLogin,
			UserPass:        user.UserPass,
			UserFinApiPass:  user.UserFinApiPass,
			UserNicename:    user.UserNicename,
			UserEmail:       user.UserEmail,
			ActivationCode:  user.ActivationCode,
			UserRegistered:  user.UserRegistered,
			IsVerified:      user.IsVerified,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			UserCredentials: user.UserCredentials,
		})
	}

	return newUsers, nil
}

func (db *DB) GetUserByID(id string) ([]models.Users, error) {
	var users []Users
	err := db.db.
		Select("paygoal_app.users.ID as ID, " +
			"paygoal_app.users.user_login as user_login, " +
			"paygoal_app.users.user_pass, " +
			"paygoal_app.users.user_fin_api_pass, " +
			"paygoal_app.users.user_nicename, " +
			"paygoal_app.users.user_email, " +
			"paygoal_app.users.activation_code, " +
			"paygoal_app.users.user_registered, " +
			"paygoal_app.users.is_verified, " +
			"paygoal_app.users.last_name, " +
			"paygoal_app.users.first_name, " +
			"paygoal_app.users.user_credentials ").
		Where(fmt.Sprintf("paygoal_app.users.ID = %s", id)).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	var newUsers []models.Users
	for _, user := range users {
		newUsers = append(newUsers, models.Users{
			ID:              user.ID,
			UserLogin:       user.UserLogin,
			UserPass:        user.UserPass,
			UserFinApiPass:  user.UserFinApiPass,
			UserNicename:    user.UserNicename,
			UserEmail:       user.UserEmail,
			ActivationCode:  user.ActivationCode,
			UserRegistered:  user.UserRegistered,
			IsVerified:      user.IsVerified,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			UserCredentials: user.UserCredentials,
		})
	}

	return newUsers, err
}

func (db *DB) UpdateUserByID(id string, field string, value string) error {
	var users []Users
	err := db.db.Model(&users).
		Where(fmt.Sprintf("paygoal_app.users.ID = %s", id)).
		Update(fmt.Sprintf("paygoal_app.users.%s", field), value).
		Find(&users).Error
	if err != nil {
		return err
	}
	return err
}

func (db *DB) DeleteUserByID(id string) error {
	var users []Users
	err := db.db.
		Where(fmt.Sprintf("paygoal_app.users.ID = %s", id)).
		Delete(&users).
		Find(&users).Error
	if err != nil {
		return err
	}
	return err
}

func GenerateFromPassword(password string, p *params) (encodedHash string, err error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2i$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s
}
