package handler

import (
	"net/http"

	"github.com/VJ-Vijay77/LoginPageNew/pkg/config"
	"github.com/VJ-Vijay77/LoginPageNew/pkg/models"
	"github.com/VJ-Vijay77/LoginPageNew/pkg/render"
)

var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}


func(m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)
	
	render.RenderTemplate(w,"home.gohtml",&models.TemplateData{})
}

func(m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP	

	render.RenderTemplate(w,"about.gohtml",&models.TemplateData{
		StringMap: stringMap,
	})
	
}
