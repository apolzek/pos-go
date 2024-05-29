package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html")) //fazendo a alteração com o must
	err := t.Execute(os.Stdout, Cursos{
		{"GO", 40},
		{"Java", 30},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}

}
