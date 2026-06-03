package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetProdpoID(userID int, search string, limit int, lastID int, lastPrice float64, sort string, categoryID int) ([]models.Product, int) {
	var imgURL sql.NullString
	var desc sql.NullString
	var totalCount int

	whereClause := `
		WHERE o.seller_id = $1 
		  AND ($2 = '' OR p.name %% $2)
		  AND ($3 = 0 OR p.category_id = $3)
	`

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM products p
		INNER JOIN product_offers o ON p.id = o.product_id
		%s
	`, whereClause)

	err := db.QueryRow(countQuery, userID, search, categoryID).Scan(&totalCount)
	if err != nil {
		panic(err)
	}

	if totalCount == 0 {
		return []models.Product{}, 0
	}

	var orderBy string

	switch sort {
	case "price_asc":
		orderBy = "o.price ASC, p.id ASC"
		if lastID > 0 {

			whereClause += fmt.Sprintf(" AND (o.price, p.id) > (%f, %d)", lastPrice, lastID)
		}
	case "price_desc":
		orderBy = "o.price DESC, p.id DESC"
		if lastID > 0 {
			whereClause += fmt.Sprintf(" AND (o.price, p.id) < (%f, %d)", lastPrice, lastID)
		}
	case "id_asc":

		orderBy = "p.id ASC"
		if lastID > 0 {
			whereClause += fmt.Sprintf(" AND p.id > %d", lastID)
		}
	case "new":
		orderBy = "p.id DESC"
		if lastID > 0 {
			whereClause += fmt.Sprintf(" AND p.id < %d", lastID)
		}
	default:

		if search != "" {
			orderBy = "similarity(p.name, $2) DESC, p.id DESC"
		} else {
			orderBy = "p.id DESC"
		}
		if lastID > 0 {
			whereClause += fmt.Sprintf(" AND p.id < %d", lastID)
		}
	}

	dataQuery := fmt.Sprintf(`
		SELECT p.id, p.name, p.description, o.price, o.count, p.img_url, p.category_id
		FROM products p
		INNER JOIN product_offers o ON p.id = o.product_id
		%s
		ORDER BY %s
		LIMIT $4
	`, whereClause, orderBy)

	rows, err := db.Query(dataQuery, userID, search, categoryID, limit)
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
			product.ImgURL = imgURL.String
		}
		products = append(products, product)
	}

	return products, totalCount
}
