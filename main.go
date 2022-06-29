package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

func (u User) getAllInto() string {
	return fmt.Sprintf("User name is: %s. He is %d and has money "+
		" equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{Name: "Bob", Age: 25, Money: -50, Avg_grades: 4.2, Happiness: 0.8}
	//bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInto())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
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
