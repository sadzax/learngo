package routes

import (
	"activity-tracker/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// r.LoadHTMLGlob("views/templates/*")

	api := r.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:user_id", controllers.GetUserWithActivities)

		api.GET("/users/:user_id/activities", controllers.GetUserActivities)

		api.POST("/activities", controllers.CreateActivity)
		api.PUT("/activities/:id", controllers.UpdateActivity)
		api.DELETE("/activities/:id", controllers.DeleteActivity)

		api.POST("/track", controllers.LogActivityHours)
	}

	// HTML view
	// r.GET("/users/:user_id/dashboard", controllers.ShowUserDashboard)

	return r // нужно всегда возвращать gin.Default ?
}
