package main

import "fmt"


// 12) Se tiene las calificaciones de 6 asignaturas (enumeradas del 1 al 6) correspondientes a 20 alumnos
// (enumerados del 1 al 20):
// a) Cargar una matriz AL con los datos de las calificaciones (cada fila representa un alumno y cada columna
// una materia).
// b) Crear una función que calcule y devuelva como resultado un vector con el promedio de cada asignatura.
// c) Crear una función que calcule y devuelva como resultado un vector con el promedio de cada estudiante.
// d) Generar una función que calcule y devuelva como resultado la cantidad de materias que aprobó y
// desaprobó un estudiante “x” pasado como parámetro.

const(
	fi = 3
	co = 4
)

func main(){

var(
	AL [fi][co]int
	prom [co]int
	promA [fi]int
	ap int
	des int
)

cargarMatriz(&AL)
imprimirM(AL)
prom = promM(AL)
fmt.Print("El promedio de cada materia es: ", prom)
promA = promAl(AL)
fmt.Print("El promedio de cada alumno es: ", promA)
materiasA(AL)
fmt.Printf("Cantidad de materias aprobadas: %d: \n", ap)
fmt.Printf("Cantidad de materias desaprobados: %d: ", des)







}

func cargarMatriz(matriz *[fi][co]int){

	for f := 0; f< fi; f++{
		for c := 0; c<co; c++{
			fmt.Printf("Ingrese la nota del alumno numero: %d, de la materia numero: %d: ", f+1, c+1)
			fmt.Scanln(&matriz[f][c])
		}
	}


}

func imprimirM(matriz [fi][co]int){

	for f := 0; f< fi; f++{
		for c := 0; c<co; c++{
			fmt.Printf("d\t", matriz[f][c] )
			fmt.Println("")
		}
	}
}

func promM(matriz [fi][co]int)(prom [co]int){
	var(
		ac int
	)


	for c := 0; c< co; c++{
		for f := 0; f<fi; f++{
			ac =+ matriz[f][c]
		}
		prom[c] = ac/fi
	}

	return 
}

func promAl(matriz [fi][co]int)(prom [fi]int){
	var(
		ac int
	)


	for f := 0; f< fi; f++{
		for c := 0; c<co; c++{
			ac =+ matriz[f][c]
		}
		prom[f] = ac/co
	}

	return 
}

func materiasA(matriz [fi][co]int)(ap int,des int){
var(
	fila int
)

	fmt.Print("Ingrese el codigo de alumno: ")
	fmt.Scan(&fila)
	fila = fila - 1
	

	for c:=0; c<co; c++{
		if matriz[fila][c] >= 6{
			ap++
		}else{
			des++
		}
	}

	return
}