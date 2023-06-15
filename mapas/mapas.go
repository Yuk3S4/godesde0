package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	// fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	// fmt.Println(paises)
	// fmt.Println(paises["Mexico"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	// }

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("Puntaje %d, existe = %t \n", puntaje, existe)
}
