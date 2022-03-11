package utils

import (
	"app/models/assets"
	"time"
)

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

func GetAssetsData() ([]*AssetData, error) {
	assets, assets_err := assets.All()
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
	asset, asset_err := assets.Find(id)

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

func CheckAssetExistsInAssets(needle *assets.Asset, haystack []*assets.Asset) bool {
	for _, needleInHay := range haystack {
		if needleInHay.ID == needle.ID {
			return true
		}
	}
	return false
}