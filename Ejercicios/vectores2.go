package main

import "fmt"

var (
	
	Nro_Movil int
	Rem[7] int
	Cant_Cuadras int
	n1 int
)

func main() {
	Menu()
 }

func cargaRem(Nro_Movil int) {

	if Nro_Movil > 0 && Nro_Movil <= 6 {
		fmt.Println("Ingrese la cantidad de cuadras: ")
		fmt.Scanln(&Cant_Cuadras)
		Rem[Nro_Movil] += Cant_Cuadras 
	} else {
		fmt.Println("Numero de movil fuera de rango vuelva a ingresarlo. ")
		return
	}
 
}

func calculo_ganancias(Nro_Movil int, Rem[7] int)  {

	Cuadras_Realizadas := Rem[Nro_Movil]
	Total := (Rem[Nro_Movil] * 500)

	fmt.Printf("El Nro.Movil %d, realizó %d cuadras. Con un total de ganancias de %d. \n", Nro_Movil, Cuadras_Realizadas, Total)

}

func Reporte_General(Rem[7] int) {

	for x := 1 ; x < 7; x++ {
		Cuadras_Realizadas := Rem[x]
		Total := (Rem[x] * 500)
		fmt.Printf("El Nro.Movil %d, realizó %d cuadras. Con un total de ganancias de %d. \n", x, Cuadras_Realizadas, Total)
	}
}

func OrdenDescendente(Rem *[7]int){

	for x := 0; x < 7; x++{
		for y := 0; y <7; y++{
			if Rem[x]>Rem[y]{
				aux := Rem[x]
				Rem[x] = Rem[y]
				Rem[y] = aux
			}
		}
	}

}

func Menu() {
	for {
		fmt.Println("Bienvenido al Sistmea. Elija alguna de las opciones.")
		fmt.Println("1. Carga de Moviles.")
		fmt.Println("2. Visualizar Ganancias de cada movil.")
		fmt.Println("3. Reporte Total.")
		fmt.Println("4. Ordenar Movil s/ganancia.")
		fmt.Println("0. Salir del programa.")

		fmt.Println("Ingrese alguna de las opciones: ")
		fmt.Scan(&n1)

		opcion := n1

		switch opcion {
		case 1:
			for {
			fmt.Println("Ingrese el Nro. de Móvil: ")
			fmt.Scanln(&Nro_Movil)
		
			if Nro_Movil == 0 {
				break
			}
			cargaRem(Nro_Movil)
			}
		case 2:
			fmt.Println("Ingrese el Nro. de Móvil: ")
			fmt.Scanln(&Nro_Movil)
			calculo_ganancias(Nro_Movil, Rem)
		case 3:
			Reporte_General(Rem)
		case 4:
			OrdenDescendente(&Rem)
			fmt.Println(Rem)
		case 0:
			return
		}
	}
		
}