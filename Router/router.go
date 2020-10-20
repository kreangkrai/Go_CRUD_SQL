package Router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/SQL/RUNSQL/Controller"
)

func SetupRouter() *gin.Engine {
	Controller.Connect()
	r := gin.Default()
	userRoute := r.Group("/user")
	userRoute.Use(UserMiddleware())
	{
		userRoute.POST("/insert", Insert)
		userRoute.GET("/get", Get)
		userRoute.PUT("/update", Update)
		userRoute.DELETE("/delete/:id", Delete)
	}

	LocationRoute := r.Group("/location")
	LocationRoute.Use(LocationMiddleware())
	{
		LocationRoute.GET("/get", GetLocation)
		LocationRoute.POST("/insert", InsertLocation)
		LocationRoute.PUT("/update", UpdateLocation)
		LocationRoute.DELETE("/delete/:id", DeleteLocation)
	}

	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	r.Run(":" + port)

	return r
}
