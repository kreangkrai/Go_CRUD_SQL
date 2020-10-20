package Router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/SQL/RUNSQL/Controller"
	"github.com/kriangkrai/SQL/RUNSQL/Models"
)

func GetLocation(c *gin.Context) {
	location := Controller.GetLocation()
	c.JSON(http.StatusOK, gin.H{
		"Location": location,
	})
}

func InsertLocation(c *gin.Context) {
	var input Models.Location
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	insert := Controller.InsertLocation(input)
	c.JSON(200, gin.H{
		"Message": insert,
	})
}

func UpdateLocation(c *gin.Context) {
	var input Models.Location
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	update := Controller.UpdateLocation(input)
	c.JSON(200, gin.H{
		"Message": update,
	})
}

func DeleteLocation(c *gin.Context) {
	id := c.Params.ByName("id")
	delete := Controller.DeleteLocation(id)
	c.JSON(200, gin.H{
		"Message": delete,
	})
}
