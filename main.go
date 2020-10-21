package main

import (
	"github.com/kriangkrai/SQL/RUNSQL/Router"
)

func main() {

	r, port := Router.SetupRouter()
	r.Run(":" + port)

}
