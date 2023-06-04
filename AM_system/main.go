package main

import (
	"AM_system/model"
	"AM_system/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
