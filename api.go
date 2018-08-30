package lazada

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
	Version = "lazop-sdk-go-20180830"

	ApiGatewaySG = "https://api.lazada.sg/rest"
	ApiGatewayMY = "https://api.lazada.com.my/rest"
	ApiGatewayVN = "https://api.lazada.vn/rest"
	ApiGatewayTH = "https://api.lazada.co.th/rest"
	ApiGatewayPH = "https://api.lazada.com.ph/rest"
	ApiGatewayID = "https://api.lazada.co.id/rest"
)

type Request struct {
	serverUrl string
	apiKey    string
	apiSecret string

	apiPath    string
	method     string
	sysParams  map[string]string
	apiParams  map[string]string
	fileParams map[string][]byte
}

func NewRequest(apiKey, secret, serverUrl string) *Request {
	return &Request{
		serverUrl: serverUrl,
		apiKey:    apiKey,
		apiSecret: secret,
		sysParams: map[string]string{
			"app_key":     apiKey,
			"sign_method": "sha256",
			"timestamp":   fmt.Sprintf("%d000", time.Now().Unix()),
			"partner_id":  Version,
		},
		apiParams:  map[string]string{},
		fileParams: map[string][]byte{},
	}
}

func (r *Request) ApiPath(path string) *Request {
	r.apiPath = path
	return r
}

func (r *Request) Method(method string) *Request {
	r.method = method
	return r
}

func (r *Request) Debug(enableDebug bool) *Request {
	if enableDebug {
		r.sysParams["debug"] = "true"
	} else {
		r.sysParams["debug"] = "false"
	}
	return r
}

func (r *Request) AccessToken(accessToken string) *Request {
	r.sysParams["access_token"] = accessToken
	return r
}

func (r *Request) AddApiParam(key string, val string) *Request {
	r.apiParams[key] = val
	return r
}

func (r *Request) AddFileParam(key string, val []byte) *Request {
	r.fileParams[key] = val
	return r
}

func (r *Request) sign() string {
	keys := []string{}
	union := map[string]string{}
	for key, val := range r.sysParams {
		union[key] = val
		keys = append(keys, key)
	}
	for key, val := range r.apiParams {
		union[key] = val
		keys = append(keys, key)
	}

	// sort sys params and api params by key
	sort.Strings(keys)

	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("%s", r.apiPath))
	for _, key := range keys {
		message.WriteString(fmt.Sprintf("%s%s", key, union[key]))
	}

	hash := hmac.New(sha256.New, []byte(r.apiSecret))
	hash.Write(message.Bytes())
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}

type Response struct {
	Code      string          `json:"code"`
	Type      string          `json:"type"`
	Message   string          `json:"message"`
	RequestId string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
}

// Do sends the request though http.request and collect the response
func (r *Request) Do() (*Response, error) {
	var req *http.Request
	var err error
	var contentType string

	body := &bytes.Buffer{}
	if r.method == http.MethodPost {
		writer := multipart.NewWriter(body)
		contentType = writer.FormDataContentType()
		if len(r.fileParams) > 0 {
			// add formfile to handle file upload
			for key, val := range r.fileParams {
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
		if err = writer.Close(); err != nil {
			return nil, err
		}
	}
	// add query params
	values := url.Values{}
	for key, val := range r.sysParams {
		values.Add(key, val)
	}
	for key, val := range r.apiParams {
		values.Add(key, val)
	}
	values.Add("sign", r.sign())
	fullUrl := fmt.Sprintf("%s%s?%s", r.serverUrl, r.apiPath, values.Encode())
	req, err = http.NewRequest(r.method, fullUrl, body)
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
	return resp, err
}
