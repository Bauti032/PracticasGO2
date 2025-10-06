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
)

const (

	Columnas = 4
)

var Clibros int 
func main() {
	var (
	MatrizA [][Columnas]int	
	)


		matriz := cargamatriz(MatrizA)
		fmt.Println(matriz)
		var codautor int
		fmt.Println("Ingrese el codigo de autor: ")
		fmt.Scanln(&codautor)
		vector := codigoautor(matriz, codautor)
		fmt.Println(vector)
		CLM,CAM := CLMC(matriz)
		fmt.Printf("El codigo del libro mayor %d, codigo autor %d", CLM,CAM)
		NMatriz:= NMatriz(matriz)
		fmt.Println("El la matriz con los que son mas de 45: ")
		fmt.Print(NMatriz)
}

func cargamatriz(MatrizA [][Columnas]int) [][Columnas]int {
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

// func codigoautor(MatrizA [][Columnas]int, codautor int) []int {


// 	var vector []int
// 	for f := 0; f < Clibros; f++ {
// 		if codautor == MatrizA[f][1] {
// 			vector = append(vector, MatrizA[f][1])
// 		}

// 	} 

// 	return vector
// }

func CLMC(matriz [][Columnas]int)  (int, int){
	var (PLM int
		CAM int
		CLM int)

for i:=0; i< Clibros; i++{
	if i == 0{
		PLM = matriz[i][3]
		CAM = matriz[i][1]
		CLM = matriz[i][0]
	}else if PLM < matriz[i][3]{
		PLM = matriz[i][3]
		CAM = matriz[i][1]
		CLM = matriz[i][0]
	}
}


return  CLM, CAM
}

//c) Generar una función que devuelva una matriz quecontenga el  código de libro y el código del autor de
//aquellos libros cuyos números de ejemplares sea mayor a 45 unidades.
func  NMatriz(matriz [][Columnas]int) [][2]int{
	var nMatriz [][2]int

	for i:=0; i<Clibros; i++{
		if matriz [i][2] > 45{
				nMatriz = append(nMatriz, [2]int{matriz[i][0],matriz[i][1]})
		}

	} 
 return nMatriz
			//vector = append(vector, MatrizA[f][1])

}


func ordenarautor(nMatriz [][2]int) [][2]int {

	var vec [2]int
	var CMayor int 

	for i := 0; i < Clibros; i++ {
		if nMatriz[i][1] > CMayor {
			vec[0] = nMatriz[i][0]
			vec[1] = nMatriz[i][1]

			
		}

	}


}