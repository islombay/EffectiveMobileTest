package adduser

import (
	"effectiveMobile/pkg/database"
	"effectiveMobile/pkg/httpserver/responses"
	"effectiveMobile/pkg/util/log/sl"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type AddUserBody struct {
	AgeUrl, GenderUrl, NationalityUrl string
}

// AddUser godoc
// @Summary AddUser
// @Tag root
// @Description Creates new user
// @ID add_user
// @Accept json
// @Produce json
// @Param input body RequestAddUser true "user info"
// @Success 200 {object} database.User
// @Failure 400,500 {object} responses.responsesStruct
// @Router / [post]
func AddUser(log *slog.Logger, db *database.Database, body AddUserBody) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestAddUser
		if err := c.BindJSON(&reqBody); err != nil {
			responses.ErrorResponse(c, responses.ErrorBadRequest)
			return
		}
		log.Debug("Getting age for user")
		userAge, err := GetAge(reqBody.Name, body.AgeUrl)
		if err != nil {
			log.Error("could not get age of user", sl.Err(err))
		}
		log.Debug("Getting gender for user")
		userGender, err := GetGender(reqBody.Name, body.GenderUrl)
		if err != nil {
			log.Error("could not get age of user", sl.Err(err))
		}

		log.Debug("Getting nationality for user")
		userNat, err := GetNationality(reqBody.Name, body.NationalityUrl)
		if err != nil {
			log.Error("could not get age of user", sl.Err(err))
		}

		log.Debug("saving user")
		user, err := db.AddUser(
			reqBody.Name,
			reqBody.Surname,
			reqBody.Patronymic,
			userGender,
			userNat.CountryID,
			userAge,
			userNat.Probability)
		if err != nil {
			log.Error("could not save user", sl.Err(err))
			responses.ErrorResponse(c, responses.ErrorServerError)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
