package main

import (
	"github.com/gorilla/mux"
	"github.com/yun/UserManger/controller"
	"github.com/yun/UserManger/db"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/middleware"
	"github.com/yun/UserManger/service"
	"net/http"
)

func main() {

	db.InitDB()

	loginMapper := mapper.NewLoginMapper(db.DB)
	loginService := service.NewLoginService(loginMapper)
	loginController := controller.NewLoginController(loginService)

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

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/users", userController.GetUsers).Methods(http.MethodGet)
	protected.HandleFunc("/createUser", userController.CreateUser).Methods(http.MethodPost)
	protected.HandleFunc("/updateuser", userController.UpdateUser).Methods(http.MethodPut)
	protected.HandleFunc("/deleteuser", userController.DeleteUser).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", router)

}
