package lazadago

// GetShopInfoResponse response
type GetShopInfoResponse struct {
	NameCompany string `json:"name_company,omitempty"`
	SellerID    int64  `json:"seller_id,omitempty"`
	Name        string `json:"name,omitempty"`
	ShortCode   string `json:"short_code,omitempty"`
	LogoURL     string `json:"logo_url,omitempty"`
	Email       string `json:"email,omitempty"`
	CB          bool   `json:"cb,omitempty"`
	Location    string `json:"location,omitempty"`
}
