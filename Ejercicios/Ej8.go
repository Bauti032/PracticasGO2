package main

import "fmt"

func main() {
	var (
		hora string
		minuto string
		segundo string
	)

	fmt.Println("Ingrese la hora: ")
	fmt.Scanln(&hora)
	fmt.Println("Ingrese los minutos: ")
	fmt.Scanln(&minuto)
	fmt.Println("Ingrese los segundos: ")
	fmt.Scanln(&segundo)

	fmt.Printf("%s:%s:%s", hora, minuto, segundo)
}