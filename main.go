package main

import (
	"fmt"
	"html/template"
	"net/http"

)

func main() {
	port:=":9999"
  err:=	http.ListenAndServe(port, nil)
	if err != nil{
		log.Fatal("ListenAndServe", err)
	}
	http.HandleFunc("/", mainPage)

}

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`

}

func mainPage(w http.ResponseWriter, r *http.Request)  {
	 user := User{"Pavlo" , "Just"}
	// js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("GitHub/go-tour/go-tour/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, nil); err !=nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
