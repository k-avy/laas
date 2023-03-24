package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type License struct{
	Id int64 `json:"id"`
	ShortName string `json:"shortname"`
	LongName string `json:"longname"`
	Url string `json:"url"`

}
type JsonResponse struct {
    Type    string `json:"type"`
    Data    []License `json:"data"`
    Message string `json:"message"`
}

const(
	DB_USER     = "kavya"
    DB_PASSWORD = "kavyash"
    DB_NAME     = "licenses"
)

func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    if err!=nil{
		panic(err)
	}
    return db
}
func GetLicense(w http.ResponseWriter, r *http.Request){
	db:= setupDB()
	rows, err := db.Query("SELECT * FROM licenses")
	if err!=nil{
		panic(err)
	}

	var licenses []License

	for rows.Next() {
        var id int64
        var shortname string
        var longname string
		var url string

        err = rows.Scan(&id,&shortname,&longname,&url)
		if err!=nil{
			panic(err)
		}
		licenses= append(licenses, License{id,shortname,longname,url})
	}
		var response = JsonResponse{
			Type:    "success",
			Data:    licenses,
			Message: "",
		}

		json.NewEncoder(w).Encode(response)
}

func GetLicenseByName(w http.ResponseWriter, r *http.Request){
	Name:= r.URL.Query().Get("name")
	db:= setupDB()
	rows, err := db.Query("SELECT * FROM licenses WHERE shortname=$1",Name)
	if err!=nil{
		panic(err)
	}
	var id int64
    var shortname string
    var longname string
	var url string

    err = rows.Scan(&id,&shortname,&longname,&url)
	if err!=nil{
		panic(err)
	}
	license:= License{id,shortname,longname,url}
	
	foo_marshalled, err1:= json.Marshal(license)
	
	if err1!=nil{
		panic(err1)
	}
	fmt.Fprint(w, string(foo_marshalled))
	
}
