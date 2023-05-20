package render

import (
	"bytes"
	"fmt"
	"github.com/rstkhldntsk97/awesome-airbnb/pkg/config"
	"github.com/rstkhldntsk97/awesome-airbnb/pkg/model"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config to the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData) *model.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTmplCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser:", err)
	}
}

// CreateTmplCache creates a template cache as a map
func CreateTmplCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)
	pages, err := filepath.Glob("./template/*.page.gohtml")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./template/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
