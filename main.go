package main

import (
	"fmt"
	"runtime"

	"github.com/Yuk3S4/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	// os := runtime.GOOS - Obtener sistema operativo actual
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}
}
