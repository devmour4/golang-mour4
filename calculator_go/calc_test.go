package main

import (
	"testing"
)

// padrão triple A - AAA
// A - Arrange (Preparar)
// A - Act (Rodar)
// A - Assert (verificar as asserções)

/*padrão exemplificado abaixo*/



//****--Soma--****//
func TestShouldSumCorrect (t *testing.T) {
//arrange
	teste := Sum(3, 2, 1)
//Act
	resultado := 6
//Assert
	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}
func TestShouldSumIncorrect(t *testing.T) {
//Arrange
	teste := Sum(3, 2, 1)
//Act
	resultado := 7
//Assert
	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}

//****--Sub--****//

func TestShouldSubCorrect (t *testing.T) {
//arrange
	teste := Sub(10, 5)
	resultado := -5

	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}
func TestShouldSubIncorrect(t *testing.T) {
	teste := Sub(10, 5)

	resultado := 5

	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}

//****--Multiplication--****//
func TestShouldMultCorrect (t *testing.T) {
//arrange
	teste := Mult(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}
func TestShouldMultIncorrect(t *testing.T) {
	teste := Mult(10, 10)

	resultado := 101

	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}
//****--DIVISION--****/]
func TestShouldDivisionCorrect (t *testing.T) {
//arrange
	teste := Division(20)
	resultado := 20

	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}
func TestShouldDivisionIncorrect(t *testing.T) {
	teste := Division(20)

	resultado := 21

	if teste != resultado {
		t.Error("valor esperado: ", resultado, "valor obtido: ", teste)
	}
}