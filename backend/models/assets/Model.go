package assets

import (
	"app/database"
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

func All() ([]*Asset, error) {
	var assets []*Asset
	if db, err := database.Open(); err == nil {
		results := db.Find(&assets)
		return assets, results.Error
	} else {
		return nil, err
	}
}

func (asset Asset) Create() error {
	return database.CreateRecord(&asset)
}

func (asset Asset) Delete() error {
	return database.DeleteRecord(&Asset{}, asset.ID)
}

func Find(id interface{}) (*Asset, error) {
	var asset *Asset
	err := database.GetRecord(&asset, "id = ?", id)
	return asset, err
}
