package lazadago

// ProductDetailSkus data
type ProductDetailSkus struct {
	Status          string   `json:"Status"`
	Quantity        int      `json:"quantity"`
	ProductWeight   float64  `json:"product_weight"`
	Images          []string `json:"Images"`
	SellerSku       string   `json:"SellerSku"`
	ShopSku         string   `json:"ShopSku"`
	URL             string   `json:"Url"`
	PackageWidth    string   `json:"package_width"`
	SpecialToTime   string   `json:"special_to_time"`
	ColorFamily     string   `json:"color_family"`
	SpecialFromTime string   `json:"special_from_time"`
	PackageHeight   string   `json:"package_height"`
	SpecialPrice    float64  `json:"special_price"`
	Price           float64  `json:"price"`
	PackageLength   string   `json:"package_length"`
	SpecialFromDate string   `json:"special_from_date"`
	PackageWeight   string   `json:"package_weight"`
	Available       int      `json:"Available"`
	SkuID           int      `json:"SkuId"`
	SpecialToDate   string   `json:"special_to_date"`
}

// ProductDetailAttributes data
type ProductDetailAttributes struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Brand            string `json:"brand"`
	WarrantyType     string `json:"warranty_type"`
}

// ProductDetail data
type ProductDetail struct {
	Skus            ProductDetailSkus       `json:"skus"`
	ItemID          int                     `json:"item_id"`
	PrimaryCategory int                     `json:"primary_category"`
	Attributes      ProductDetailAttributes `json:"attributes"`
}

// GetProductsResponse response
type GetProductsResponse struct {
	TotalProducts int             `json:"total_products"`
	Products      []ProductDetail `json:"products"`
}
