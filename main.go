package main

import (
	"fmt"
	"LAB3TPS-master/TP3"
	"LAB3TPS-master/TP4"
	"LAB3TPS-master/TP5"
	"LAB3TPS-master/TP6"
	"LAB3TPS-master/TP7"
	"LAB3TPS-master/TP8"
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
