package lazada

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	// please to go https://open.lazada.com/ to register one
	apiKey      = "<app-id>"
	apiSecret   = "<app-secret>"
	accessToken = "<access-token>"
)

func TestPhotoUpload(t *testing.T) {
	filename := "/Users/hsinhoyeh/Desktop/1.png"
	blob, _ := ioutil.ReadFile(filename)
	r := NewRequest(apiKey, apiSecret, ApiGatewayID).
		Method(http.MethodPost).
		ApiPath("/image/upload").
		AddFileParam("image", blob).
		AccessToken(accessToken)
	resp, err := r.Do()
	fmt.Printf("resp:%v, err:%v\n", string(resp.Data), err)
}

func TestGetProduct(t *testing.T) {
	r := NewRequest(apiKey, apiSecret, ApiGatewayID).
		Method(http.MethodGet).
		ApiPath("/products/get").
		AddApiParam("filter", "live").
		AddApiParam("limit", "10").
		AddApiParam("offset", "0").
		AccessToken(accessToken)
	resp, err := r.Do()
	fmt.Printf("resp:%v, err:%v\n", string(resp.Data), err)
}
