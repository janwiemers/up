package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/janwiemers/up/websockets"
)

func wsHandler(c *gin.Context) {
	websockets.ServeWs(websockets.HubInstance, c.Writer, c.Request)
}
