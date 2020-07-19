package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"shakilakhtar/go-microservices/types"
	"net/http"
	"strconv"
)

type decodeRequestFunc func(_ context.Context, r *http.Request) (interface{}, error)

func DecodeListAssetsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return types.ListAssetsRequest{}, nil
}
func DecodeGetAssetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	value := vars["assetId"]
	assetId, _ := strconv.Atoi(value)
	return types.GetAssetRequest{Id: assetId}, nil
}

func DecodeAddAssetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var payload types.AddAssetRequest
	json.NewDecoder(r.Body).Decode(&payload)
	defer r.Body.Close()
	return payload, nil
}

func DecodeUpdateAssetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var payload types.UpdateAssetRequest
	json.NewDecoder(r.Body).Decode(&payload)
	defer r.Body.Close()
	return payload, nil
}

//func DecodeRequest(fn decodeRequestFunc) (interface{}, error){
//	return fn
//}