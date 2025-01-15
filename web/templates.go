package web

import (
	"embed"
	"html/template"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// Embed all files in the templates directory.
//
//go:embed templates/**/*
var templatesFS embed.FS

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c *gin.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *Templates) Instance(name string, data any) render.Render {
	return &templateRender{
		Template: t.templates.Lookup(name),
		Data:     data,
	}
}

type templateRender struct {
	Template *template.Template
	Data     any
}

func (tr *templateRender) Render(w http.ResponseWriter) error {
	tr.WriteContentType(w)
	return tr.Template.Execute(w, tr.Data)
}

func (tr *templateRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func NewTemplate() *Templates {
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/**/*"))

	return &Templates{
		templates: tmpl,
	}
}
