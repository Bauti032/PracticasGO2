package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Procedimiento para ordenar un vector de forma ascendente (por referencia)
func ordenarAscendente(vec *[20]int) {
	for i := 0; i < len(vec)-1; i++ {
		for j := i + 1; j < len(vec); j++ {
			if vec[i] > vec[j] {
				vec[i], vec[j] = vec[j], vec[i]
			}
		}
	}
}

// Función para intercalar dos vectores ordenados sin elementos repetidos
func intercalarOrdenado(a, b [20]int) []int {
	c := make([]int, 0, 40)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else if a[i] > b[j] {
			c = append(c, b[j])
			j++
		} else { // Si son iguales, solo uno y avanzar ambos
			c = append(c, a[i])
			i++
			j++
		}
	}
	// Agregar los elementos restantes de A
	for i < len(a) {
		c = append(c, a[i])
		i++
	}
	// Agregar los elementos restantes de B
	for j < len(b) {
		c = append(c, b[j])
		j++
	}
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var A, B [20]int

	// Llenar los vectores con números aleatorios entre 0 y 99
	for i := 0; i < 20; i++ {
		A[i] = rand.Intn(100)
		B[i] = rand.Intn(100)
	}

	fmt.Println("Vector A:", A)
	fmt.Println("Vector B:", B)

	// Ordenar los vectores por separado
	ordenarAscendente(&A)
	ordenarAscendente(&B)

	fmt.Println("Vector A ordenado:", A)
	fmt.Println("Vector B ordenado:", B)

	// Intercalar los vectores ordenados
	C := intercalarOrdenado(A, B)
	fmt.Println("Vector C (intercalado):", C)
}
