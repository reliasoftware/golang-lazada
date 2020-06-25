# go-lazada

```golang
// Using
var clientOptions = ClientOptions{
  APIKey:    apiKey,
  APISecret: apiSecret,
  Region: "VN",
}

lc := NewClient(&clientOptions)
lc.SetAccessToken("example token get from auth")
lc.ChangeRgion("MY") // Change region

resp, _ := lc.GetSeller()
fmt.Println(resp.Name)

```
