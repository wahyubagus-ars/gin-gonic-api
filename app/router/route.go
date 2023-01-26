package router

import (
	"gin-gonic-api/config"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	}

	return router
}
