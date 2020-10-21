package Router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/SQL/RUNSQL/Controller"
	"github.com/kriangkrai/SQL/RUNSQL/Models"
)

func GetLocation(c *gin.Context) {
	location := Controller.GetLocation()
	// c.JSON(http.StatusOK, gin.H{
	// 	"Location": location,
	// })

	c.JSON(200, location)
}

func InsertLocation(c *gin.Context) {
	var input Models.Location
	e := c.ShouldBindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	_ = Controller.InsertLocation(input)
	// c.JSON(200, gin.H{
	// 	"Message": input,
	// })
	c.JSON(200, input)
}

func UpdateLocation(c *gin.Context) {
	var input Models.Location
	e := c.ShouldBindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	update := Controller.UpdateLocation(input)
	// c.JSON(200, gin.H{
	// 	"Message": update,
	// })
	c.JSON(200, update)
}

func DeleteLocation(c *gin.Context) {
	id := c.Params.ByName("id")
	delete := Controller.DeleteLocation(id)
	// c.JSON(200, gin.H{
	// 	"Message": delete,
	// })

	c.JSON(200, delete)
}
