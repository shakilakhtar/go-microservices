package assets

import "net/http"

const (
	ContextPath    = "/goseed/v1"
	ListAssetsURI  = ContextPath + "/assets"
	GetAssetURI    = ContextPath + "/assets/{assetId}"
	AddAssetURI    = ContextPath + "/assets"
	UpdateAssetURI = ContextPath + "/assets"
	DeleteAssetURI = ContextPath + "/assets"

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"GetAsset",
		GET,
		GetAssetURI,
		GetAssetHandler,
	},
	Route{
		"ListAssets",
		GET,
		ListAssetsURI,
		ListAssetsHandler,
	},

	Route{
		"AddAsset",
		POST,
		AddAssetURI,
		AddAssetHandler,
	},
	Route{
		"UpdateAsset",
		PUT,
		UpdateAssetURI,
		UpdateAssetHandler,
	},
	Route{
		"DeleteAsset",
		DELETE,
		DeleteAssetURI,
		DeleteAssetHandler,
	},
}
