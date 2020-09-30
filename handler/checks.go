package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janwiemers/up/database"
)

func checks(c *gin.Context) {
	id, _ := c.Params.Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	checks, err := database.Checks(i)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, checks)
}
