package web

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer() (*Renderer, error) {
	rootPath := "assets\\templates"
	pfx := len(rootPath) + 1
	funcMap := template.FuncMap{}

	root := template.New("")

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".go.html") {
			if e1 != nil {
				return e1
			}

			b, e2 := os.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := strings.ReplaceAll(strings.ReplaceAll(path[pfx:], "\\", "/"), ".go.html", "")
			t := root.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Renderer{
		templates: root,
	}, nil
}

// Render renders a template document
func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	buffer := bytes.Buffer{}

	// Execute Page Template
	err := r.templates.ExecuteTemplate(&buffer, name, data)
	if err != nil {
		return err
	}
	pageHTML := buffer.String()

	layoutData := map[string]interface{}{
		"MetaData": MakeMetaData(),
		"PageHTML": template.HTML(pageHTML),
	}

	err = r.templates.ExecuteTemplate(w, "layout", layoutData)
	if err != nil {
		log.Printf("Error when rendering template: %v\n", err.Error())
	}
	return err
}

func (r *Renderer) RenderString(name string, data interface{}) (string, error) {
	var buffer bytes.Buffer
	err := r.templates.ExecuteTemplate(&buffer, name, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
