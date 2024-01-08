package main

import (
	"log"
	"net/http"
	"subproblem/rest-api/pkg/controller"
	"subproblem/rest-api/pkg/db"
	"subproblem/rest-api/pkg/service"
	"github.com/gorilla/mux"
	"subproblem/rest-api/pkg/util"
)

func main() {


	db, err := db.NewPostgresDb()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	

	util.LoadEnv()


	userService := service.NewUserService(db)
	authService := service.NewAuthService(db)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/register", authController.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", authController.Login).Methods("POST")

	// Secured Endpoints
	r.HandleFunc("/api/v1/user/{id}", authController.Middleware(http.HandlerFunc(userController.GetUserById))).Methods("GET")
	

	http.ListenAndServe(":8080", r)
}