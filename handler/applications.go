package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janwiemers/up/database"
)

func applications(c *gin.Context) {
	apps := database.Applications()
	c.JSON(http.StatusOK, apps)
}
