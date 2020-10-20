package Router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/SQL/RUNSQL/Controller"
	"github.com/kriangkrai/SQL/RUNSQL/Models"
)

func Insert(c *gin.Context) {
	var input Models.Data
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	insert := Controller.Insert(input)
	c.JSON(200, gin.H{
		"Message": insert,
	})
}

func Get(c *gin.Context) {
	datas := Controller.GetData()
	c.JSON(http.StatusOK, gin.H{
		"Data": datas,
	})
}

func Update(c *gin.Context) {
	var input Models.Data
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	update := Controller.Update(input)
	c.JSON(200, gin.H{
		"Message": update,
	})
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	delete := Controller.Delete(id)
	c.JSON(200, gin.H{
		"Message": delete,
	})
}
