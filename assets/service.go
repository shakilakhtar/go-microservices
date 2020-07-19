package assets

import (
	"shakilakhtar/go-microservices/types"
	"shakilakhtar/go-microservices/repository"
)

//A service endpoint interface
type Service interface {
	GetAsset(id int) (types.Asset, error)
	ListAssets() []types.Asset
	Save(asset types.Asset) (types.Asset, error)
	Update(asset types.Asset) (types.Asset, error)
}

type service struct {
	catalogRepo repository.Repository
}

func (s *service) GetAsset(id int) (types.Asset, error) {
	asset, err := s.catalogRepo.Get(id)
	return asset, err
}

func (s *service) ListAssets() []types.Asset {
	assets := s.catalogRepo.FindAll()
	return assets
}

func (s *service) Save(asset types.Asset)(types.Asset, error) {
	s.catalogRepo.Save(asset)
	return asset, nil
}

func (s *service) Update(asset types.Asset) (types.Asset, error) {
	s.catalogRepo.Update(asset)
	return asset, nil
}

// NewService creates a inventory service with necessary dependencies.
func NewService() Service {
	return &service{
		catalogRepo: repository.NewCatalogRepository(),
	}
}
