package repository

import (
	"database/sql"
	"fmt"
	"quiz3/structs"
	"time"
)

func GetAllBook(db *sql.DB) (results []structs.Book, err error) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryId, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, book)
	}
	return
}

func GetBook(db *sql.DB, book structs.Book) (results []structs.Book, err error) {
	sql := "SELECT * FROM books WHERE id = $1"

	rows, err := db.Query(sql, book.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryId, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, book)
	}
	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	if book.ReleaseYear >= 1980 && book.ReleaseYear <= 2024 {

		sql := "INSERT INTO books(id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

		var thick string
		total := &book.TotalPage

		if *total > 100 {
			thick = "tebal"
		} else {
			thick = "tipis"
		}

		errs := db.QueryRow(sql, &book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, thick, &book.CategoryId, time.Now().Format("2006.01.02 15:04:05"), &book.CreatedBy)

		return errs.Err()
	} else {
		errs := fmt.Errorf("Tahun rilis lebih tua dari 1980 atau berada di masa depan")
		return errs
	}
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	if book.ReleaseYear >= 1980 && book.ReleaseYear <= 2024 {

		sql := "UPDATE books SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, modified_at = $9, modified_by = $10 WHERE id = $11"

		var thick string
		total := &book.TotalPage

		if *total > 100 {
			thick = "tebal"
		} else {
			thick = "tipis"
		}

		upd, errs := db.Exec(sql, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, thick, &book.CategoryId, time.Now().Format("2006.01.02 15:04:05"), time.Now().Format("2006.01.02 15:04:05"), book.ID)

		count, err := upd.RowsAffected()

		if err == nil {
			if count == 0 {
				errs = fmt.Errorf("Book not found!")
			}
		} else {
			errs = err
		}

		return errs
	} else {
		errs := fmt.Errorf("Tahun rilis lebih tua dari 1980 atau berada di masa depan")
		return errs
	}
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM books WHERE id = $1"

	del, errs := db.Exec(sql, book.ID)

	count, err := del.RowsAffected()
	if err == nil {
		if count == 0 {
			errs = fmt.Errorf("Book not found!")
		}
	} else {
		errs = err
	}

	return errs
}
