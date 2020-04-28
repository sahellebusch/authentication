package main

import (
	"authentication/models"
	"authentication/router"
)

func main() {
	db := models.SetupModels()
	defer db.Close()

	router := router.SetupRouter(db)
	router.Run() // listen and serve on 0.0.0.0:8080
}
