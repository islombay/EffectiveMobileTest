package deleteuser

import (
	"effectiveMobile/pkg/database"
	"effectiveMobile/pkg/httpserver/responses"
	"effectiveMobile/pkg/util/log/sl"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// DeleteUser godoc
// @Summary DeleteUser
// @Tag root
// @Description delete user by ID
// @ID delete_user
// @Accept json
// @Produce json
// @Param input body RequestDeleteUser true "user id"
// @Success 200 {boolean} ok
// @Failure 400,404,500 {object} responses.responsesStruct
// @Router / [delete]
func DeleteUser(log *slog.Logger, db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestDeleteUser
		if err := c.BindJSON(&reqBody); err != nil {
			responses.ErrorResponse(c, responses.ErrorBadRequest)
			return
		}
		log.Debug("deleting user")
		if err := db.DeleteUser(reqBody.ID); err != nil {
			if err.Error() == database.ErrorNotFound {
				responses.ErrorResponse(c, responses.ErrorNotFound)
				return
			}
			log.Error("could not delete user", sl.Err(err))
			responses.ErrorResponse(c, responses.ErrorServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true})
	}
}
