package controller

import (
	"encoding/json"
	"github.com/yun/UserManger/service"
	"github.com/yun/UserManger/untils"
	"html/template"
	"net/http"
	"time"
)

type LoginController struct {
	LoginService *service.LoginService
	Tmpl         *template.Template
}

func NewLoginController(loginService *service.LoginService) *LoginController {
	return &LoginController{
		LoginService: loginService,
		Tmpl:         template.Must(template.ParseFiles("view/login.html")),
	}
}

func (lc *LoginController) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusBadRequest, "请求方式错误"))
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	token, err := lc.LoginService.LoginUser(username, password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(untils.Fail[any](http.StatusBadRequest, err.Error()))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * 24),
	})
	json.NewEncoder(w).Encode(untils.Success("登录成功", token))
}
func (lc *LoginController) LoginPage(w http.ResponseWriter, r *http.Request) {
	if err := lc.Tmpl.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, "模板错误", http.StatusInternalServerError)
	}
}
