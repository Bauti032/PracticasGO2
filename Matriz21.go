package main

// 21) Cargar una matriz A de Nx4 elementos, donde cada fila contiene los datos correspondientes a un libro de una
// librería (la primera columna contiene el código del libro, la segunda el código del autor, la tercera el número de
// ejemplares y la cuarta el precio del libro). Se pide:
// a) Generar un procedimiento que reciba como parámetro un código de un autor e imprima el código de todos
// los libros que éste publicó. El código de autor que se debe pasar como parámetro debe ser leído antes de
// invocar al procedimiento.
// b) Generar una función que devuelva el código del libro y el código del autor del libro más caro. Imprimir los
// resultados.
// c) Generar una función que devuelva una matriz que contenga el código de libro y el código del autor de
// aquellos libros cuyos números de ejemplares sea mayor a 45 unidades.
// d) Ordenar la nueva matriz generado por código de autor.
// e) Imprimir la matriz generada.


import (
	"fmt"
	// "math/rand"
	// "time"
)

const (

	Columnas = 4
)

var Clibros int 
func main() {
	// rand.Seed(time.Now().UnixNano())

	var (
	MatrizA [][Columnas]int
	// opcion int
	
	)


	
	

	// fmt.Println("1. Carga matriz")
	// fmt.Println("2 Codigo autor, devuelve todos los libros")
	// fmt.Println("3. Devuelve mas caro")
	// fmt.Println("4. Ejemplares mayor a 45")
	// fmt.Println("6. Ordenar e imprimir la matriz anterior")
	// fmt.Println("Ingrese la opcion: ")
	// fmt.Scanln(&opcion)



		matriz := cargamatriz(MatrizA)
		fmt.Println(matriz)
		var codautor int
		fmt.Println("Ingrese el codigo de autor: ")
		fmt.Scanln(&codautor)
		vector := codigoautor(matriz, codautor)
		fmt.Println(vector)
		
// for opcion < 5 {
	// switch opcion {
	// case 1:
	// 	matriz := cargamatriz(MatrizA)
	// 	fmt.Println(matriz)
	// 	fmt.Println("Ingrese la opcion: ")
	// 	fmt.Scanln(&opcion)
	// case 2:
	// 	var codautor int
	// 	fmt.Println("Ingrese el codigo de autor: ")
	// 	fmt.Scanln(&codautor)
	// 	vector := codigoautor(matriz, codautor)
	// 	fmt.Println(vector)
	// 	fmt.Println("Ingrese la opcion: ")
	// 	fmt.Scanln(&opcion)

	// case 3:
	// case 4:
	// case 5: 
	// default :
	// 	return
	// 	}
	// }
// }
	}

func cargamatriz(MatrizA [][Columnas]int) [][Columnas]int {

	// var nlibro int
	// var nautor int
	// var nejemplares int
	// var precio int
	
	var fila [Columnas]int

	fmt.Println("Ingrese la cantidad de libros a ingresar: ")
	fmt.Scanln(&Clibros)
	

	for i := 0; i < Clibros; i++ {
		fmt.Println("Ingrese el codigo del libro: ")
		fmt.Scanln(&fila[0])
		fmt.Println("Ingrese el codigo del autor: ")
		fmt.Scanln(&fila[1])
		fmt.Println("Ingrese el numero de ejemplares: ")
		fmt.Scanln(&fila[2])
		fmt.Println("Ingrese el precio del libro: ")
		fmt.Scanln(&fila[3])
	

		MatrizA = append(MatrizA, fila)
	}
	
	return MatrizA 
}

func codigoautor(MatrizA [][Columnas]int, codautor int) []int {


	var vector []int
	for f := 0; f < Clibros; f++ {
		if codautor == MatrizA[f][1] {
			vector = append(vector, MatrizA[f][1])
		}

	} 

	return vector
}