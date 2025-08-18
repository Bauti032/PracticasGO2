package main

import "fmt"

var (
	Notas     [50]int
	Promedio  int
	Nmayor    int
	Lmayor    int
	Nmenor    int
	Lmenor    int
	NroLegajo int
	Nota      int
	Opcion    int
)

func main() {
	for {
		if !menu() {
			break
		}
	}
}

func menu() bool {
	fmt.Println("\n1. Cargar Notas.")
	fmt.Println("2. Ver Notas Individuales")
	fmt.Println("3. Reporte General.")
	fmt.Println("4. Orden Descendente.")
	fmt.Println("5. Estadísticas.")
	fmt.Println("6. Salir")
	fmt.Print("Elija la opción a realizar: ")
	fmt.Scanln(&Opcion)

	switch Opcion {
	case 1:
		for {
			fmt.Print("Ingrese el Nro de Legajo (0 para volver al menú): ")
			fmt.Scanln(&NroLegajo)
			if NroLegajo == 0 {
				break
			}
			cargaVector(NroLegajo)
		}
	case 2:
		fmt.Print("Ingrese el Nro de legajo: ")
		fmt.Scanln(&NroLegajo)
		ReporteIndividual(NroLegajo, Notas)
	case 3:
		Reporte_General(Notas)
	case 4:
		OrdenDescendente(&Notas)
		fmt.Println("Notas ordenadas descendentemente:", Notas)
	case 5:
		Estadísticas(Notas)
	case 6:
		fmt.Println("Saliendo del programa...")
		return false
	default:
		fmt.Println("Opción no válida. Intente de nuevo.")
	}
	return true
}

func cargaVector(NroLegajo int) {
	if NroLegajo >= 1 && NroLegajo <= 50 {
		fmt.Print("Ingrese la nota del alumno (del 1 al 10): ")
		fmt.Scanln(&Nota)
		Notas[NroLegajo-1] = Nota
	} else {
		fmt.Println("Número de legajo fuera de rango. Vuelva a ingresar uno entre el 1 y 50.")
	}
}

func ReporteIndividual(Nro_Legajo int, Notas [50]int) {
	if Nro_Legajo < 1 || Nro_Legajo > 50 {
		fmt.Println("Número de legajo fuera de rango.")
		return
	}
	Nota = Notas[Nro_Legajo-1]
	if Nota >= 6 {
		fmt.Printf("El alumno con legajo Nro. %d aprobó con una Nota de: %d\n", Nro_Legajo, Nota)
	} else {
		fmt.Printf("El alumno con legajo Nro. %d desaprobó con una Nota de: %d\n", Nro_Legajo, Nota)
	}
}

func Reporte_General(Notas [50]int) {
	for i := 0; i < 50; i++ {
		fmt.Printf("El alumno con Nro. Legajo %d tiene una nota de %d.\n", i+1, Notas[i])
	}
}

func OrdenDescendente(Notas *[50]int) {
	for x := 0; x < 50; x++ {
		for y := x + 1; y < 50; y++ {
			if Notas[x] < Notas[y] {
				aux := Notas[x]
				Notas[x] = Notas[y]
				Notas[y] = aux
			}
		}
	}
}

func Estadísticas(Notas [50]int) {
	if len(Notas) == 0 {
		fmt.Println("No hay notas cargadas.")
		return
	}
	Nmayor = Notas[0]
	Lmayor = 1
	Nmenor = Notas[0]
	Lmenor = 1
	suma := 0

	for i := 0; i < 50; i++ {
		if Notas[i] > Nmayor {
			Nmayor = Notas[i]
			Lmayor = i + 1
		}
		if Notas[i] < Nmenor {
			Nmenor = Notas[i]
			Lmenor = i + 1
		}
		suma += Notas[i]
	}
}
