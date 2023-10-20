package getuser

import (
	"effectiveMobile/pkg/database"
	"effectiveMobile/pkg/httpserver/responses"
	"effectiveMobile/pkg/util/log/sl"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary GetUser
// @Tag root
// @Description get users filtering
// @ID get_user
// @Accept json
// @Produce json
// @Param userID query integer false "userid to find"
// @Param maxAge query integer false "maxAge for filtering"
// @Param minAge query integer false "minAge for filtering"
// @Param limit query integer false "limit of objects"
// @Param nationality query string false "users with specific nationality"
// @Param gender query string false "users with specific gender"
// @Success 200 {object} []database.User
// @Failure 400,404,500 {object} responses.responsesStruct
// @Router / [get]
func GetUser(log *slog.Logger, db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		/*
			query params
			id=
			max_age
			min_age
			limit
			nationality
			gender
		*/
		val, ok := queryParams["id"]
		if ok {
			valInt, err := strconv.Atoi(val[0])
			if err != nil {
				responses.ErrorResponse(c, responses.ErrorBadRequest)
				return
			}
			res, err := db.GetUserID(valInt)
			if err != nil {
				if err.Error() == database.ErrorNotFound {
					responses.ErrorResponse(c, responses.ErrorNotFound)
					return
				}
				log.Error("could not get user by id", sl.Err(err))
				responses.ErrorResponse(c, responses.ErrorServerError)
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}

		limit := 0
		if lmval, ok := queryParams["limit"]; ok {
			valInt, err := strconv.Atoi(lmval[0])
			if err != nil {
				responses.ErrorResponse(c, responses.ErrorBadRequest)
				return
			}
			limit = valInt
		}

		minAge := 0
		if lmval, ok := queryParams["min_age"]; ok {
			valInt, err := strconv.Atoi(lmval[0])
			if err != nil {
				responses.ErrorResponse(c, responses.ErrorBadRequest)
				return
			}
			minAge = valInt
		}

		maxAge := 0
		if lmval, ok := queryParams["max_age"]; ok {
			valInt, err := strconv.Atoi(lmval[0])
			if err != nil {
				responses.ErrorResponse(c, responses.ErrorBadRequest)
				return
			}
			maxAge = valInt
		}

		nat := ""
		if lmval, ok := queryParams["nationality"]; ok {
			nat = lmval[0]
		}

		gender := ""
		if lmval, ok := queryParams["gender"]; ok {
			gender = lmval[0]
		}

		users, err := db.GetUsers(maxAge, minAge, limit, gender, nat)
		if err != nil {
			if err.Error() == database.ErrorNotFound {
				responses.ErrorResponse(c, responses.ErrorNotFound)
				return
			}
			log.Error("could not get users by filters and pagination", sl.Err(err))
			responses.ErrorResponse(c, responses.ErrorServerError)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
