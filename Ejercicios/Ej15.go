package main
import (
	"fmt"
)
var (
	Edad int
	Nombre string
	NMenor string
	Nmayor string
	EMayor int
	EMenor int
	Prom_Edad int
	Cant_Menor int
	Cant_Entre int
	Acum_Edad int

)
// 10 alumnos. Nombre de menor y mayor, cantidad de <12, cantidad <10<16, promedio edades


func main() {
	for x := 0; x < 10; x ++ {
		fmt.Println("Ingrese el nombre del alumno: ")
		fmt.Scanln(&Nombre)
		fmt.Println("Ingrese La edad del alumno: ")
		fmt.Scanln(&Edad)

		Acum_Edad += Edad

		if x == 1 {
			EMayor = Edad
			EMenor = Edad
			Nmayor = Nombre
			NMenor = Nombre
		}

		if EMayor < Edad {
			EMayor = Edad
			Nmayor = Nombre
		}

		if EMenor > Edad {
			EMenor = Edad
			NMenor = Nombre
		}

		if Edad < 12 {
			Cant_Menor += 1
		}
		if Edad > 10 && Edad < 16 {
			Cant_Entre += 1
		}
	}
	Prom_Edad = Acum_Edad / 10
	fmt.Println("El nombre del alumno menor es: ", NMenor)
	fmt.Println("El nombre del alumno mayor es: ", Nmayor)
	fmt.Println("La cantidad de alumnos menores a 12 es: ", Cant_Menor)
	fmt.Println("La cantidad de alumnos entre 10 y 16 es: ", Cant_Entre)
	fmt.Println("El promedio de las edades es: ", Prom_Edad)
}