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

func (u User) getAllInto() string {
	return fmt.Sprintf("User name is: %s. He is %d and has money "+
		" equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInto())

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
