package controller

import (
	"encoding/json"
	"github.com/yun/UserManger/service"
	"github.com/yun/UserManger/utils"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type HomeController struct {
	HomeService *service.HomeService
	Tmpl        *template.Template
}

func NewHomeController(homeService *service.HomeService) *HomeController {
	return &HomeController{
		HomeService: homeService,
		Tmpl:        template.Must(template.ParseFiles("views/index.html")),
	}
}

func (hc *HomeController) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := hc.Tmpl.Execute(w, nil); err != nil {
		http.Error(w, "模板错误", http.StatusInternalServerError)
	}
}
func (hc *HomeController) UploadAvatar(w http.ResponseWriter, r *http.Request) {
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

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusBadRequest, "文件过大"))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	file, handler, _ := r.FormFile("avatar")
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Printf("file close error: %v", err)
			return
		}
	}(file)
	data, _ := io.ReadAll(file)
	id := r.Context().Value("userID").(int)
	avatarUrl, err := hc.HomeService.SaveAvatarFile(id, data, handler.Filename)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusInternalServerError, "上传失败:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	err = hc.HomeService.UpdateAvatarURL(avatarUrl, id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(utils.Fail[string](http.StatusInternalServerError, "更新失败:"+err.Error()))
		if err != nil {
			log.Printf("json encode error: %v", err)
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(utils.Success("上传成功", avatarUrl))
}

func (hc *HomeController) GetAvatarURL(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int)
	user, err := hc.HomeService.GetUserByID(id)
	if err != nil {
		http.Redirect(w, r, "https://cdn.jsdelivr.net/gh/feathericons/feather@main/icons/user.svg", http.StatusFound)
		return
	}
	filePath := "." + user.AvatarURL // ./static/avatar/1_xxxx.jpg
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Redirect(w, r, "https://cdn.jsdelivr.net/gh/feathericons/feather@main/icons/user.svg", http.StatusFound)
		return
	}

	// 禁止浏览器缓存
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// 再返回文件
	http.ServeFile(w, r, filePath)
}
