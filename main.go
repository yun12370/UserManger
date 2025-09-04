package main

import (
	"github.com/gorilla/mux"
	"github.com/yun/UserManger/controller"
	"github.com/yun/UserManger/db"
	"github.com/yun/UserManger/mapper"
	"github.com/yun/UserManger/service"
	"net/http"
)

func main() {

	db.InitDB()

	userMapper := mapper.NewUserMapper(db.DB)
	userService := service.NewUserService(userMapper)
	userController := controller.NewUserController(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users", userController.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/createUser", userController.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/updateuser", userController.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/deleteuser", userController.DeleteUser).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", router)

}
