package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese el primer numero: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error al ingresar el numero " + err.Error())
		}
	}
	fmt.Println("Ingrese el segundo numero: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error al ingresar el numero " + err.Error())
		}
	}
	fmt.Println("Ingrese el leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
		fmt.Println(leyenda, numero1*numero2)
	}
}
