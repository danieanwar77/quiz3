package repository

import (
	"database/sql"
	"quiz3/structs"
)

func AuthUser(db *sql.DB, users structs.Users) (results structs.Users, err error) {
	sql := "SELECT * FROM users WHERE username = $1 AND password = $2"

	rows, err := db.Query(sql, users.UserName, users.Password)
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

		results = users
	}
	return

}
