package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetAllProd(search string, limit int, offset int, sort string, categoryID int) ([]models.Product, int) {
	var imgURL sql.NullString
	var desc sql.NullString
	search = "%" + search + "%"

	var totalCount int
	countQuery := `
		SELECT COUNT(*) 
		FROM products 
		WHERE ($1 = '%%' OR name ILIKE $1 OR description ILIKE $1)
		  AND ($2 = 0 OR category_id = $2)
	`
	err := db.QueryRow(countQuery, search, categoryID).Scan(&totalCount)
	if err != nil {
		panic(err)
	}

	orderBy := "id DESC"
	switch sort {
	case "price_asc":
		orderBy = "price ASC"
	case "price_desc":
		orderBy = "price DESC"
	}

	dataQuery := fmt.Sprintf(`
		SELECT id, name, description, price, count, seller_id, img_url, category_id
		FROM products
		WHERE ($1 = '%%' OR name ILIKE $1 OR description ILIKE $1)
		  AND ($2 = 0 OR category_id = $2)
		ORDER BY %s
		LIMIT $3 OFFSET $4
	`, orderBy)

	rows, err := db.Query(dataQuery, search, categoryID, limit, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&desc,
			&product.Price,
			&product.Count,
			&product.SellerID,
			&imgURL,
			&product.Category_id,
		)
		if err != nil {
			panic(err)
		}

		if desc.Valid {
			product.Description = desc.String
		}
		if imgURL.Valid {
			product.ImageURL = imgURL.String
		}

		products = append(products, product)
	}

	return products, totalCount
}
