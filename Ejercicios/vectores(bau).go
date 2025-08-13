package main

import "fmt"

// Procedimiento para ordenar un vector de 20 enteros de forma ascendente
func ordenarAscendente(v *[20]int) {
    // Algoritmo de burbuja simple
    for i := 0; i < len(v)-1; i++ {
        for j := 0; j < len(v)-i-1; j++ {
            if v[j] > v[j+1] {
                v[j], v[j+1] = v[j+1], v[j] // intercambiar
            }
        }
    }
}

// Función para intercalar dos vectores ordenados en uno nuevo
func intercalar(A, B [20]int) [40]int {
    var C [40]int
    i, j, k := 0, 0, 0

    for i < 20 && j < 20 {
        if A[i] < B[j] {
            C[k] = A[i]
            i++
        } else {
            C[k] = B[j]
            j++
        }
        k++
    }

    // Copiar los elementos restantes si quedan
    for i < 20 {
        C[k] = A[i]
        i++
        k++
    }
    for j < 20 {
        C[k] = B[j]
        j++
        k++
    }

    return C
}

func main() {
    A := [20]int{5, 12, 1, 9, 7, 2, 15, 18, 3, 10, 6, 8, 11, 14, 16, 4, 13, 19, 17, 20}
    B := [20]int{21, 25, 22, 24, 23, 27, 26, 28, 30, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}

    // Ordenar A y B
    ordenarAscendente(&A)
    ordenarAscendente(&B)

    fmt.Println("Vector A ordenado:", A)
    fmt.Println("Vector B ordenado:", B)

    // Intercalar
    C := intercalar(A, B)
    fmt.Println("Vector C intercalado:", C)
}
