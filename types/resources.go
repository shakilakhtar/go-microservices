package types


type (
	//For GET - /catalog/v1/assets
	ListAssetsRequest struct{}

	ListAssetsResponse struct {
		Assets []Asset `json:"assets,omitempty"`
		Err   error  `json:"error,omitempty"`
	}

	//For Post - /catalog/v1/assets/{assetId}
	GetAssetRequest struct {
		Id int `json:"id,omitempty"`
	}

	GetAssetResponse struct {
		Asset Asset  `json:"asset,omitempty"`
		Err  error `json:"error,omitempty"`
	}

	//For POST - /catalog/v1/assets
	AddAssetRequest struct {
		Data Asset `json:"data,omitempty"`
	}

	AddAssetResponse struct {
		Id  int   `json:"id,omitempty"`
		Err error `json:"error,omitempty"`
	}

	//For POST - /catalog/v1/assets
	UpdateAssetRequest struct {
		Data Asset `json:"data,omitempty"`
	}

	UpdateAssetResponse struct {
		Id  int   `json:"id,omitempty"`
		Err error `json:"error,omitempty"`
	}
)
