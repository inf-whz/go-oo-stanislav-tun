package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w,"index", nil)
}

func main(){
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3030", nil)
}
