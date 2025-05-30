package main

import (
	"encoding/json"
	"strconv"
	"net/http"
	"log"
	"github.com/gorilla/mux"
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

func GetClients(w http.ResponseWriter, r *http.Request) { //EXIBE todos os usuarios 
    w.Header().Set("Content-Type", "application/json") // Define o tipo de resposta como JSON
    json.NewEncoder(w).Encode(clients) // Converte a lista de pessoas para JSON e envia na resposta
}
func GetClient(w http.ResponseWriter, r *http.Request) { //EXIBE o usuario através do ID
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	
	for _, item := range clients {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func CreateClient (w http.ResponseWriter, r *http.Request) {
	var client Person
	err := json.NewDecoder(r.Body).Decode(&client) //pega os dados de requisição do body(json), transforma os JSON em uma pessoa (Person)
	if err != nil { //como eu adcionei a variavel err na linha de cima, e fiz o laço aqui, eu verifico se tem erros e informo ao usuário se houve ou não um problema
		http.Error(w, "Erro ao processar JSON: formato inválido", http.StatusBadRequest)
		return
	}
	//e armazena o objeto(a pessoa) convertido na variável client.

	//🔹 json.NewDecoder(r.Body) cria um decodificador JSON que lê r.Body.
	//🔹 .Decode(&person) transforma os dados JSON no objeto Person dentro do Go

	client.ID = "client_" + strconv.Itoa(len(clients)+1) //Gera um ID automaticamente para o cliente, baseado no tamanho da lista clients.
								 
	clients = append(clients, client)
	w.Header().Set("Content-Type", "application/json") //defino o formato da resposta
	json.NewEncoder(w).Encode(client)
}
func UpdateClient (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	var updateClient Person
	err := json.NewDecoder(r.Body).Decode(&updateClient)

	if err != nil {
		http.Error(w, "Erro ao processar o JSON: formato inválido", http.StatusBadRequest)
		return
	}
	for index, item := range clients {
		if item.ID == params["id"] {
			updateClient.ID = item.ID
			clients[index] = updateClient
			json.NewEncoder(w).Encode(updateClient)
		}

		http.Error(w, "Cliente não encontrado", http.StatusNotFound)

	}

}
func RemoveClient (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for index, item := range clients {
		if item.ID == params["id"] { //compara o id do usuario com o da URL e ve se é o mesmo para que remova
			clients = append(clients[:index], clients[index+1:]...) //Remove o client, com a função append utilizada em slices
			w.WriteHeader(http.StatusOK) //define o código de status HTTP na resposta. No seu código, ela indica que a remoção do cliente foi bem-sucedida.
			return
		}
		
	}
	http.Error(w, "Cliente não encontrado", http.StatusNotFound)
}


func main() {
	router := mux.NewRouter()

	//o router é a função que vai usar a porta, o handlefunc é pra acessar o local que vou direcionar
	router.HandleFunc("/contact", GetClients).Methods("GET") // Exibir todas as pessoas
	router.HandleFunc("/contact/{id}", GetClient).Methods("GET") // Localizar pessoa por ID
	router.HandleFunc("/contact", CreateClient).Methods("POST") // Criar pessoa
	router.HandleFunc("/contact/{id}", UpdateClient).Methods("PUT") // Atualiza a pessoa pelo ID
	router.HandleFunc("/contact/{id}",RemoveClient).Methods("DELETE") // Deletar pessoa utiliza ID
	

	clients = append(clients, Person{ID: "1", Firstname: "Abigail", Lastname: "Brithanny", Address: &Address{City: "São Paulo", State: "São Paulo"}})
	clients = append(clients, Person{ID: "2", Firstname: "Josivaldo", Lastname: "Leite", Address: &Address{City: "Nova Friburgo", State: "Rio de Janeiro"}})

	err := http.ListenAndServe(":8000", router)
	
	if err != nil {
    log.Fatal("Erro ao iniciar o servidor:", err)
	}

}