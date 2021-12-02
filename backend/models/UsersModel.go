package models

import (
	"app/database"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Email        string             `json:"email" bson:"email"`
	Locked       bool               `json:"locked" bson:"locked"`
	DateCreated  string             `json:"dateCreated" bson:"dateCreated"`
	DateUpdated  string             `json:"dateUpdated" bson:"dateUpdated"`
	LastLoggedIn string             `json:"lastLoggedIn" bson:"lastLoggedIn"`
}

const UserCol = "users"

func CreateOrEditUser(user *User) error {
	found, err := FindUser(user)
	if err != nil {
		return err
	}

	if found {
		return UpdateUser(user)
	} else {
		username_exist, username_err := FindUserByUsername(user)
		if username_err != nil {
			return username_err
		}

		if username_exist {
			return fmt.Errorf("Username already exists")
		}

		return CreateUser(user)
	}
}

func GetUsers() ([]*User, error) {
	var users []*User
	err := database.GetDocuments(UserCol, func(cur *mongo.Cursor) error {
		var user *User
		if err := cur.Decode(&user); err != nil {
			return err
		} else {
			users = append(users, user)
			return nil
		}
	})

	return users, err
}

func FindUser(user *User) (bool, error) {
	return database.DocumentExist(UserCol, bson.M{"_id": user.ID})
}

func FindUserByUsername(user *User) (bool, error) {
	return database.DocumentExist(UserCol, bson.M{"username": user.Username})
}

func UpdateUser(user *User) error {
	if hash, hash_err := HashPassword(user.Password); hash_err != nil {
		return hash_err
	} else {
		user.Password = hash
		return database.UpdateDocument(ProjCol, bson.M{"_id": user.ID}, bson.M{"$set": user})
	}
}

func CreateUser(user *User) error {
	if hash, hash_err := HashPassword(user.Password); hash_err != nil {
		return hash_err
	} else {
		user.Password = hash
		return database.CreateDocument(UserCol, user)
	}
}

func GetUser(id string) (*User, error) {
	if pid, pid_err := primitive.ObjectIDFromHex(id); pid_err != nil {
		return nil, pid_err
	} else {
		var user *User
		err := database.GetDocument(UserCol, bson.M{"_id": pid}, func(res *mongo.SingleResult) error {
			return res.Decode(&user)
		})
		return user, err
	}
}

func DeleteUser(id string) error {
	if pid, pid_err := primitive.ObjectIDFromHex(id); pid_err != nil {
		return pid_err
	} else {
		return database.DeleteDocument(UserCol, bson.M{"_id": pid})
	}
}

func GetUserWithPassword(username string, password string) (*User, error) {
	var user *User
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		result := client.Database("db").Collection("users").FindOne(ctx, bson.M{"username": username})
		var storedUser *User

		if result.Err() != nil {
			return result.Err()
		}

		err := result.Decode(&storedUser)

		if err != nil {
			return err
		}

		// Verify password
		pass_err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password))

		// no password error, then the user is correct
		if pass_err == nil {
			user = storedUser
		}

		return pass_err
	})

	return user, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
