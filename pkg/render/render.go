package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/VJ-Vijay77/LoginPageNew/pkg/config"
	"github.com/VJ-Vijay77/LoginPageNew/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//sets the config for the template package
func NewTemplate(a *config.AppConfig){
	app = a
}

 func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	
	if app.UseCache{
	tc = app.TemplateCache
	}else{
		tc ,_ = CreateTemplateCache()
	}
	
	t,ok := tc[tmpl]
	if !ok {
		log.Fatalln("Could not get template from tempate cache")
	}

	buf := new(bytes.Buffer)

	_=t.Execute(buf,td)

	_,err := buf.WriteTo(w)
	if err != nil{
		fmt.Println(err)
	}

}
  

func CreateTemplateCache()(map[string]*template.Template,error){
	templatecache := map[string] *template.Template{}

	page , err := filepath.Glob("./templates/*.gohtml")
	if err != nil {
		return templatecache,err
	}
	for _,page := range page {
		name := filepath.Base(page)
		 ts,err := template.New(name).Funcs(functions).ParseFiles(page)
		 if err != nil {
			return templatecache,err 
		 }
		 matches,err := filepath.Glob("./templates/*.layout.gohtml")
		 if err != nil {
			return templatecache,err
		 }
		 if len(matches) > 0 {
			 ts ,err = ts.ParseGlob("./templates/*.layout.gohtml")
			 if err != nil {
				return templatecache,err   
			 }
		 }
	templatecache[name] =ts
		}

		return templatecache,nil
		 
}

