package repository


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"shakilakhtar/go-microservices/types"
	//"go-microservices/utils"
	"shakilakhtar/go-microservices-platform/dataaccess"
)

type Repository interface {
	FindAll() []types.Asset
	Get(id int) (types.Asset, error)
	Save(item types.Asset) (types.Asset, error)
	Update(item types.Asset) (types.Asset, error)
}

type catalogRepository struct {
	db *gorm.DB
}

const (
	ASSET_TABLE = "asset"
)

func (repo *catalogRepository) FindAll() []types.Asset {
	var assets []types.Asset
	repo.db.Table(ASSET_TABLE).Find(&assets)
	return assets
}

func (repo *catalogRepository) Get(id int) (types.Asset, error) {
	var asset types.Asset
	repo.db.Table(ASSET_TABLE).Find(&asset, id)
	return asset, nil
}

func (repo *catalogRepository) Save(asset types.Asset) (types.Asset, error) {
	repo.db.Table(ASSET_TABLE).Create(&asset)
	return asset, nil
}

func (repo *catalogRepository) Update(asset types.Asset) (types.Asset, error) {
	repo.db.Table(ASSET_TABLE).Save(&asset)
	return asset, nil
}

// NewCatalogRepository returns a new instance of a asset catalog repository.
func NewCatalogRepository() Repository {
	return &catalogRepository{
		//db: utils.GetConnection(),
		db: dataaccess.GetConnection(),
	}
}