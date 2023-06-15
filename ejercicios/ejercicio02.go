package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin) // inicializar scanner

	var numero int
	var err error
	var texto string

	for {
		fmt.Println("Ingrese un número : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 11; i++ {
		texto += fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i)
	}

	return texto
}
