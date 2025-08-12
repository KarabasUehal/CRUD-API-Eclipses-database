package main

import (
	"golang-gin/handlers"
	"golang-gin/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := storage.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	router := gin.Default()

	router.GET("/eclipse", handlers.GetAllEclipses)
	router.GET("/eclipse/:id", handlers.GetEclipseByID)
	router.POST("/eclipse/add", handlers.AddEclipse)
	router.PUT("/eclipse/update/:id", handlers.UpdateEclipseByID)
	router.DELETE("/eclipse/delete/:id", handlers.DeleteEclipseByID)

	router.Run(":3000")

}
