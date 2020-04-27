package main

import (

	"github.com/labstack/echo"
	"github.com/linda/auth/infrastructure/datastore"
	"github.com/linda/auth/infrastructure/router"
	"github.com/linda/auth/registry"

)

func main() {
	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())


	e.Logger.Fatal(e.Start(":1323"))
}