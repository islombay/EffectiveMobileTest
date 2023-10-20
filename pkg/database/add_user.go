package database

func (db *Database) AddUser(name, surname, pat, gender, nat string, age int, natProb float64) (User, error) {
	sql := `INSERT INTO users (
                   first_name,
                   last_name,
                   patronymic,
                   age,
                   gender,
                   nationality,
                   nationality_probability
	) VALUES($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.db.Exec(sql, name, surname, pat, age, gender, nat, natProb)
	if err != nil {
		return User{}, err
	}

	return db.findUser(name, surname, pat, natProb)
}

func (db *Database) findUser(fn, ln, pat string, natProb float64) (User, error) {
	sql := `SELECT * FROM users WHERE first_name = $1 AND last_name = $2 AND patronymic = $3 AND nationality_probability = $4`
	var user User
	if err := db.db.Get(&user, sql, fn, ln, pat, natProb); err != nil {
		return user, err
	}
	return user, nil
}
