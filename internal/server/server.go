package server

import (
	"log"

	"github.com/fentezi/api-rate/config"
	"github.com/fentezi/api-rate/internal/controllers"
	"github.com/fentezi/api-rate/internal/db"
	"github.com/fentezi/api-rate/internal/repositories"
	"github.com/fentezi/api-rate/internal/services"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	cfg := config.NewConfig()
	db, err := db.InitDatabase(cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	rep := repositories.NewRepository(db)
	service := services.NewService(*rep, cfg)
	controller := controllers.NewController(service)

	router := gin.Default()
	router.GET("/rate", controller.GetRate)
	router.POST("/subscribe", controller.Subscribe)
	router.POST("/sendEmails", controller.Mailing)

	router.Run()
}
