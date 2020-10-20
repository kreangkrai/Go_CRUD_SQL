package main

import (
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/SQL/Controller"
	"github.com/kriangkrai/SQL/Router"
)

func main() {

	Controller.Connect()
	//version := Controller.CheckVersion()
	//fmt.Println(version)
	//update := Controller.Update(8, 31.9)
	//fmt.Println(update)

	//insert := Controller.Insert("Sensor2", "2020-10-19", 25.9)
	//fmt.Println(insert)
	//datas := Controller.Select()

	// for i := range datas {
	// 	fmt.Println(datas[i].Device + "," + datas[i].Date + "," + strconv.FormatFloat(datas[i].Value, 'f', 2, 64))
	// }

	//Router.InitRoute(datas)

	r := gin.Default()
	userRoute := r.Group("/user")
	userRoute.Use(Router.UserMiddleware())
	{
		userRoute.POST("/insert", Router.Insert)
		userRoute.GET("/get", Router.Get)
		userRoute.PUT("/update", Router.Update)
		userRoute.DELETE("/delete/:id", Router.Delete)
	}

	LocationRoute := r.Group("/location")
	LocationRoute.Use(Router.LocationMiddleware())
	{
		LocationRoute.GET("/get", Router.GetLocation)
		LocationRoute.POST("/insert", Router.InsertLocation)
		LocationRoute.PUT("/update", Router.UpdateLocation)
		LocationRoute.DELETE("/delete/:id", Router.DeleteLocation)
	}

	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	r.Run(":" + port)

}
