package main

import (
	"html/template"
	"net/http"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html")) //fazendo a alteração com o must
		err := t.Execute(os.Stdout, Cursos{
			{"GO", 40},
			{"Java", 30},
			{"Python", 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)

}
