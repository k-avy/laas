package app

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/k-avy/laas/license"
)
 
func RunRepositoryDemo(ctx context.Context, licenseRepository *license.PostgreSQLRepository){
	fmt.Println("1.MIGRATE REPOSITORY")
	if err:= licenseRepository.Migrate(ctx); err!= nil {
		log.Fatal(err)
	}
	fmt.Println("2. CREATE RECORDS OF LICENSE")
	first:= license.License{
		Shortname: "LGPL-2.1-or-later",
		Longname: "GNU Library General Public License v2.1 or later",
		Url: "https:www.gnu.org/licenses/old-licenses/lgpl-2.1-standalone.html",
	}
	second:=license.License{
		Shortname: "LGPL-3.0-or-later",
		Longname: "GNU Lesser General Public License v3.0 or later",
		Url: "https:www.gnu.org/licenses/lgpl-3.0-standalone.html",
	}
	createdfirst, err:= licenseRepository.Create(ctx, first)
	if errors.Is(err,license.ErrCheck) {
		log.Printf("record: %+v already exists\n",first)
	}else if err != nil {
        log.Fatal(err)
    }
	createdsecond, err:= licenseRepository.Create(ctx, second)
	if errors.Is(err, license.ErrCheck){
		log.Printf("record: %+v already exists\n", second)
	}else if err!= nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\n%+v\n", createdfirst, createdsecond)

	fmt.Println("3. GET RECORD BY NAME")
	gotfirst, err := licenseRepository.GetByName(ctx, "LGPL-2.1-or-later")
	if errors.Is(err,license.ErrNotExist){
		log.Printf("record: LGPL-2.1-or-later does not exist in the repository")
	} else if err!= nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", gotfirst)

	fmt.Println("5. GET ALL")
	all, err:= licenseRepository.All(ctx)
	if err!= nil {
		log.Fatal(err)
	}
	for _, license := range all{
		fmt.Printf("%+v\n", license)
	}

	fmt.Println("6. DELETE RECORD")
	if err:= licenseRepository.Delete(ctx, createdfirst.ID); err!=nil {
		if errors.Is(err, license.ErrDeleteFailed){
			fmt.Printf("delete of record: %d failed", createdfirst.ID)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("7. GET ALL")

	if err!= nil {
		log.Fatal(err)
	}
	for _, license := range all{
		fmt.Printf("%+v\n", license)
	}

}