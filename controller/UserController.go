package controller

import (
	"encoding/json"
	"github.com/yun/UserManger/models"
	"github.com/yun/UserManger/service"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	page, pageSize := 1, 1
	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if ps := r.URL.Query().Get("pageSize"); ps != "" {
		pageSize, _ = strconv.Atoi(ps)
	}

	users, err := uc.UserService.GetUsers(page, pageSize)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(users)

}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode("请求方法错误！")
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {

		//fmt.Sprint(err)
		//fmt.Print(err)
		json.NewEncoder(w).Encode("表单解析错误！")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	role := 2
	status := 1
	user := &models.User{
		Username: username,
		Password: password,
		Role:     role,
		Status:   status,
	}

	err = uc.UserService.CreateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode("添加用户成功！")
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		return
	}
	role := 2
	status := 1
	user := &models.User{
		Username: "username",
		Password: "password",
		Role:     role,
		Status:   status,
	}
	err := uc.UserService.UpdateUser(user)
	if err != nil {
		return
	}
	return
}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		return
	}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := uc.UserService.DeleteUser(id)
	if err != nil {
		return
	}
	return
}
