package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type Page struct {
	Link        string
	Description string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/footer.html", "template/header.html")
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}

	data := struct {
		Title string
		Pages []Page
	}{
		Title: "Welcom to ATS System",
		Pages: []Page{
			{Link: "/", Description: "Main Page"},
			{Link: "/bootstrap", Description: "bootstrap test page"},
			{Link: "/invalid", Description: "Invalid http request page."},
			{Link: "/notfind", Description: "This should be the 404 page."},
			{Link: "/inputtest", Description: "This is a page for test html form."},
			{Link: "/registernewcase", Description: "Register a new ATS case."},
			{Link: "/index", Description: "This is also the main page."},
			{Link: "/pagefooter", Description: "This is also the main page."},
			{Link: "/formsubmit", Description: "This is a page for form submit example."},
			{Link: "/modularcase", Description: "Try to make the test case create page more flexibale."},
			{Link: "/newcase", Description: "Re-Design the case create function."},
		},
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
	}
}

func BootStrap(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/bootstrap.html", "template/footer.html", "template/header.html")
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
	log.Println(r.Method)
	if r.Method == "GET" {
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
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}

		log.Println(r.Form)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func PageFoorter(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/page.html", "template/footer.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}

		log.Println(r.Form)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func ModularCase(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/modularcase.html", "template/footer.html", "template/header.html", "template/task.html", "template/routine.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}

		log.Println(r.Form)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func NewCase(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/newcase.html", "template/footer.html", "template/header.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}

		log.Println(r.Form)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func FormSubmit(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/formsubmit.html", "template/footer.html", "template/header.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}

		log.Println(r.Form)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func InputTest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/inputtest.html", "template/footer.html", "template/header.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}

		log.Println(r.Form)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func ResourceNotFoundHandler(w http.ResponseWriter, r *http.Request) {

}

func InvalidReqMethodHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//@liwei: This need more analysis.
	http.HandleFunc("/", RedirectToMain)
	http.HandleFunc("/index", MainPage)
	http.HandleFunc("/notfind", ResourceNotFoundHandler)
	http.HandleFunc("/invalid", InvalidReqMethodHandler)
	http.HandleFunc("/bootstrap", BootStrap)
	http.HandleFunc("/inputtest", InputTest)
	http.HandleFunc("/formsubmit", FormSubmit)
	http.HandleFunc("/modularcase", ModularCase)
	http.HandleFunc("/newcase", NewCase)
	http.HandleFunc("/pagefooter", PageFoorter)
	http.HandleFunc("/registernewcase", RegisterNewCase)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
