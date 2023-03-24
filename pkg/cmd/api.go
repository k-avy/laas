package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-avy/laas/pkg/api"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func Licenseapi(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",homeLink)
	router.HandleFunc("/get",api.GetLicense).Methods("GET")
	router.HandleFunc("/getbyname", api.GetLicenseByName).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}