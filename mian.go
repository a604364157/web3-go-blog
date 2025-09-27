package main

import (
	"web3-go-blog/config"
	"web3-go-blog/models"
	"web3-go-blog/router"
)

func main() {
	models.InitDB()
	r := router.SetupRouter()
	if err := r.Run(":" + config.Port); err != nil {
		panic("start web server error:" + err.Error())
	}
}
