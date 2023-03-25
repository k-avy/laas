package api

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k-avy/laas/pkg/model"
)

func FindLicenses(c *gin.Context) {
	var licenses []model.License
	model.DB.Find(&licenses)

	c.JSON(http.StatusOK, gin.H{"data": licenses})
}

func CreateLicense(c *gin.Context) {
	// Validate input

	var input model.LicenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	license := model.License{ShortName: input.ShortName, LongName: input.LongName, Url: input.Url}
	model.DB.Create(&license)
  
	c.JSON(http.StatusOK, gin.H{"data": license})
}

func FindLicense(c *gin.Context) {  // Get model if exist
	var license model.License
	var queryParam string
	var ok bool
	if queryParam, ok = c.GetQuery("shortname"); ok {
		if queryParam == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound,
				gin.H{"Error: ": "Invalid startingIndex on search filter!"})
			c.Abort()
			return
		}

	if err := model.DB.Where("shortname = ?", queryParam).First(&license).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": license})
  }
  
}