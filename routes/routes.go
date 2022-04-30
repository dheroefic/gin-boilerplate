package routes

import (
	"fmt"
	"net/http"

	"github.com/dheroefic/gin-boilerplate/models/structs"
	"github.com/dheroefic/gin-boilerplate/utils/helpers"
	"github.com/gin-gonic/gin"
)

func Load(route *gin.Engine) {
	// 404 states
	route.NoRoute(func(ctx *gin.Context) {
		helpers.Logger("ROUTING PROCESS", fmt.Sprintf("accessing undefined endpoint: %s", ctx.Request.RequestURI), false)
		obj := structs.APIResponse{Code: http.StatusNotFound, Message: "Endpoint is not found"}
		ctx.SecureJSON(http.StatusNotFound, obj)
	})

	// Health
	// Default health-check route
	route.GET("/health", func(ctx *gin.Context) {
		obj := structs.APIResponse{Code: 1, Message: "System is normal"}
		ctx.SecureJSON(http.StatusOK, obj)
	})

}
