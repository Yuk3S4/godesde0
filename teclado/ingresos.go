package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin) // inicializar scanner

	fmt.Println("Ingrese número 1 : ")
	if scanner.Scan() { // ? si el usuario ingresó algo por teclado
		num1, err = strconv.Atoi(scanner.Text()) // scanner.Text() - todo lo que el usuario ingresó
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2 : ")
	if scanner.Scan() { // ? si el usuario ingresó algo por teclado
		num2, err = strconv.Atoi(scanner.Text()) // scanner.Text() - todo lo que el usuario ingresó
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() { // ? si el usuario ingresó algo por teclado
		leyenda = scanner.Text() // scanner.Text() - todo lo que el usuario ingresó
	}

	fmt.Println(leyenda, num1*num2)

}
