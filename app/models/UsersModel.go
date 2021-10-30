package models

import (
	"app/database"
	"context"

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
		return CreateUser(user)
	}
}

func GetUsers() ([]*User, error) {
	var users []*User
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		cur, err := client.Database("db").Collection(UserCol).Find(ctx, bson.D{})

		if err != nil {
			return err
		}

		for cur.Next(ctx) {
			var result *User
			err := cur.Decode(&result)

			if err != nil {
				return err
			}

			users = append(users, result)
		}

		return nil
	})

	return users, err
}

func GetUser(id string) (*User, error) {
	var user *User
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		pid, pid_err := primitive.ObjectIDFromHex(id)

		if pid_err != nil {
			return pid_err
		}

		result := client.Database("db").Collection(UserCol).FindOne(ctx, bson.M{"_id": pid})

		if result.Err() != nil {
			return result.Err()
		}

		return result.Decode(&user)
	})

	return user, err
}

func FindUser(user *User) (bool, error) {
	var found bool = false
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		count, err := client.Database("db").Collection(UserCol).CountDocuments(ctx, bson.M{"_id": user.ID})

		if err != nil {
			return err
		}

		found = count != 0
		return nil
	})

	return found, err
}

func CreateUser(user *User) error {
	return database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		hash, hash_err := HashPassword(user.Password)

		if hash_err != nil {
			return hash_err
		}

		user.Password = hash

		_, err := client.Database("db").Collection(UserCol).InsertOne(ctx, user)
		return err
	})
}

func UpdateUser(user *User) error {
	return database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		hash, hash_err := HashPassword(user.Password)

		if hash_err != nil {
			return hash_err
		}

		user.Password = hash

		_, err := client.Database("db").Collection(UserCol).UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
		return err
	})
}

func DeleteUser(id string) error {
	return database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		pid, pid_err := primitive.ObjectIDFromHex(id)

		if pid_err != nil {
			return pid_err
		}

		_, err := client.Database("db").Collection(UserCol).DeleteOne(ctx, bson.M{"_id": pid})

		return err
	})
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
