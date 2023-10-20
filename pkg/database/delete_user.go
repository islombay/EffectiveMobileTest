package database

func (db *Database) DeleteUser(id int) error {
	sql := "DELETE FROM users WHERE id = $1"
	if _, err := db.db.Exec(sql, id); err != nil {
		return err
	}
	return nil
}
