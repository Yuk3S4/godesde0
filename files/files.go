package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Yuk3S4/godesde0/ejercicios"
)

var filePath string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()
	archivo, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err.Error())
		return
	}
	// Grabar en un archivo
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()
	if !Append(filePath, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filePath, texto string) bool {
	// Abrir un archivo en modo escritura y append
	archivo, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append ", err.Error())
		return false
	}

	// Grabar texto en el archivo
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString ", err.Error())
		return false
	}

	archivo.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error durante Open: ", err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> ", registro)
	}

	archivo.Close()
}
