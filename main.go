package main

//import "fmt"
import (
	"fmt"

	"github.com/goCourse/variables"

	"runtime"

	"github.com/goCourse/ejercicios"

	"github.com/goCourse/teclado"
)

func main() {
	variables.MuestraEnteras()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(2345)
	fmt.Println("Estado = ", estado)
	fmt.Println("Texto = ", texto)

	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Linux")
	} else if os == "Windows" {
		fmt.Println("Windows")
	} else {
		fmt.Println("Sistema Operativo desconocido", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "Windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Sistema Operativo desconocido", os)
	}
	numero, mensaje := ejercicios.RetornoValores("100")
	fmt.Println("Numero = ", numero)
	fmt.Println("Mensaje = ", mensaje)
	numero2, mensaje2 := ejercicios.RetornoValores("40")
	fmt.Println("Numero2 = ", numero2)
	fmt.Println("Mensaje2 = ", mensaje2)
	teclado.IngresoNumeros()
}
