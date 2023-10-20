package updateuser

import (
	"effectiveMobile/pkg/database"
	"effectiveMobile/pkg/httpserver/responses"
	"effectiveMobile/pkg/util/log/sl"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func UpdateUser(log *slog.Logger, db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestUpdateUser
		if err := c.BindJSON(&reqBody); err != nil {
			responses.ErrorResponse(c, responses.ErrorBadRequest)
			return
		}
		updatedUser, err := db.UpdateUser(database.User{
			ID:              reqBody.ID,
			Name:            reqBody.Name,
			Surname:         reqBody.Surname,
			Patronymic:      reqBody.Patronymic,
			Age:             reqBody.Age,
			Gender:          reqBody.Gender,
			Nationality:     reqBody.Nationality,
			NationalityProb: reqBody.NationalityProb,
		})
		if err != nil {
			if err.Error() == database.ErrorNotFound {
				responses.ErrorResponse(c, responses.ErrorNotFound)
				return
			}
			log.Error("could not update user in db", sl.Err(err))
			responses.ErrorResponse(c, responses.ErrorServerError)
			return
		}
		c.JSON(http.StatusOK, updatedUser)
	}
}
