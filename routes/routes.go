package routes

import (
	"fmt"
	"net/http"

	"github.com/dheroefic/gin-boilerplate/database"
	"github.com/dheroefic/gin-boilerplate/models/structs"
	"github.com/dheroefic/gin-boilerplate/utils/helpers"
	"github.com/gin-gonic/gin"
)

// All route defined here
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
		db := database.GetSession().Begin()
		obj := structs.APIResponse{Code: http.StatusOK, Message: "System is normal"}
		if err := db.Exec("select 1 + 1").Error; err != nil {
			obj.Code = http.StatusInternalServerError
			obj.Message = "There's something wrong with the system"
		}
		db.Commit()
		ctx.SecureJSON(http.StatusOK, obj)
	})

}
