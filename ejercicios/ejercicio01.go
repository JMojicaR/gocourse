package ejercicios

import (
	"strconv"
)

func RetornoValores(numero string) (int, string) {
	numeroInt, err := strconv.Atoi(numero)
	if err != nil {
		return 0, "Error " + err.Error()
	}
	if numeroInt >= 100 {
		return numeroInt, "El numero es mayor o igual a 100"
	} else {
		return numeroInt, "El numero es menor a 100"
	}
}
