package main

import (
	"backend_template/config"
	"backend_template/router"
	)


func main() {
	router:=router.InitRouter()
	config.ConnectDataBase()
	router.Run(":8080")
	defer config.DB.Close()
}