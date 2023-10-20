package updateuser

import (
	"effectiveMobile/pkg/database"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func UpdateUser(log *slog.Logger, db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
