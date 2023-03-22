package main

import (
	"context"
	"database/sql"
	"time"

	"log"

	"github.com/k-avy/laas/license"
	"github.com/k-avy/laas/app/classic"
)
func main(){
	db, err:= sql.Open("pqx","postgres:kavya:kavyash@localhost:5432/license")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	licenseRepository:= license.NewPostgreSQLRepository(db)

	ctx, cancel:= context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	app.RunRepositoryDemo(ctx, licenseRepository)
}