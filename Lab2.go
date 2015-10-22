package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"httprouter"
	"net/http"
)

var string1 string

// Req represents the structure of request
type Req struct {
	Name string `json:"name"`
}

// Resp represents the structure of request
type Resp struct {
	Greeting string `json:"greeting"`
}

// GreetUser creates a new user resource
func GreetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u := Resp{}
	u2 := Req{}
	json.NewDecoder(r.Body).Decode(&u2)

	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString(u2.Name)
	buffer.WriteString("!")
	u.Greeting = buffer.String()

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func main() {
	// Instantiate a new router
	r := httprouter.New()

	r.POST("/hello", GreetUser)

	http.ListenAndServe("localhost:8080", r)
}
