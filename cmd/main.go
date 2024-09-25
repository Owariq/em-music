package main

import (
	"log"

	_ "github.com/Owariq/em-music/docs"
	"github.com/Owariq/em-music/internal/db"
	"github.com/Owariq/em-music/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//	@title			Music Service
//	@version		1.0
//	@description	API for music service
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/lib

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	

	r := gin.Default()



	err = db.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	routes.Routes(r)

	log.Fatal(r.Run(":8080"))

}