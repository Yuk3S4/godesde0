package ejercicios

import "strconv"

func Retornar2Valores(valor string) (int, string) {

	if num, _ := strconv.Atoi(valor); num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
