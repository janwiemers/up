package handler

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter will set up a gin router
func SetupRouter(r *gin.Engine) {
	r.GET("/ping", ping)
	// r.GET("/applications", applications)
	// r.GET("/application/:id/checks", checks)
	r.GET("/ws", wsHandler)
}
