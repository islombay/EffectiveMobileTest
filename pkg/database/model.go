package database

type User struct {
	ID              int64   `json:"id"`
	Name            string  `db:"first_name" json:"name"`
	Surname         string  `db:"last_name" json:"surname"`
	Patronymic      string  `json:"patronymic,omitempty"`
	Age             int     `json:"age,omitempty"`
	Gender          string  `json:"gender,omitempty"`
	Nationality     string  `json:"nationality,omitempty"`
	NationalityProb float64 `db:"nationality_probability" json:"nationality_probability,omitempty"`
}
