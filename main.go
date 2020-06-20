package main

import (
	"fmt"
	"github.com/dbobadillabarcelo/TP3"
	"github.com/dbobadillabarcelo/TP4"
	"github.com/dbobadillabarcelo/TP5"
	"github.com/dbobadillabarcelo/TP6"
	"github.com/dbobadillabarcelo/TP7"
	"github.com/dbobadillabarcelo/TP8"
)

func main(){
	var asd int64
	fmt.Printf("Ingrese el numero del TP que quiere ejecutar: \n")
	fmt.Printf("3): Introduccion a Go \n")
	fmt.Printf("4): Variables y estructuras en Go \n")
	fmt.Printf("5): Consumo de APIs en Go \n")
	fmt.Printf("6): Concurrencia y Paralelismos en Go \n")
	fmt.Printf("7): Testing \n")
	fmt.Printf("8): Creacion de una API \n\n ")
	fmt.Scan(&asd)
	switch asd {
	case 3:
		tp3_hello_world.HelloWorld()
		break
	case 4:
		tp4_calculdora.Calculadora()
		break
	case 5:
		tp5_consumo_apis.ConsumoApi()
		break
	case 6:
		tp6_calculadora_hilos.CalculadoraHilos()
		var input string
		fmt.Scanln(&input)
		break
	case 7:
		tp7_calculadora_test.CalculadoraTest()
		break
	case 8:
		tp_8_calculadoraApi.CalculadoraApi()
		break
	}
}
