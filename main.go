package main

import (
	"control/router"
)

func main() {
	s := router.GetServer()
	s.Run(":" + "9000")
}
