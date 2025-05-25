package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data any) {
	tmp, err := template.New("test").Parse(text)
	check(err)
	err = tmp.Execute(os.Stdout, data)
	check(err)
}

func main() {
	fmt.Println("---------------")
	executeTemplate("Template start:\nAction: {{.}}\nTemplate end.\n", "ABC")
	fmt.Println("---------------")
	executeTemplate("start {{if .}} dot is true {{end}} finish\n", true)
	executeTemplate("start {{if .}} dot is true {{end}} finish\n", false)
	fmt.Println("---------------")
	executeTemplate("Before loop: {{.}}\n{{range .}}Inside loop: {{.}}\n{{end}}After loop: {{.}}\n", []string{"do", "re", "mi"})
	fmt.Println("---------------")
	executeTemplate("Name: {{.Name}}, Count: {{.Count}}\n", Part{Name: "Fuses", Count: 5})
	executeTemplate("Name: {{.Name}}, Count: {{.Count}}\n", Part{Name: "Cables", Count: 2})
	fmt.Println("---------------")
	executeTemplate("Name: {{.Name}}\n{{if .Active}}Rate: {{.Rate}}{{end}}\n", Subscriber{Name: "Joe McRae", Active: true, Rate: 0.05})
	executeTemplate("Name: {{.Name}}\n{{if .Active}}Rate: {{.Rate}}{{end}}\n", Subscriber{Name: "Lucas Howes", Active: false, Rate: 0.05})
}
