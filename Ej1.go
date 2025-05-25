package ejercicios

import "fmt"

// Leer 3 numeros enteros, la diferencia del primero respecto al segundo, el producto de los dos ultimos, la division entre el primero y el tercero, y la suma de los tres numeros. Mostrar los resultados en pantalla.
func main_ej1() {
	var (
		num1 float64
		num2 float64
		num3 float64
	)

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanln(&num2)
	fmt.Print("Ingrese el tercer numero:")
	fmt.Scanln(&num3)

	dif := num1 - num2
	prod := num2 * num3
	div := num1 / num3

	fmt.Printf("Su diferencia es %.2f, su poducto es %.2f, su division es %.2f.", dif, prod, div)

}
