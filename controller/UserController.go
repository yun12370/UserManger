package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yun/UserManger/models"
	"github.com/yun/UserManger/service"
	"github.com/yun/UserManger/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	page, pageSize := 1, 100
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
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusInternalServerError, "系统错误:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(utils.Success("获取用户列表成功", users))
	if err != nil {
		log.Printf("json encode error: %v", err)
		return
	}
	return

}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "请求方法错误"))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "表单解析错误:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))
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
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusInternalServerError, "创建失败:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(utils.Success("添加用户成功", models.ToVO(user)))
	if err != nil {
		log.Printf("json encode error: %v", err)
		return
	}
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "请求方法错误"))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "表单解析错误:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))
	role, _ := strconv.Atoi(r.FormValue("role"))
	status, _ := strconv.Atoi(r.FormValue("status"))
	user := &models.User{
		ID:       id,
		Username: username,
		Password: password,
		Role:     role,
		Status:   status,
	}

	userRole := r.Context().Value("role").(int)
	err = uc.UserService.UpdateUser(user, userRole)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusInternalServerError, "修改失败:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(utils.Success("修改用户成功", ""))
	if err != nil {
		log.Printf("json encode error: %v", err)
		return
	}
}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "请求方法错误"))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	userRole := r.Context().Value("role").(int)
	userID := r.Context().Value("userID").(int)

	err := uc.UserService.DeleteUser(id, userRole, userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusInternalServerError, "删除失败:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(utils.Success("删除用户成功", ""))
	if err != nil {
		log.Printf("json encode error: %v", err)
		return
	}
}
