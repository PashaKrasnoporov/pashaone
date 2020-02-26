package main

import (
	_ "encoding/json"
	"html/template"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", mainPage)
	port:=":9999"
	println("Server listen on port:", port)
err :=http.ListenAndServe(port, nil)
if err != nil{
	log.Fatal("ListenAndServe",err)
}

}
type User struct {
FirstName string `json:"first_name"`
LastName string `json:"last_name"`
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	//user := User{"Паша","Краснопьоров"}
	//js, _ :=json.Marshal(user)

tmpl, err := template.ParseFiles("statik/index")
if err !=nil {
	http.Error(w, err.Error(), 400)
	return
}

if err := tmpl.Execute(w, nil); err != nil {
	http.Error(w, err.Error(), 400)
	return
}
//w.Write(js)
}
