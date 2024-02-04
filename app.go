package main

import (
	"api/database"
	"api/models"
	"api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type App struct {
	Router *mux.Router
}

func initDatabase() {

	if err := database.Open(); err != nil {
		panic("Fail to connect to database")
	}
}

func (a *App) Initialize() {

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	c := cors.New(cors.Options{
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowedOrigins:     []string{"*"},
	})

	initDatabase()

	models.DB = database.Connect
	a.Router = routers.LoadRouter()

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(a.Router)

}

func (a *App) Run(port string) {
	fmt.Print("http://localhost" + port + "\n")
	log.Fatal(http.ListenAndServe(port, a.Router))
}
