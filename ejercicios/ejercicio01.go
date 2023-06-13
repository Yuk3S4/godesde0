package ejercicios

import "strconv"

func Retornar2Valores(valor string) (int, string) {

	if numero, _ := strconv.Atoi(valor); numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
