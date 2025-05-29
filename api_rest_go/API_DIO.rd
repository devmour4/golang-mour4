# 📇 API REST em Go - Gerenciamento de Contatos

Este projeto é uma API REST simples feita em Go (Golang), que permite cadastrar, listar, buscar
 e remover contatos (pessoas). O backend usa o framework **Gorilla Mux** para gerenciamento de rotas HTTP.

---

## 🚀 Como executar

### Pré-requisitos:
- Go instalado ([Download aqui](https://go.dev/dl/))

### Instalação das dependências:

```bash
go get -u github.com/gorilla/mux

Executar o servidor:
go run main.go

A API estará disponível em: http://localhost:8000

//-----COMANDOS DE TERMINAL-----//

✅ Listar todos os clientes (GET /contact)
curl -X GET http://localhost:8000/contact

✅ Buscar um cliente específico pelo ID (GET /contact/{id})
curl -X GET http://localhost:8000/contact/1

✅ Criar um novo cliente (POST /contact)
curl -X POST http://localhost:8000/contact

✅ Remover um cliente pelo ID (DELETE /contact/{id})
curl -X DELETE http://localhost:8000/contact/1