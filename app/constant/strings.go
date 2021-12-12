package constant

const (
	PSQL  = "psql"
	REDIS = "redis"

	TX       = "tx"
	ENABLETX = true

	// Table database
	USER                = "user"
	USERADDRESS         = "user_address"
	PRODUCT             = "product"
	PRODUCTCATEGORY     = "product_category"
	PRODUCTIMAGE        = "product_image"
	MEDIAFILE           = "media_file"
	CART                = "cart"
	CART_PRODUCT        = "cart_product"
	COURIER             = "courier"
	STOCK               = "stock"
	PAYMENTMETHOD       = "payment_method"
	ORDER               = "order"
	ORDERPRODUCT        = "order_product"
	ORDERPAYMENT        = "order_payment"
	ORDERDELIVERY       = "order_delivery"
	ORDERDELIVERYSTATUS = "order_delivery_status"

	// Usecase
	CHECKOUT = "checkout"

	// Key Payment third party
	XENDITEWALLET        = "xenditewallet"
	XENDITVIRTUALACCOUNT = "xenditvirtualaccount"
)
