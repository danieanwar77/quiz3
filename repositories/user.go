package repository

import (
	"database/sql"
	"quiz3/structs"
	"time"
)

func GetAllUser(db *sql.DB) (results []structs.Users, err error) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var users structs.Users

		err = rows.Scan(&users.ID, &users.UserName, &users.Password, &users.CreatedAt, &users.CreatedBy, &users.ModifiedAt, &users.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, users)
	}
	return
}

func GetUser(db *sql.DB, users structs.Users) (results []structs.Users, err error) {
	sql := "SELECT * FROM users WHERE id = $1"

	rows, err := db.Query(sql, users.ID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var users structs.Users

		err = rows.Scan(&users.ID, &users.UserName, &users.Password, &users.CreatedAt, &users.CreatedBy, &users.ModifiedAt, &users.ModifiedBy)
		if err != nil {
			return
		}

		results = append(results, users)
	}
	return

}

func InsertUser(db *sql.DB, users structs.Users) (err error) {

	sql := "INSERT INTO users(id, username, password, created_at, created_by) VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sql, &users.ID, &users.UserName, &users.Password, time.Now().Format("2006.01.02 15:04:05"), &users.CreatedBy)

	return errs.Err()
}

func UpdateUser(db *sql.DB, users structs.Users) (err error) {
	sql := "UPDATE users SET username = $1, password = $2, modified_at = $4, modified_by = $5 WHERE id = $6"

	errs := db.QueryRow(sql, users.UserName, users.Password, time.Now().Format("2006.01.02 15:04:05"), time.Now().Format("2006.01.02 15:04:05"), users.ID)

	return errs.Err()
}

func DeleteUser(db *sql.DB, users structs.Users) (err error) {
	sql := "DELETE FROM users WHERE id = $1"

	errs := db.QueryRow(sql, users.ID)
	return errs.Err()
}
