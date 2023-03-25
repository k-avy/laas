package model

type License struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  ShortName string `json:"shortname"`
  LongName  string `json:"longname"`
  Url string `json:"url"`
}

 type LicenseInput struct{
  ShortName string `json:"shortname" binding:"required"`
  LongName  string `json:"longname" binding:"required"`
  Url string `json:"url" binding:"required"`

 }