package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/footer.html", "template/header.html")
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}

	data := struct {
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

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
	}
}

func RedirectToMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
}

func RegisterNewCase(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/registernewcase.html", "template/footer.html", "template/header.html")
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	//@liwei: This need more analysis.
	http.HandleFunc("/", RedirectToMain)
	http.HandleFunc("/index", MainPage)
	http.HandleFunc("/registernewcase", RegisterNewCase)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
