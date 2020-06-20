package operaciones

const c float64 = 5

type Historial struct {
	Operandos []float64
	Operacion string
	Resultado float64
}

func Sumar(operandos ...float64) (Historial) {
	resultado :=operandos[0] + operandos[1] + c
	return  Historial{
		Operandos: operandos,
		Operacion: "suma",
		Resultado: resultado,
	}
}

func Restar(operandos ...float64) (Historial) {
	resultado := operandos[0] - operandos[1] + c
	return  Historial{

		Operandos: operandos,
		Operacion: "resta",
		Resultado: resultado,
	}
}

func Multiplicar(operandos ...float64) (Historial) {
	resultado := operandos[0] * operandos[1] + c

	return  Historial{

		Operandos: operandos,
		Operacion: "multiplicacion",
		Resultado: resultado,
	}
}

func Division(operandos ...float64) (Historial) {

	resultado := operandos[0] / operandos[1] + c

	return  Historial{

		Operandos: operandos,
		Operacion: "Division",
		Resultado: resultado,
	}
}