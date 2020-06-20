package tp4_calculdora


import (
	"fmt"
	"github.com/dbobadillabarcelo/TP4/operaciones"
)



func Calculadora()  {
	var a, b float64
	var operacion string
	var operandos []float64

	fmt.Printf("La calculadora está desviada y sumará 5 a cada resultado \n")
	fmt.Printf("Ingrese la operación a realizar: \n")
	fmt.Printf("Presione 1 para sumar \n")
	fmt.Printf("Presione 2 para restar \n")
	fmt.Printf("Presione 3 para multiplicar\n")
	fmt.Printf("Presione 4 para dividir \n")
	fmt.Scan(&operacion)
	fmt.Printf("Por favor, ingrese un numero: \n")
	fmt.Scan(&a)
	fmt.Printf("Ingrese otro numero: \n")
	fmt.Scan(&b)
	operandos = []float64{a,b}

	switch operacion {
	case "1":
		fmt.Print(operaciones.Sumar(operandos...))

	case "2":
		fmt.Print(operaciones.Restar(operandos...))
	case "3":
		fmt.Print(operaciones.Multiplicar(operandos...))
	case "4":
		fmt.Print(operaciones.Division(operandos...))
	default:
		fmt.Print("Operacion inválida.")
	}
}
