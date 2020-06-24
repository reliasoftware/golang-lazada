package lazadago

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	// Version lazada api latest updated
	Version = "lazop-sdk-go-20200515"

	// APIGatewaySG endpoint
	APIGatewaySG = "https://api.lazada.sg/rest"
	// APIGatewayMY endpoint
	APIGatewayMY = "https://api.lazada.com.my/rest"
	// APIGatewayVN endpoint
	APIGatewayVN = "https://api.lazada.vn/rest"
	// APIGatewayTH endpoint
	APIGatewayTH = "https://api.lazada.co.th/rest"
	// APIGatewayPH endpoint
	APIGatewayPH = "https://api.lazada.com.ph/rest"
	// APIGatewayID endpoint
	APIGatewayID = "https://api.lazada.co.id/rest"
)

// ClientOptions params
type ClientOptions struct {
	ServerURL string
	APIKey    string
	APISecret string
}

// LazadaClient represents a client to Lazada
type LazadaClient struct {
	ServerURL string
	APIKey    string
	APISecret string

	Method     string
	SysParams  map[string]string
	APIParams  map[string]string
	FileParams map[string][]byte
}

// NewClient init
func NewClient(opts *ClientOptions) Client {
	return &LazadaClient{
		ServerURL: opts.ServerURL,
		APIKey:    opts.APIKey,
		APISecret: opts.APISecret,
		SysParams: map[string]string{
			"app_key":     opts.APIKey,
			"sign_method": "sha256",
			"timestamp":   fmt.Sprintf("%d000", time.Now().Unix()),
			"partner_id":  Version,
		},
		APIParams:  map[string]string{},
		FileParams: map[string][]byte{},
	}
}

// Debug setter
func (lc *LazadaClient) Debug(enableDebug bool) *LazadaClient {
	if enableDebug {
		lc.SysParams["debug"] = "true"
	} else {
		lc.SysParams["debug"] = "false"
	}
	return lc
}

// AccessToken setter
func (lc *LazadaClient) AccessToken(accessToken string) *LazadaClient {
	lc.SysParams["access_token"] = accessToken
	return lc
}

// AddAPIParam setter
func (lc *LazadaClient) AddAPIParam(key string, val string) *LazadaClient {
	lc.APIParams[key] = val
	return lc
}

// AddFileParam setter
func (lc *LazadaClient) AddFileParam(key string, val []byte) *LazadaClient {
	lc.FileParams[key] = val
	return lc
}

// Create sign from system params and api params
func (lc *LazadaClient) sign(url string) string {
	keys := []string{}
	union := map[string]string{}
	for key, val := range lc.SysParams {
		union[key] = val
		keys = append(keys, key)
	}
	for key, val := range lc.APIParams {
		union[key] = val
		keys = append(keys, key)
	}

	// sort sys params and api params by key
	sort.Strings(keys)

	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("%s", url))
	for _, key := range keys {
		message.WriteString(fmt.Sprintf("%s%s", key, union[key]))
	}

	hash := hmac.New(sha256.New, []byte(lc.APISecret))
	hash.Write(message.Bytes())
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}

// Response success
type Response struct {
	Code      string          `json:"code"`
	Type      string          `json:"type"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
}

// ResponseError defines a error response
type ResponseError struct {
	Code      string `json:"code"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func (lc *LazadaClient) getPath(apiName string) string {
	return fmt.Sprintf("%s", availablePaths[apiName])
}

// Execute sends the request though http.request and collect the response
func (lc *LazadaClient) Execute(apiName string, apiMethod string, apiParams interface{}) (*Response, error) {
	var req *http.Request
	var err error
	var contentType string

	// bodyParams, err := json.Marshal(apiParams)
	// if err != nil {
	// 	return nil, err
	// }

	// add query params
	values := url.Values{}
	for key, val := range lc.SysParams {
		values.Add(key, val)
	}

	// POST handle
	body := &bytes.Buffer{}
	if apiMethod == http.MethodPost {
		writer := multipart.NewWriter(body)
		contentType = writer.FormDataContentType()
		if len(lc.FileParams) > 0 {
			// add formfile to handle file upload
			for key, val := range lc.FileParams {
				part, err := writer.CreateFormFile("image", key)
				if err != nil {
					return nil, err
				}
				_, err = part.Write(val)
				if err != nil {
					return nil, err
				}
			}
		}

		for key, val := range lc.APIParams {
			_ = writer.WriteField(key, val)
		}

		if err = writer.Close(); err != nil {
			return nil, err
		}
	}

	// GET handle
	if apiMethod == http.MethodGet {
		for key, val := range lc.APIParams {
			values.Add(key, val)
		}
	}

	apiPath := lc.getPath(apiName)
	values.Add("sign", lc.sign(apiPath))
	fullURL := fmt.Sprintf("%s%s?%s", lc.ServerURL, apiPath, values.Encode())
	req, err = http.NewRequest(apiMethod, fullURL, body)

	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	resp := &Response{}
	err = json.Unmarshal(respBody, resp)

	// lc.APIParams = nil
	// lc.FileParams = nil

	return resp, err
}

// Client interface api
type Client interface {
	//=======================================================
	// Shop
	//=======================================================

	// GetSeller Use this call to get information of shop
	GetSeller() (*GetShopInfoResponse, error)
}

//=======================================================
// Shop
//=======================================================

// GetSeller Use this call to get information of shop
func (lc *LazadaClient) GetSeller() (resp *GetShopInfoResponse, err error) {
	b, err := lc.Execute("GetShopInfo", "GET", nil)

	if err != nil {
		return
	}
	err = json.Unmarshal(b.Data, &resp)

	if err != nil {
		return
	}
	return
}
