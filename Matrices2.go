package main

import "fmt"

	// 5) Cargar una matriz de 6x6 elementos y poner un 0 en donde encuentre un valor par.

	const f = 2
	const c = 2

func main(){
	var(
		M[f][c] int
		idM int 
	)	


	cargarM(&M,&idM )
	imprimirM(M,&idM)
	buscarC(&M, &idM)
	imprimirM(M, &idM)

}

func buscarC(m *[f][c]int, idm *int){
	
	for x :=0; x < f; x++{
		for y :=0; y < c; y++{
			if (m[x][y] % 2) == 0{
				m[x][y] = 0
			}else{

			}
		}
	}


}

func cargarM(m *[f][c]int, idm *int){
	var n int

	for x := 0 ; x < f; x++{
		for y :=0; y < c; y++{
			fmt.Print("ingrese numero: ")
			fmt.Scanln(&n)
			
			m[x][y] = n
		}
	}
	
	idm = 0
}

func imprimirM(m [f][c]int, idm *int){
	if idm == 0{
		fmt.Println("Funcion ingresada: ")
	}else{
		fmt.Println("Funcion modificada: ")		
	}


	for x := 0; x < f; x++ {
		for y := 0; y < c; y++ {
			fmt.Printf("%d\t", m[x][y])
		}
		fmt.Println("")
	} 

	fmt.Println("")

}