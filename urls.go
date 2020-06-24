package lazadago

var availablePaths map[string]string = map[string]string{
	//=======================================================
	// Shop
	//=======================================================

	"GetSeller": "seller/get",

	//=======================================================
	// Products
	//=======================================================

	"GetProducts": "products/get",

	//=======================================================
	// Orders
	//=======================================================

	"GetOrders":               "orders/get",
	"GetOrderItems":           "order/items/get",
	"SetStatusToReadyToShip":  "order/rts",
	"SetStatusToSOFDelivered": "order/sof/delivered",
}
