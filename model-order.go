package lazadago

// OrderDetailAddressBilling data
type OrderDetailAddressBilling struct {
	Country       string `json:"country"`
	Address3      string `json:"address3"`
	Address2      string `json:"address2"`
	City          string `json:"city"`
	Address1      string `json:"address1"`
	Phone2        string `json:"phone2"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	CustomerEmail string `json:"customer_email"`
	PostCode      string `json:"post_code"`
	Address5      string `json:"address5"`
	Address4      string `json:"address4"`
	FirstName     string `json:"first_name"`
}

// OrderDetailAddressShipping data
type OrderDetailAddressShipping struct {
	Country       string `json:"country"`
	Address3      string `json:"address3"`
	Address2      string `json:"address2"`
	City          string `json:"city"`
	Address1      string `json:"address1"`
	Phone2        string `json:"phone2"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	CustomerEmail string `json:"customer_email"`
	PostCode      string `json:"post_code"`
	Address5      string `json:"address5"`
	Address4      string `json:"address4"`
	FirstName     string `json:"first_name"`
}

// OrderDetail data
type OrderDetail struct {
	VoucherPlatform            float64                    `json:"voucher_platform"`
	Voucher                    float64                    `json:"voucher"`
	OrderNumber                int64                      `json:"order_number"`
	VoucherSeller              float64                    `json:"voucher_seller"`
	CreatedAt                  string                     `json:"created_at"`
	VoucherCode                string                     `json:"voucher_code"`
	GiftOption                 bool                       `json:"gift_option"`
	CustomerLastName           string                     `json:"customer_last_name"`
	UpdatedAt                  string                     `json:"updated_at"`
	PromisedShippingTimes      string                     `json:"promised_shipping_times"`
	Price                      string                     `json:"price"`
	NationalRegistrationNumber string                     `json:"national_registration_number"`
	PaymentMethod              string                     `json:"payment_method"`
	CustomerFirstName          string                     `json:"customer_first_name"`
	ShippingFee                float64                    `json:"shipping_fee"`
	ItemsCount                 int                        `json:"items_count"`
	DeliveryInfo               string                     `json:"delivery_info"`
	Statuses                   []string                   `json:"statuses"`
	AddressBilling             OrderDetailAddressBilling  `json:"address_billing"`
	ExtraAttributes            string                     `json:"extra_attributes"`
	OrderID                    int64                      `json:"order_id"`
	GiftMessage                string                     `json:"gift_message"`
	Remarks                    string                     `json:"remarks"`
	AddressShipping            OrderDetailAddressShipping `json:"address_shipping"`
	TaxCode                    string                     `json:"tax_code,omitempty"`
}

// GetOrdersResponse api response
type GetOrdersResponse struct {
	Count  int           `json:"count"`
	Orders []OrderDetail `json:"orders"`
}

// OrderItem data
type OrderItem struct {
	PaidPrice             float64 `json:"paid_price"`
	ProductMainImage      string  `json:"product_main_image"`
	TaxAmount             float64 `json:"tax_amount"`
	VoucherPlatform       float64 `json:"voucher_platform"`
	Reason                string  `json:"reason"`
	ProductDetailURL      string  `json:"product_detail_url"`
	PromisedShippingTime  string  `json:"promised_shipping_time"`
	PurchaseOrderID       string  `json:"purchase_order_id"`
	VoucherSeller         float64 `json:"voucher_seller"`
	ShippingType          string  `json:"shipping_type"`
	CreatedAt             string  `json:"created_at"`
	VoucherCode           string  `json:"voucher_code"`
	PackageID             string  `json:"package_id"`
	Variation             string  `json:"variation"`
	UpdatedAt             string  `json:"updated_at"`
	PurchaseOrderNumber   string  `json:"purchase_order_number"`
	Currency              string  `json:"currency"`
	ShippingProviderType  string  `json:"shipping_provider_type"`
	Sku                   string  `json:"sku"`
	InvoiceNumber         string  `json:"invoice_number"`
	CancelReturnInitiator string  `json:"cancel_return_initiator"`
	ShopSku               string  `json:"shop_sku"`
	IsDigital             int     `json:"is_digital"`
	ItemPrice             float64 `json:"item_price"`
	ShippingServiceCost   int     `json:"shipping_service_cost"`
	TrackingCodePre       string  `json:"tracking_code_pre"`
	TrackingCode          string  `json:"tracking_code"`
	ShippingAmount        float64 `json:"shipping_amount"`
	OrderItemID           int64   `json:"order_item_id"`
	ReasonDetail          string  `json:"reason_detail"`
	ShopID                string  `json:"shop_id"`
	ReturnStatus          string  `json:"return_status"`
	Name                  string  `json:"name"`
	ShipmentProvider      string  `json:"shipment_provider"`
	VoucherAmount         float64 `json:"voucher_amount"`
	DigitalDeliveryInfo   string  `json:"digital_delivery_info"`
	ExtraAttributes       string  `json:"extra_attributes"`
	OrderID               int64   `json:"order_id"`
	Status                string  `json:"status"`
}

// GetOrderItemsResponse api response
type GetOrderItemsResponse struct {
	Items []OrderItem `json:"items"`
}
