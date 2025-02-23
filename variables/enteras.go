package variables

import "fmt"

func MuestraEnteras() {
	var intcomun int     //variable por declaracion
	intde32 := int32(10) //variable por asignacion
	intde64 := int64(100)
	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
