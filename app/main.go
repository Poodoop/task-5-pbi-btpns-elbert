package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/username/app/controllers"
	"github.com/username/app/database"
	"github.com/username/app/router"
)

func main() {
	db := database.Init()
	defer db.Close()

	database.RunMigrations(db)

	r := gin.Default()

	router.SetupRoutes(r)

	r.Run(":8080")
}
