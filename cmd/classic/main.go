package main

import (
	"context"
	"database/sql"
	"time"

	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/k-avy/laas/app/classic"
	"github.com/k-avy/laas/license"
	"github.com/k-avy/laas/pkg/cmd"
)
func main(){
	db, err:= sql.Open("pgx","postgres://kavya:kavya123@localhost:5432/license")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	licenseRepository:= license.NewPostgreSQLRepository(db)

	ctx, cancel:= context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	app.RunRepositoryDemo(ctx, licenseRepository)
	cmd.Licenseapi()
}