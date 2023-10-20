package adduser

type RequestAddUser struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

type ResponseAddUser struct {
}

type ResponseAgeService struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ResponseGenderService struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type ResponseNationalityService struct {
	Name    string                              `json:"name"`
	Country []ResponseNationalityCountryService `json:"country"`
}

type ResponseNationalityCountryService struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
