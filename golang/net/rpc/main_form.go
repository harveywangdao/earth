package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hey)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

const tpl = `
<html>
	<head>
		<title>Hey</title>
	</head>

	<body>
		<form method="post" action="/">
			Username: <input type="text" name="uname">
			Password: <input type="password" name="pwd">
			<button type="submit">Submit</button>
		</form>
	</body>
</html>
`

func Hey(w http.ResponseWriter, r *http.Request) {
	log.Print("111")
	if r.Method == "GET" {
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)
	} else {
		fmt.Println(r.FormValue("uname") + "   " + r.FormValue("pwd"))
	}
}
