package main

import "fmt"

func main() {
	// Constants = name and type
	const pi float64 = 3.14
	const pi2 = 3.1425

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Variables integer
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)
	//fmt.Println("Hola Mundo")

	// Zero values - Default values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area calculator
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)
}
