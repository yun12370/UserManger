package main

import (
	"github.com/gorilla/mux"
	"github.com/yun/UserManger/controller"
	"github.com/yun/UserManger/db"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/middleware"
	"github.com/yun/UserManger/service"
	"log"
	"net/http"
)

func main() {

	db.InitDB()

	loginMapper := mapper.NewLoginMapper(db.DB)
	loginService := service.NewLoginService(loginMapper)
	loginController := controller.NewLoginController(loginService)

	registerMapper := mapper.NewRegisterMapper(db.DB)
	registerService := service.NewRegisterService(registerMapper)
	registerController := controller.NewRegisterController(registerService)

	homeController := controller.NewHomeController()

	userMapper := mapper.NewUserMapper(db.DB)
	userService := service.NewUserService(userMapper)
	userController := controller.NewUserController(userService)

	router := mux.NewRouter()
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			loginController.LoginPage(w, r)
		} else if r.Method == http.MethodPost {
			loginController.LoginUser(w, r)
		}
	}).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			registerController.RegisterPage(w, r)
		} else if r.Method == http.MethodPost {
			registerController.RegisterUser(w, r)
		}
	})

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/users", userController.GetUsers).Methods(http.MethodGet)
	protected.HandleFunc("/user", userController.CreateUser).Methods(http.MethodPost)
	protected.HandleFunc("/user/{id}", userController.UpdateUser).Methods(http.MethodPut)
	protected.HandleFunc("/user/{id}", userController.DeleteUser).Methods(http.MethodDelete)
	protected.HandleFunc("/logout", loginController.Logout).Methods(http.MethodGet)
	protected.HandleFunc("/index", homeController.HomePage).Methods(http.MethodGet)

	chain := middleware.Chain(
		middleware.LoggerMiddleware,
		middleware.RecoverMiddleware,
	)
	log.Fatal(http.ListenAndServe(":8080", chain(router)))
}
