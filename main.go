package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Student struct {
	Name string
	Age  int64
}

func main() {
	fmt.Println("AB")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		students := map[string][]Student{
			"Students": {
				{Name: "Rama", Age: 20},
				{Name: "Sita", Age: 19},
			},
		}
		tmpl.Execute(w, students)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		name := r.PostFormValue("name")
		age := r.PostFormValue("age")
		fmt.Println(name, age)
		tmpl := template.Must(template.ParseFiles("index.html"))
		intAge, _ := strconv.ParseInt(age, 10, 64)
		tmpl.ExecuteTemplate(w, "student-list-element", Student{Name: name, Age: intAge})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-student", h2)

	err := http.ListenAndServe(":5172", nil)
	if err != nil {
		log.Fatal("test", err)
	}
}
