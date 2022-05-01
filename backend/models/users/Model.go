package users

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

func All() ([]*User, error) {
	var user []*User
	err := database.GetRecords(&user)
	return user, err
}

func Exists(cond ...interface{}) (bool, error) {
	return database.RecordExist(&User{}, cond)
}

func Find(id interface{}) (*User, error) {
	var user *User
	if db, err := database.Open(); err == nil {
		results := db.
			Where("id = ?", id).
			Find(&user)
		return user, results.Error
	} else {
		return nil, err
	}
}

func (user User) Create() error {
	return database.CreateRecord(&user)
}

func (user User) Update() error {
	userRecord, err := Find(user.ID)
	if err != nil {
		return err
	}

	hash, hash_err := HashPassword(user.Password)
	if hash_err != nil {
		return hash_err
	}

	userRecord.Username = user.Username
	userRecord.Password = hash
	userRecord.Email = user.Email

	return database.UpdateRecord(&userRecord)
}

func (user User) Delete() error {
	return database.DeleteRecord(&User{}, user.ID)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
