package hateSpeech

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
}

func ProcessSpeech (w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	sentence := r.FormValue("sentence")
	_ = sentence
	newSpeech := speech{
		body: "Buhari",
	}

	if newSpeech.isRelatedToPresident() {
		_, _ = fmt.Fprintf(w, "You entered something related to the president!")
	}
}