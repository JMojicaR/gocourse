package main

//import "fmt"
import (
	"fmt"

	"github.com/goCourse/variables"
)

func main() {
	variables.MuestraEnteras()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(2345)
	fmt.Println("Estado = ", estado)
	fmt.Println("Texto = ", texto)
}
