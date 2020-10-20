package main

import (
	"github.com/kriangkrai/SQL/RUNSQL/Router"
)

func main() {

	r := Router.SetupRouter()
	r.Run()

}
