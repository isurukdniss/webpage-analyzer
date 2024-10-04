package utils

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, templatePath string, data any) error { //TODO: test the any type

	t := template.Must(template.ParseFiles(templatePath))

	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}
