package main

import (
	"event-app/config"
	"event-app/routes"

	_ "event-app/docs" // Import Swagger docs
)

// @title Event App API
// @version 1.0
// @description This is a simple API for managing events.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email youremail@domain.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()
	r.Run(":8080")
}
