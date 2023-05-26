package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	
)

type License struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	ShortName string `json:"shortname"`
	LongName  string `json:"longname"`
	Url       string `json:"url"`
}

type LicenseInput struct {
	ShortName string `json:"shortname" binding:"required"`
	LongName  string `json:"longname" binding:"required"`
	Url       string `json:"url" binding:"required"`
}

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}
// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}