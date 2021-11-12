package productcategory

import "github.com/pandudpn/shopping-cart/app/adapter/dbc"

type ProductCategoryRepository struct {
	DB dbc.SqlDbc
}
