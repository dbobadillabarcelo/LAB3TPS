package tp6_calculadora_hilos

import (
	"fmt"
	"time"
)

type Base struct {
	Operandos []float64
	Operacion string
	Resultado float64
}

var chan1 chan Base = make(chan Base)
var chan2 chan Base = make(chan Base)
var chan3 chan Base = make(chan Base)
var chan4 chan Base = make(chan Base)

func CalculadoraHilos() {

	var num1, num2 float64
	var operandos []float64

	fmt.Printf("Ingrese dos operandos: \n")
	fmt.Scan(&num1, &num2)
	operandos = []float64{num1, num2}

	go func() {

		resultado := operandos[0] + operandos[1]

		chan1 <- Base{

			Operandos: operandos,
			Operacion: "suma",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		resultado := operandos[0] - operandos[1]

		chan2 <- Base{

			Operandos: operandos,
			Operacion: "resta",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)
	}()

	go func() {

		resultado := operandos[0] * operandos[1]

		chan3 <- Base{

			Operandos: operandos,
			Operacion: "multiplicacion",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {

		resultado := operandos[0] / operandos[1]

		chan4 <- Base{

			Operandos: operandos,
			Operacion: "Division",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		for {
			select {
			case suma := <-chan1:
				fmt.Println(suma)
			case resta := <-chan2:
				fmt.Println(resta)
			case multiplicar := <-chan3:
				fmt.Println(multiplicar)
			case dividir := <-chan4:
				fmt.Println(dividir)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
