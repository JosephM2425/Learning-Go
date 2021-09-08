package main

import "fmt"

func printHelloWorld(num int) {
	fmt.Println("Hola Mundo", num)
}

func twoArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobuleReturn(a int) (c, d int) {
	return a, a * 2
}

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

	helloMessage := "Hello"
	worldMessage := "World"

	// println

	fmt.Println(helloMessage, worldMessage)

	// printf

	name := "Joseph"
	cursos := 500

	// Con esta es de mejor practica si conocemos el tipo de dato
	fmt.Printf("%s tiene mas de %d cursos\n", name, cursos)

	// La v es para cuando no se sabe el tipo de dato
	fmt.Printf("%v tiene mas de %v cursos\n", name, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", name, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("Cursos: %T\n", cursos)

	// Funciones
	printHelloWorld(5)

	twoArguments(5, 6, "Hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := dobuleReturn(2)
	fmt.Println("Value1 y Value2: ", value1, value2)

	value3, _ := dobuleReturn(2)
	fmt.Println("Value3: ", value3)
}
