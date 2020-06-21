package main

import (
	"control/controller"
	"control/router"
)

func main() {
	controller.NewPrometheusClient()
	s := router.GetServer()
	s.Run(":" + "9000")
}
