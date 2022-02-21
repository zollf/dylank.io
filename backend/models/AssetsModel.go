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

type AssetData struct {
	ID        uint64
	Index     int
	Title     string
	Slug      string
	Url       string
	CreatedAt string
	UpdatedAt string
}

type ProjectAssetData struct {
	Title   string
	Slug    string
	Url     string
	Checked bool
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

func GetAssetsData() ([]*AssetData, error) {
	assets, assets_err := GetAssets()
	if assets_err != nil {
		return nil, assets_err
	}

	var asset_data []*AssetData
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, asset := range assets {
		asset_data = append(asset_data, &AssetData{
			ID:        asset.ID,
			Index:     i + 1,
			Url:       asset.Url,
			Title:     asset.Title,
			Slug:      asset.Slug,
			CreatedAt: asset.CreatedAt.In(zone).Format(time.RFC822),
			UpdatedAt: asset.UpdatedAt.In(zone).Format(time.RFC822),
		})
	}

	return asset_data, nil
}

func GetAssetData(id string) (*AssetData, error) {
	asset, asset_err := GetAsset(id)

	if asset_err != nil {
		return nil, asset_err
	}

	zone, _ := time.LoadLocation("Australia/Perth")

	asset_data := &AssetData{
		ID:        asset.ID,
		Title:     asset.Title,
		Url:       asset.Url,
		CreatedAt: asset.CreatedAt.In(zone).Format(time.RFC822),
		UpdatedAt: asset.UpdatedAt.In(zone).Format(time.RFC822),
	}

	return asset_data, nil
}

func CheckAssetExistsInAssets(needle *Asset, haystack []*Asset) bool {
	for _, needleInHay := range haystack {
		if needleInHay.ID == needle.ID {
			return true
		}
	}
	return false
}
