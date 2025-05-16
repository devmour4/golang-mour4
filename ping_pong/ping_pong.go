package main

import (
	"fmt"
	"time"
)
//criar canais com concorrencia para exibir com alternancia ping pong
//criar go routines com a função de exibir e com o valor dentro dos canais
func calculationResult (values chan string) {
	for i := 0; ; i++ {
		values <- "ping"
	}
}
func calculationResults (values chan string) {
	for i := 0; ; i++ {
		values <- "pong"
	}
}
func showResults(values chan string) {
	for {
		value := <- values
		fmt.Println(value)
		time.Sleep(time.Second * 1)
	}
}
func main() {

	var values chan string = make(chan string)
	go calculationResult(values)
	go showResults(values)
	go calculationResults(values)
	var userInput string
	fmt.Scanln(&userInput)

}