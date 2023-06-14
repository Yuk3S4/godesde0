package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin) // inicializar scanner

	var numero int
	var err error

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
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}

}
