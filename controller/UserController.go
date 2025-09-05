package controller

import (
	"encoding/json"
	"github.com/yun/UserManger/models"
	"github.com/yun/UserManger/service"
	"github.com/yun/UserManger/untils"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "系统错误:"+err.Error()))
		return
	}
	if len(users) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusNotFound, "暂无用户数据"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(untils.Success("获取用户列表成功", users))
	return

}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusBadRequest, "请求方法错误"))
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusBadRequest, "表单解析错误:"+err.Error()))
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "系统错误:"+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(untils.Success("添加用户成功", models.ToVO(user)))
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusBadRequest, "请求方法错误"))
		return
	}

	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInsufficientStorage)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "获取token失败"))
		return
	}
	token, err := untils.ParseToken(cookie.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "解析token失败"))
		return
	}
	if token.Role != 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusForbidden, "无权限"))
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusBadRequest, "表单解析错误:"+err.Error()))
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	username := r.FormValue("username")
	password := r.FormValue("password")
	role, _ := strconv.Atoi(r.FormValue("role"))
	status, _ := strconv.Atoi(r.FormValue("status"))
	user := &models.User{
		ID:       id,
		Username: username,
		Password: password,
		Role:     role,
		Status:   status,
	}

	err = uc.UserService.UpdateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "系统错误:"+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(untils.Success("修改用户成功", ""))
}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusBadRequest, "请求方法错误"))
		return
	}

	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInsufficientStorage)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "获取token失败"))
		return
	}
	token, err := untils.ParseToken(cookie.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "解析token失败"))
		return
	}
	if token.Role != 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusForbidden, "无权限"))
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if token.UserID == id {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusForbidden, "不能删除自己"))
		return
	}
	err = uc.UserService.DeleteUser(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusInternalServerError, "系统错误:"+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(untils.Success("删除用户成功", ""))

}
