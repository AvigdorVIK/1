package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
	fmt.Fprintf(page, "User name is:"+bob.name)

}

func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contact/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
