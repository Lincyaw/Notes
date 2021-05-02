package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	u1 := User{
		Name:   "猪头",
		Gender: "女",
		Age:    20,
	}
	t.Execute(w, u1)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v\n", err)
		return
	}
}

