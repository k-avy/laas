package main

import (

	"github.com/gin-gonic/gin"
	"github.com/k-avy/laas/pkg/api"
	"github.com/k-avy/laas/pkg/model"
)

func main() {
  r := gin.Default()

  model.ConnectDatabase()
	r.GET("/licenses", api.FindLicenses)
	r.POST("/licenses", api.CreateLicense)
	r.GET("/license", api.FindLicense)
	
  r.Run()
}