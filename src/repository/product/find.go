package product

import (
	"strings"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_PRODUCT_SELECT = "select p.id, p.name, p.slug, description, price, discounted_price, " +
		"qty, p.enabled, p.created_at, c.id, c.name, c.slug, s.id, coalesce(s.quantity_hold, 0) " +
		"from public.product as p inner join public.product_category as c on c.id = p.category_id " +
		"left join stock s on p.id = s.product_id "
	QUERY_PRODUCT         = QUERY_PRODUCT_SELECT + "where p.enabled = true order by p.created_at desc"
	QUERY_PRODUCT_BY_ID   = QUERY_PRODUCT_SELECT + "where p.id = $1"
	QUERY_PRODUCT_BY_SLUG = QUERY_PRODUCT_SELECT + "where p.slug = $1"
	QUERY_PRODUCT_BY_NAME = QUERY_PRODUCT_SELECT + "where p.enabled = true and lower(p.name) like "
)

func (pr *ProductRepository) FindAllProducts() ([]*model.Product, error) {
	var products = make([]*model.Product, 0)

	rows, err := pr.DB.Query(QUERY_PRODUCT)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		product, err := rowsToProduct(rows)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) FindProductsByName(name string) ([]*model.Product, error) {
	var products = make([]*model.Product, 0)
	name = strings.ToLower(name)

	rows, err := pr.DB.Query(QUERY_PRODUCT_BY_NAME + "'" + name + "%'")
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		product, err := rowsToProduct(rows)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) FindProductById(id int) (*model.Product, error) {
	row := pr.DB.QueryRow(QUERY_PRODUCT_BY_ID, id)

	product, err := rowToProduct(row)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepository) FindProductBySlug(slug string) (*model.Product, error) {
	row := pr.DB.QueryRow(QUERY_PRODUCT_BY_SLUG, slug)

	product, err := rowToProduct(row)
	if err != nil {
		return nil, err
	}

	return product, nil
}
