package database

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Keeping them as separate variables so they never can get mixed up
var DB *gorm.DB
var TEST_DB *gorm.DB

func Open() (db *gorm.DB, err error) {
	if os.Getenv("ENV") == "test" {
		if TEST_DB != nil {
			return TEST_DB, nil
		}
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		DB = db
		return db, err
	} else {
		if DB != nil {
			log.Printf("Using current mysql db connection")
			return DB, nil
		}
		log.Printf("Open new mysql db connection")
		db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
		DB = db
		return db, err
	}
}

func RecordExist(value interface{}, args ...interface{}) (bool, error) {
	db, err := Open()
	db.DB()
	if err == nil {
		var count int64
		db.Model(&value).Where(args).Count(&count)
		if count == 1 {
			return true, nil
		} else {
			return false, errors.New("row does not exist")
		}
	} else {
		return false, err
	}
}

func GetRecords(value interface{}) error {
	if db, err := Open(); err == nil {
		results := db.Find(value)
		return results.Error
	} else {
		return err
	}
}

func CreateRecord(value interface{}) error {
	if db, err := Open(); err == nil {
		return db.Create(value).Error
	} else {
		return err
	}
}

func GetRecord(value interface{}, query interface{}, args ...interface{}) error {
	if db, err := Open(); err == nil {
		return db.Where(query, args).First(value).Error
	} else {
		return err
	}
}

func DeleteRecord(value interface{}, id string) error {
	if db, err := Open(); err == nil {
		return db.Delete(value, id).Error
	} else {
		return err
	}
}

func DeleteMany(value interface{}, id []string) error {
	if db, err := Open(); err == nil {
		return db.Delete(value, id).Error
	} else {
		return err
	}
}

func UpdateRecord(value interface{}) error {
	if db, err := Open(); err == nil {
		return db.Save(value).Error
	} else {
		return err
	}
}
