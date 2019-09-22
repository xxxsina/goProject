package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Form data:", r.Form)
	fmt.Println("PostForm data:", r.PostForm)
	fmt.Println("Url.Path :", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("URL", r.URL)
	fmt.Println("r :", r)
	for k, v := range r.Form {
		fmt.Println("k :", k, " v :", v)
	}
	fmt.Fprintf(w, "Hello astx")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
		return
	}
	fmt.Print("xxxooo")
}
