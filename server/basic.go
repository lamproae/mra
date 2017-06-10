package main

import (
	"ccase"
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Link        string
	Description string
}

func ADD(a, b int) string {
	return strconv.Itoa(a + b)
}

func Encap(i interface{}, a, b string) interface{} {
	return struct {
		o interface{}
		a string
	}{
		o: i,
		a: b,
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/footer.html", "template/header.html")
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
		t, err := template.ParseFiles("template/registernewcase.html", "template/footer.html", "template/header.html", "template/casenavigator.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title string
			DB    *ccase.CaseDBInMem
		}{
			Title: "Device test cases",
			DB:    DB,
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/dumpcase.html", "template/footer.html", "template/header.html", "template/caseheader.html", "template/casenavigator.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		r.ParseForm()
		log.Println(r.Form)

		c, err := DB.Get(&ccase.Case{
			Group:    r.FormValue("group"),
			SubGroup: r.FormValue("sgroup"),
			Feature:  r.FormValue("feature"),
			Name:     r.FormValue("case"),
		})

		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		err = c.DelTask(&ccase.Task{
			Name: r.FormValue("task"),
		})

		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		DB.Save()

		err = t.Execute(w, struct {
			Title string
			Case  *ccase.Case
			DB    *ccase.CaseDBInMem
		}{
			Title: "Del Task",
			Case:  c,
			DB:    DB,
		})
		if err != nil {
			log.Println(err.Error())
		}

	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func DumpCase(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/dumpcase.html", "template/footer.html", "template/header.html", "template/caseheader.html", "template/casenavigator.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		r.ParseForm()
		log.Println(r.Form)

		c, err := DB.Get(&ccase.Case{
			Group:    r.FormValue("group"),
			SubGroup: r.FormValue("sgroup"),
			Feature:  r.FormValue("feature"),
			Name:     r.FormValue("name"),
		})

		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title string
			Case  *ccase.Case
			DB    *ccase.CaseDBInMem
		}{
			Title: "Create new Test Case",
			Case:  c,
			DB:    DB,
		})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func DumpTask(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/dumptask.html", "template/footer.html", "template/header.html", "template/taskheader.html", "template/casenavigator.html", "template/dumpcondition.html", "template/dumproutine.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		r.ParseForm()
		log.Println(r.Form)

		c, err := DB.Get(&ccase.Case{
			Group:    r.FormValue("group"),
			SubGroup: r.FormValue("sgroup"),
			Feature:  r.FormValue("feature"),
			Name:     r.FormValue("case"),
		})

		task := c.GetTask(r.FormValue("task"))

		if task == nil {
			io.WriteString(w, "Cannot find task: "+r.FormValue("task"))
			return
		}

		err = t.Execute(w, struct {
			Title string
			Case  *ccase.Case
			DB    *ccase.CaseDBInMem
			Task  *ccase.Task
		}{
			Title: "Dump Case Task",
			Case:  c,
			DB:    DB,
			Task:  task,
		})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func NewCase(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/newcase.html", "template/footer.html", "template/header.html", "template/casenavigator.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title string
			DB    *ccase.CaseDBInMem
		}{
			Title: "Create new Test Case",
			DB:    DB,
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

		newcase, err := ccase.CreateNewCase(r.Form)
		if err != nil {
			log.Println("Cannot create new Case: ", err)
		} else {
			log.Printf("%#v", newcase)
			DB.Add(newcase)
			DB.Save() //This should be more flexible
			log.Printf("%v", DB.Dump())
			log.Printf("%#v", DB.Dump()[0].DUTs[0])
			log.Println(DB)
		}
		//@liwei: This is a very stuiped method. We need just return the opertion status to user.
		t, err := template.ParseFiles("template/newcase.html", "template/footer.html", "template/header.html", "template/casenavigator.html")
		//t, err := template.ParseFiles("template/newnewtask.html", "template/footer.html", "template/header.html", "template/newroutine.html", "template/condition.html", "template/casenavigator.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title   string
			ID      string
			Device  string
			Group   string
			Feature string
			Case    string
			DB      *ccase.CaseDBInMem
		}{
			Title:   "NewTask",
			ID:      "1",
			Device:  "V8500",
			Group:   "L2",
			Feature: "VLAN",
			Case:    "Vlan Create",
			DB:      DB,
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

func EditTask(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	err := r.ParseForm()
	if err != nil {
		log.Println("Cannot parse form: ", err.Error())
		return
	}

	log.Println(r.Form)

	c, err := DB.Get(&ccase.Case{
		Group:    r.FormValue("group"),
		SubGroup: r.FormValue("sgroup"),
		Feature:  r.FormValue("feature"),
		Name:     r.FormValue("name"),
	})

	if r.Method == "GET" {
		t, err := template.ParseFiles("template/dumptask.html", "template/footer.html", "template/header.html", "template/taskheader.html", "template/casenavigator.html", "template/dumpcondition.html", "template/dumproutine.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		r.ParseForm()
		log.Println(r.Form)

		c, err := DB.Get(&ccase.Case{
			Group:    r.FormValue("group"),
			SubGroup: r.FormValue("sgroup"),
			Feature:  r.FormValue("feature"),
			Name:     r.FormValue("case"),
		})

		task := c.GetTask(r.FormValue("task"))

		if task == nil {
			io.WriteString(w, "Cannot find task: "+r.FormValue("task"))
			return
		}

		t.Funcs(template.FuncMap{"ADD": ADD, "ENCAP": Encap})

		err = t.Execute(w, struct {
			Title string
			Case  *ccase.Case
			DB    *ccase.CaseDBInMem
			Task  *ccase.Task
		}{
			Title: "Dump Case Task",
			Case:  c,
			DB:    DB,
			Task:  task,
		})
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		newTask, err := ccase.CreateNewTask(r.Form)
		if err != nil {
			log.Printf("%q, %q", newTask, err)
			io.WriteString(w, err.Error())
			return
		}

		if c.IsTaskExist(newTask) {
			c.DelTask(newTask)
		}

		if err := c.AddTask(newTask); err != nil {
			io.WriteString(w, err.Error())
			return
		}

		DB.Save() //This should be more flexible
		t, err := template.ParseFiles("template/dumptask.html", "template/footer.html", "template/header.html", "template/taskheader.html", "template/casenavigator.html", "template/dumpcondition.html", "template/dumproutine.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		t.Funcs(template.FuncMap{"ADD": ADD, "ENCAP": Encap})
		err = t.Execute(w, struct {
			Title string
			Case  *ccase.Case
			DB    *ccase.CaseDBInMem
			Task  *ccase.Task
		}{
			Title: "Dump Case Task",
			Case:  c,
			DB:    DB,
			Task:  newTask,
		})
		if err != nil {
			log.Println(err.Error())
		}

	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func NewNewTask(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	err := r.ParseForm()
	if err != nil {
		log.Println("Cannot parse form: ", err.Error())
		return
	}

	log.Println(r.Form)

	c, err := DB.Get(&ccase.Case{
		Group:    r.FormValue("group"),
		SubGroup: r.FormValue("sgroup"),
		Feature:  r.FormValue("feature"),
		Name:     r.FormValue("name"),
	})
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/newnewtask.html", "template/footer.html", "template/header.html", "template/newroutine.html", "template/condition.html", "template/casenavigator.html", "template/taskheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, &struct {
			Title string
			Case  *ccase.Case
			DB    *ccase.CaseDBInMem
		}{
			Title: c.Name,
			Case:  c,
			DB:    DB,
		})
		if err != nil {
			log.Println(err.Error())
		}

	} else if r.Method == "POST" {

		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		newTask, err := ccase.CreateNewTask(r.Form)
		if err != nil {
			log.Printf("%q, %q", newTask, err)
			io.WriteString(w, err.Error())
			return
		}

		if err := c.AddTask(newTask); err != nil {
			io.WriteString(w, err.Error())
			return
		}

		DB.Save()

		if _, ok := r.Form["continue"]; ok {
			t, err := template.ParseFiles("template/newnewtask.html", "template/footer.html", "template/header.html", "template/newroutine.html", "template/condition.html", "template/casenavigator.html", "template/taskheader.html")
			if err != nil {
				log.Println(err)
				io.WriteString(w, err.Error())
				return
			}

			err = t.Execute(w, &struct {
				Title string
				Case  *ccase.Case
				DB    *ccase.CaseDBInMem
			}{
				Title: c.Name,
				Case:  c,
				DB:    DB,
			})
			if err != nil {
				log.Println(err.Error())
			}
		} else {
			t, err := template.ParseFiles("template/dumpcase.html", "template/footer.html", "template/header.html", "template/caseheader.html", "template/casenavigator.html", "template/taskheader.html")
			if err != nil {
				log.Println(err)
				io.WriteString(w, err.Error())
				return
			}

			err = t.Execute(w, &struct {
				Title string
				Case  *ccase.Case
				DB    *ccase.CaseDBInMem
			}{
				Title: c.Name,
				Case:  c,
				DB:    DB,
			})
			if err != nil {
				log.Println(err.Error())
			}
		}

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

func ProductInfo(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/productinfo.html", "template/footer.html", "template/header.html")
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

func RunScript(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/runscript.html", "template/footer.html", "template/header.html")
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

		script, err := ccase.RunUserScript(r.Form)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		t, err := template.ParseFiles("template/dumpscript.html", "template/footer.html", "template/header.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title    string
			Commands []*ccase.Command
		}{
			Title:    "Dump Script",
			Commands: script.Commands,
		})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println(r.Form)
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func DumpDeviceCurrentStatus(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/devicecurrentstatus.html", "template/footer.html", "template/header.html")
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

func DeviceTestCases(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/devicetestcases.html", "template/footer.html", "template/header.html", "template/casenavigator.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title string
			DB    *ccase.CaseDBInMem
		}{
			Title: "Device test cases",
			DB:    DB,
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

func BootstrapLayout(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/bootstraplayout.html", "template/footer.html", "template/header.html")
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

func Dashboard(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/dashboard.html", "template/footer.html", "template/header.html")
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

func SideBar(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/sidebar.html", "template/footer.html", "template/header.html")
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

func Product(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/product.html", "template/footer.html", "template/header.html")
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
		t, err := template.ParseFiles("template/devicetestcases.html", "template/footer.html", "template/header.html", "template/casenavigator.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, struct {
			Title string
			DB    *ccase.CaseDBInMem
		}{
			Title: "Device test cases",
			DB:    DB,
		})
		if err != nil {
			log.Println(err.Error())
		}

	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/product.html", "template/footer.html", "template/header.html")
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
		t, err := template.ParseFiles("template/product.html", "template/footer.html", "template/header.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}

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
	http.HandleFunc("/bootcss", BootCSS)
	http.HandleFunc("/pagefooter", PageFoorter)
	http.HandleFunc("/registernewcase", RegisterNewCase)
	http.HandleFunc("/precondition", PreCondition)
	http.HandleFunc("/postcondition", PostCondition)
	http.HandleFunc("/taskroutine", TaskRoutine)
	http.HandleFunc("/stepforward", StepForward)
	http.HandleFunc("/newnewtask", NewNewTask)
	http.HandleFunc("/prepostroutine", PrePostRoutine)
	http.HandleFunc("/product", Product)
	http.HandleFunc("/productinfo", ProductInfo)
	http.HandleFunc("/allcases", DeviceTestCases)
	http.HandleFunc("/devicestatus", DumpDeviceCurrentStatus)
	http.HandleFunc("/runscript", RunScript)
	http.HandleFunc("/bootstraplayout", BootstrapLayout)
	http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/sidebar", SideBar)
	http.HandleFunc("/dumpcase", DumpCase)
	http.HandleFunc("/deletetask", DeleteTask)
	http.HandleFunc("/dumptask", DumpTask)
	http.HandleFunc("/edittask", EditTask)
	http.HandleFunc("/login", Login)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}

var DB *ccase.CaseDBInMem

func init() {
	DB = &ccase.CaseDBInMem{
		Device: "V8500",
		Groups: make(map[string]*ccase.Group, 1),
	}

	value, err := ioutil.ReadFile("testcases.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(value, DB)
}
