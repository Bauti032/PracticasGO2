package main

// La diferencia entre variavle comun y vector es la siguiente

// Vector var vec[5] int
// Nombre CE Tipo de datos

// func ejemplo(vector*[5]int)
// Se tienen 2 vectores A y B de 20 elementos. Programar un procedimiento que permita ordenar
// vectores de forma ascendente. El procedimiento debe recibir como parámetro por referencia el
// vector a ordenar.
// Una vez ordenado los 2 vectores mediante la llamada al procedimiento, generar un nuevo vector C
// que sea la intercalación ordenada de A y de B (considere que no hay elementos repetidos).
import (
	"fmt"
)

func main(){
	var (
		vecA[5] int
		vecB[5] int
		vecC[10] int
	)


	cargarVectorA(&vecA)
	cargarVectorB(&vecB)
	ordenarAscendente(&vecA)
	ordenarAscendente(&vecB)
	vecC = cargarvecC(vecA, vecB)
	fmt.Print(vecC)
}


func cargarVectorA(vecA*[5] int){
	var num int
	for i := 0; i<5; i++{
		fmt.Print("Ingrese un numero para el vector A: ")
		fmt.Scanln(&num)
		vecA [i] = num
	}
}
 
func cargarVectorB(vecB*[5] int){
	var num int
	for i := 0; i<5; i++{
		fmt.Print("Ingrese un numero para el vector B: ")
		fmt.Scanln(&num)
		vecB [i] = num
	}
}

func cargarvecC(vecA [5]int, vecB [5]int) [10]int{
	var( ia, ib int
		vecC [10]int)


		for ic := 0; ic < 10; ic++{
			if ia < 5 && ib <5 && vecA[ia] < vecB[ib]{
					vecC[ic] = vecA[ia]
					ia++
			}else if ib < 5{
					vecC[ic] = vecB[ib]
					ib++
			}else if ia < 5{
					vecC[ic] = vecA[ia]
					ia++
			}
	}

	return vecC
}

func ordenarAscendente(vec *[5]int){

	for x := 0; x < 5; x++{
		for y := 0; y <5; y++{
			if vec[x]<vec[y]{
				aux := vec[x]
				vec[x] = vec[y]
				vec[y] = aux
			}
		}
	}
}





