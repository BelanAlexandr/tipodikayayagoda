package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetProdpoID(userID int, search string, limit int, offset int, sort string) ([]models.Product, int) {
	var imgURL sql.NullString
	var desc sql.NullString
	search = "%" + search + "%"

	var totalCount int
	countQuery := `
		SELECT COUNT(*) 
		FROM products 
		WHERE seller_id = $1 AND ($2 = '%%' OR name ILIKE $2 OR description ILIKE $2)
	`
	err := db.QueryRow(countQuery, userID, search).Scan(&totalCount)
	if err != nil {
		panic(err)
	}

	// Определяем сортировку
	orderBy := "id DESC"
	switch sort {
	case "price_asc":
		orderBy = "price ASC"
	case "price_desc":
		orderBy = "price DESC"
	}

	dataQuery := fmt.Sprintf(`
		SELECT id, name, description, price, count, seller_id, img_url
		FROM products
		WHERE seller_id = $1 AND ($2 = '%%' OR name ILIKE $2 OR description ILIKE $2)
		ORDER BY %s
		LIMIT $3 OFFSET $4
	`, orderBy)

	rows, err := db.Query(dataQuery, userID, search, limit, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &desc, &product.Price, &product.Count, &product.SellerID, &imgURL)
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
