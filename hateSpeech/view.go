package hateSpeech

import (
	"html/template"
	"net/http"
)

func render (w http.ResponseWriter, tmpl string, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	handleError(err)
	err = t.Execute(w, nil)
}
