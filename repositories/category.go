package repository

import (
	"database/sql"
	"fmt"
	"quiz3/structs"
	"time"
)

func GetAllCategory(db *sql.DB) (results []structs.Category, err error) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	count := make([]int, 0)
	for rows.Next() {
		var category structs.Category

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
		if err != nil {
			return
		}

		count = append(count, category.ID)

		results = append(results, category)
	}

	if len(count) == 0 {
		err = fmt.Errorf("Category not found!")
	}

	return
}

func GetCategory(db *sql.DB, category structs.Category) (results []structs.Category, err error) {
	sql := "SELECT * FROM category where id = $1"

	rows, err := db.Query(sql, category.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	count := make([]int, 0)
	for rows.Next() {
		var category structs.Category

		fmt.Println("rows :", rows)

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
		if err != nil {
			return
		}

		count = append(count, category.ID)

		results = append(results, category)
	}

	if len(count) == 0 {
		err = fmt.Errorf("Category not found!")
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category(id, name, created_at, created_by) VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql, &category.ID, &category.Name, time.Now().Format("2006.01.02 15:04:05"), &category.CreatedBy)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1, modified_at = $2, modified_by = $3 WHERE id = $4"

	upd, errs := db.Exec(sql, category.Name, time.Now().Format("2006.01.02 15:04:05"), time.Now().Format("2006.01.02 15:04:05"), category.ID)

	count, err := upd.RowsAffected()

	if err == nil {
		if count == 0 {
			errs = fmt.Errorf("Category not found!")
		}
	} else {
		errs = err
	}

	return errs
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	del, errs := db.Exec(sql, category.ID)

	count, err := del.RowsAffected()
	if err == nil {
		if count == 0 {
			errs = fmt.Errorf("Category not found!")
		}
	} else {
		errs = err
	}

	return errs
}
