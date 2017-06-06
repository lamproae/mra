package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl = `{{$Liwei:="ttttt"}}{{$Liwei}}`

func main() {
	t, err := template.New("text").Parse(tpl)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(os.Stdout, nil)
}
