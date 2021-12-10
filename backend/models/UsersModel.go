package models

import (
	"app/database"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint64     `json:"id"`
	Username     string     `json:"username" gorm:"index:idx_username,unique"`
	Password     string     `json:"password"`
	Email        string     `json:"email"`
	Locked       bool       `json:"locked"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	LastLoggedIn *time.Time `json:"lastLoggedIn"`
}

func GetUsers() ([]*User, error) {
	var users []*User
	err := database.GetRecords(&users)
	return users, err
}

func FindUser(user *User) (bool, error) {
	return database.RecordExist(&User{}, "id = ?", user.ID)
}

func FindUserByUsername(user *User) (bool, error) {
	return database.RecordExist(&User{}, "username = ?", user.Username)
}

func UpdateUser(user *User, id string) error {
	if hash, hash_err := HashPassword(user.Password); hash_err != nil {
		return hash_err
	} else {

		userRecord, err := GetUser(id)
		if err != nil {
			return err
		}

		userRecord.Username = user.Username
		userRecord.Password = hash
		userRecord.Email = user.Email

		return database.UpdateRecord(&userRecord)
	}
}

func CreateUser(user *User) error {
	if hash, hash_err := HashPassword(user.Password); hash_err != nil {
		return hash_err
	} else {
		user.Password = hash
		return database.CreateRecord(user)
	}
}

func GetUser(id string) (*User, error) {
	var user *User
	err := database.GetRecord(&user, "id = ?", id)
	return user, err
}

func DeleteUser(id string) error {
	return database.DeleteRecord(&User{}, id)
}

func GetUserWithPassword(username string, password string) (*User, error) {
	var user *User
	err := database.GetRecord(&user, "username = ?", username)

	if err != nil {
		return nil, err
	}

	pass_err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if pass_err != nil {
		return nil, pass_err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
