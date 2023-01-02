package main

import (
	"log"
	"main/src/keycloak"
	"main/src/router"
	"main/src/services"
	"main/src/utils"
	"net/http"

	//"github.com/jmoiron/sqlx"
	//"testing/quick"
)

func main() {
	//fmt.Println("Hello World")


	db := utils.OpenDBConnection()

	services.SetDB(db)

	var myRouter = router.CreateRouter(keycloak.NewKeycloak())

	log.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}