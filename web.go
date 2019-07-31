package main

import "fmt"
import "html/template"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "john wick",
		"Message": "have a nice day",
	}

	var t, err = template.ParseFiles("view/show.html")
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
	http.HandleFunc("/", root)

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
