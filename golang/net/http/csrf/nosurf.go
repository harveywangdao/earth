package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

var templateString string = `
<!doctype html>
<html>
<body>
{{ if .name }}
<p>Your name: {{ .name }}</p>
{{ end }}
<form action="/" method="POST">
<input type="text" name="name">

<!-- Try removing this or changing its value
     and see what happens -->
<input type="hidden" name="csrf_token" value="{{ .token }}">
<input type="submit" value="Send">
</form>
</body>
</html>
`
var templ = template.Must(template.New("t1").Parse(templateString))

func myFunc(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]string)
	context["token"] = nosurf.Token(r)
	if r.Method == "POST" {
		context["name"] = r.FormValue("name")
	}

	templ.Execute(w, context)
}

func do1() {
	myHandler := http.HandlerFunc(myFunc)
	http.ListenAndServe(":8000", nosurf.New(myHandler))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
