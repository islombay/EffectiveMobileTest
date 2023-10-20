package database

func (db *Database) UpdateUser(u User) (User, error) {
	oldU, err := db.findUserByID(u.ID)
	if err != nil {
		return User{}, err
	}
	sql := `UPDATE users SET
			first_name = $1, last_name=$2, patronymic=$3, age=$4,
			gender=$5, nationality=$6, nationality_probability=$7
			WHERE id = $8`
	if u.Name == "" {
		u.Name = oldU.Name
	}
	if u.Surname == "" {
		u.Surname = oldU.Surname
	}
	if u.Patronymic == "" {
		u.Patronymic = oldU.Patronymic
	}
	if u.Age == 0 {
		u.Age = oldU.Age
	}
	if u.Gender == "" {
		u.Gender = oldU.Gender
	}
	if u.Nationality == "" {
		u.Nationality = oldU.Nationality
	}
	if u.NationalityProb == 0 {
		u.NationalityProb = oldU.NationalityProb
	}
	_, err = db.db.Exec(sql, u.Name, u.Surname, u.Patronymic, u.Age, u.Gender, u.Nationality, u.NationalityProb, u.ID)
	if err != nil {
		return User{}, err
	}
	return db.findUserByID(u.ID)
}

func (db *Database) findUserByID(id int64) (User, error) {
	sql := `SELECT * FROM users WHERE id = $1`
	var user User
	if err := db.db.Get(&user, sql, id); err != nil {
		return user, err
	}
	return user, nil
}
