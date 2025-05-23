package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)
type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`

}
type Address struct {
	State string `json:"state,omitempty"`
	City string `json:"city,omitempty"`
}

// Lista de pessoas (simulação de banco de dados)
var people []Person

func GetPeople (w http.ResponseWriter, r *http.Request)

func main() {

}