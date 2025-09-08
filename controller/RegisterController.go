package controller

import (
	"encoding/json"
	"github.com/yun/UserManger/service"
	"github.com/yun/UserManger/utils"
	"html/template"
	"net/http"
)

type RegisterController struct {
	RegisterService *service.RegisterService
	Tmpl            *template.Template
}

func NewRegisterController(registerService *service.RegisterService) *RegisterController {
	return &RegisterController{
		RegisterService: registerService,
		Tmpl:            template.Must(template.ParseFiles("views/register.html")),
	}
}
func (rc *RegisterController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Fail[any](http.StatusBadRequest, "请求方法错误"))
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	err := rc.RegisterService.RegisterUser(username, password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Fail[any](http.StatusInternalServerError, "注册失败:"+err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.Success("注册成功", ""))
}

func (rc *RegisterController) RegisterPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := rc.Tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "模板错误", http.StatusInternalServerError)
	}
}
