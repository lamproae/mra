package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>{{.Title}}</title>
		<script type="text/javascript">
		alert("Hello world")
		</script>
	</head>
	</body>
		{{range .Items}}
			<div>{{.}}</div>
		{{else}}
			<div>no row</div>
		{{end}}
	</body>
</html>`

var data = struct {
	Title string
	Items []string
}{
	Title: "First Page",
	Items: []string{
		"My photos",
		"My movie",
		"My video",
		"My audio",
	},
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("MainPage").Parse(tpl)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", MainPage)
	http.ListenAndServe(":8080", nil)
}
