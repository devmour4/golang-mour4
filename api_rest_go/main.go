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
var clients []Person

func GetClients(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json") // Define o tipo de resposta como JSON
    json.NewEncoder(w).Encode(clients) // Converte a lista de pessoas para JSON e envia na resposta
}
func GetClient(w http.ResponseWriter, r *http.Request)

func main() {
	router := mux.NewRouter()

	//o router é a função que vai usar a porta, o handlefunc é pra acessar o local que vou direcionar
	router.HandleFunc("/contact/", GetClients).Methods("GET") // Exibir todas as pessoas
	router.HandleFunc("/contact/{id}", GetClient).Methods("GET") // Localizar pessoa por ID
	router.HandleFunc("/contact/",) // Atualizar pessoa utiliza ID
	router.HandleFunc("/contact/",) // Deletar pessoa utiliza ID
	http.ListenAndServe(":8000", router)
}