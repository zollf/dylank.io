package models

import (
	"app/database"
	"log"
	"time"
)

type Asset struct {
	ID        uint64    `json:"id"`
	Slug      string    `json:"slug" gorm:"index:idx_asset_slug,unique"`
	Title     string    `json:"title" gorm:"index:idx_asset_title,unique"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetAssets() ([]*Asset, error) {
	var assets []*Asset
	if db, err := database.Open(); err == nil {
		results := db.Find(&assets)
		return assets, results.Error
	} else {
		return nil, err
	}
}

func CreateAsset(asset *Asset) error {
	log.Printf("%v", asset)
	return database.CreateRecord(asset)
}

func DeleteAsset(id string) error {
	return database.DeleteRecord(&Asset{}, id)
}

func GetAsset(id string) (*Asset, error) {
	var asset *Asset
	err := database.GetRecord(&asset, "id = ?", id)
	return asset, err
}
