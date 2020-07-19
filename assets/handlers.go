package assets

import (
	gokithandler "shakilakhtar/go-microservices-platform/handler"
	"shakilakhtar/go-microservices/handler"
	"shakilakhtar/go-microservices/types"
	"net/http"
)

func ListAssetsHandler(w http.ResponseWriter, r *http.Request) {
	var svc Service
	svc = NewService()
	gokithandler.EncodeJSONResponse(w, types.ListAssetsResponse{Assets: svc.ListAssets(), Err: nil})
}

func GetAssetHandler(w http.ResponseWriter, r *http.Request) {
	request, err := handler.DecodeGetAssetRequest(nil, r)
	//request, err := handler.DecodeRequest(handler.DecodeGetAssetRequest(r))
	req := request.(types.GetAssetRequest)
	var svc Service
	svc = NewService()
	a, err := svc.GetAsset(req.Id)
	gokithandler.EncodeJSONResponse(w, types.GetAssetResponse{Asset: a, Err: err})
}

func AddAssetHandler(w http.ResponseWriter, r *http.Request) {
	request, err := handler.DecodeAddAssetRequest(nil, r)
	req := request.(types.AddAssetRequest)
	asset := req.Data
	var svc Service
	svc = NewService()
	i, err := svc.Save(asset)
	gokithandler.EncodeJSONResponse(w, types.AddAssetResponse{Id: i.Id, Err: err})
}

func UpdateAssetHandler(w http.ResponseWriter, r *http.Request) {
	request, err := handler.DecodeUpdateAssetRequest(nil, r)
	req := request.(types.UpdateAssetRequest)
	item := req.Data
	var svc Service
	svc = NewService()
	i, err := svc.Update(item)
	gokithandler.EncodeJSONResponse(w, types.UpdateAssetResponse{Id: i.Id, Err: err})
}

func DeleteAssetHandler(w http.ResponseWriter, r *http.Request) {
	request, err := handler.DecodeUpdateAssetRequest(nil, r)
	req := request.(types.UpdateAssetRequest)
	item := req.Data
	var svc Service
	svc = NewService()
	i, err := svc.Update(item)
	gokithandler.EncodeJSONResponse(w, types.UpdateAssetResponse{Id: i.Id, Err: err})
}
