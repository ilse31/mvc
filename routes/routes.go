package routes

import (
	"mvc/config"
	"mvc/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes() *echo.Echo {
	echoApp := echo.New()

	echoApp.DELETE("/users/:id", controllers.DeletedUserByID)
	echoApp.GET("/users/:id", controllers.GetUserbyId)
	echoApp.GET("/users", controllers.GetUserController)
	echoApp.POST("/users", controllers.SaveUserController)
	echoApp.PUT("/users/:id", controllers.UpdateUsersController)

	// echoAppBasic := echoApp.Group("/auth")
	// echoAppBasic.GET("/users", controllers.GetUserController)
	// echoAppBasic.Use(mid.BasicAuth(middleware.BasicAuthDb))

	//jwt

	echoApp.GET("/token", controllers.GetUserToken)
	echoApp.GET("/withToken", controllers.HelloWorld, middleware.JWT([]byte(config.JWT_SECRET)))

	return echoApp
}
