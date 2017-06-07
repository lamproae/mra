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
			{Link: "/bootcss", Description: "sample layout from bootcss."},
			{Link: "/pagefooter", Description: "This is also the main page."},
			{Link: "/formsubmit", Description: "This is a page for form submit example."},
			{Link: "/modularcase", Description: "Try to make the test case create page more flexibale."},
			{Link: "/newcase", Description: "Re-Design the case create function."},
			{Link: "/precondition", Description: "PreCondition."},
			{Link: "/postcondition", Description: "PostCondition."},
			{Link: "/taskroutine", Description: "TaskRoutine."},
			{Link: "/prepostroutine", Description: "PrePostRoutine."},
			{Link: "/stepforward", Description: "StepForward."},
			{Link: "/newnewtask", Description: "NewNewTask."},
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
		//@liwei: This is a very stuiped method. We need just return the opertion status to user.
		t, err := template.ParseFiles("template/newtask.html", "template/footer.html", "template/header.html", "template/routine.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title string
			ID    string
		}{
			Title: "NewTask",
			ID:    "1",
		})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func ShowTask(w http.ResponseWriter, r *http.Request) {

}

func NewTask(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/newtask.html", "template/footer.html", "template/header.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, &struct {
			Title string
			ID    string
		}{
			Title: "NewTask",
			ID:    "1",
		})
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}
		http.Redirect(w, r, "/newtask", http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func NewNewTask(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/newnewtask.html", "template/footer.html", "template/header.html", "template/newroutine.html", "template/condition.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, &struct {
			Title string
			ID    string
		}{
			Title: "NewTask",
			ID:    "1",
		})
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

func BootCSS(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	t, err := template.ParseFiles("template/bootcss_layout.html", "template/footer.html", "template/header.html")
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}

	err = t.Execute(w, &struct {
		ID string
	}{
		ID: "1",
	})
	if err != nil {
		log.Println(err.Error())
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

func PreCondition(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/precondition.html", "template/footer.html", "template/header.html", "template/routine.html", "template/stepforward.html")
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

func PostCondition(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/postcondition.html", "template/footer.html", "template/header.html", "template/routine.html", "template/stepforward.html")
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

func TaskRoutine(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/taskroutine.html", "template/footer.html", "template/header.html", "template/routine.html", "template/stepforward.html")
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

func PrePostRoutine(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/prepostroutine.html", "template/footer.html", "template/header.html", "template/routine.html", "template/stepforward.html")
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

func StepForward(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/stepforward.html", "template/footer.html", "template/header.html")
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
	http.HandleFunc("/newtask", NewTask)
	http.HandleFunc("/bootcss", BootCSS)
	http.HandleFunc("/pagefooter", PageFoorter)
	http.HandleFunc("/registernewcase", RegisterNewCase)
	http.HandleFunc("/precondition", PreCondition)
	http.HandleFunc("/postcondition", PostCondition)
	http.HandleFunc("/taskroutine", TaskRoutine)
	http.HandleFunc("/stepforward", StepForward)
	http.HandleFunc("/newnewtask", NewNewTask)
	http.HandleFunc("/prepostroutine", PrePostRoutine)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
