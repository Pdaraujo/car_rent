package views

import (
	"bytes"
	"errors"
	"github.com/gorilla/csrf"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

const (
	LayoutDir   string = "views/layout/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)


type RendersView interface {
	Render(w http.ResponseWriter, r *http.Request, data interface{})
}

type View struct {
	Layout string
	Template *template.Template
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

func NewView(layout string, files... string) *View {
	createUsableFiles(files)
	files = append(files, getLayoutFiles()...)
	t, err := template.New("").Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", errors.New("csrfField is not implemented")
		},
	}).ParseFiles(files...)

	if err != nil {
		log.Fatal(err)
	}

	return &View{
		Layout:   layout,
		Template: t,
	}
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var buf bytes.Buffer
	csrfField := csrf.TemplateField(r)
	tpl := v.Template.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrfField
		},
	})
	err := tpl.ExecuteTemplate(&buf, v.Layout, nil)
	if err != nil {
		http.Error(w, "Something went wrong. If the problem "+
			"persists, please email support@lenslocked.com",
			http.StatusInternalServerError)
		log.Println(err)
		return
	}
	io.Copy(w, &buf)
}

func getLayoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func createUsableFiles(files []string) {
	for i, file := range files {
		files[i] = TemplateDir + file + TemplateExt
	}
}
