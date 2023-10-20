package database

import (
	"fmt"
	"strings"
)

func (db *Database) GetUserID(id int) (User, error) {
	var u User
	err := db.db.Get(&u, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (db *Database) GetUsers(maxAge, minAge, limit int, gender, nat string) ([]User, error) {
	sql := "SELECT * FROM users"
	conds := []string{}
	if maxAge != 0 && maxAge > 0 {
		conds = append(conds, fmt.Sprintf("age <= %d", maxAge))
	}
	if minAge != 0 && minAge > 0 {
		conds = append(conds, fmt.Sprintf("age >= %d", minAge))
	}
	if gender != "" {
		conds = append(conds, fmt.Sprintf("gender ='%s'", gender))
	}
	if nat != "" {
		conds = append(conds, fmt.Sprintf("nationality = '%s'", nat))
	}

	if len(conds) > 0 {
		sql += " WHERE "
		sql += strings.Join(conds, " AND ")
	}
	if limit != 0 {
		sql += fmt.Sprintf(" LIMIT %d", limit)
	}

	users := []User{}
	rows, err := db.db.Queryx(sql)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.StructScan(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
