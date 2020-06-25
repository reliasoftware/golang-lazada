# go-lazada

```golang
// Using
var clientOptions = ClientOptions{
  APIKey:    apiKey,
  APISecret: apiSecret,
  Region: "VN",
}

lc := NewClient(&clientOptions)
resp, _ := lc.GetSeller()
fmt.Println(resp.Name)

```
