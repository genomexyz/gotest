package main

import "fmt"
import "html/template"
import "net/http"

type M map[string]interface{}

//var tmpl, err = template.ParseGlob("view/*")

func index(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "john wick",
		"Message": "have a nice day",
	}

	var t, err = template.ParseFiles("view/halmuka.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(w)

	t.Execute(w, data)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "halo!")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", root)

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
