package updateuser

type RequestUpdateUser struct {
	ID              int64   `json:"id" binding:"required"`
	Name            string  `json:"name"`
	Surname         string  `json:"surname"`
	Patronymic      string  `json:"patronymic"`
	Age             int     `json:"age"`
	Gender          string  `json:"gender"`
	Nationality     string  `json:"nationality"`
	NationalityProb float64 `json:"nationality_probability"`
}
