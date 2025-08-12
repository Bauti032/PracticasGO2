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
	cargarvecC(&vecA, &vecB, &vecC)
	ordenarAscendenteC(&vecC)
	fmt.Print(vecC)
}


func cargarVectorA(vecA*[5] int){
	var num int
	for i := 0; i<5; i++{
		fmt.Print("Ingrese un numero para el vector: ")
		fmt.Scanln(&num)
		vecA [i] = num
	}
}
 
func cargarVectorB(vecB*[5] int){
	var num int
	for i := 0; i<5; i++{
		fmt.Print("Ingrese un numero para el vector: ")
		fmt.Scanln(&num)
		vecB [i] = num
	}
}

func cargarvecC(vecA*[5] int, vecB*[5] int, vecC*[10] int){
	for x:= 0; x<10; x++{
		vecC[x] = vecA[x]
		x++		
		vecC[x] = vecB[x]
	}
}

func ordenarAscendenteC(vecC*[10] int){
	for x := 0; x < 10; x++{
		for y := 0; y <10; y++{
			if vecC[x]<vecC[y]{
				aux := vecC[x]
				vecC[x] = vecC[y]
				vecC[y] = aux
			}
		}
	}
}

