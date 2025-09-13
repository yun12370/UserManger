package controller

import (
	"encoding/json"
	"github.com/yun/UserManger/service"
	"github.com/yun/UserManger/utils"
	"html/template"
	"log"
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
		Tmpl:         template.Must(template.ParseFiles("views/login.html")),
	}
}

func (lc *LoginController) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "请求方式错误"))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	token, err := lc.LoginService.LoginUser(username, password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "登陆失败:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
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
	err = json.NewEncoder(w).Encode(utils.Success("登录成功", token))
	if err != nil {
		log.Printf("json encode error: %v", err)
		return
	}
}
func (lc *LoginController) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := lc.Tmpl.Execute(w, nil); err != nil {
		http.Error(w, "模板错误", http.StatusInternalServerError)
	}
}

func (lc *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
