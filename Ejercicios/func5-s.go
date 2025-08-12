package main


import (
	"fmt"
	"math/rand"
	"Time"
)

func main() {

	var (
		number int
	)


}

func Generar_random() {
	rand.Seed(Time.Now().Unix())

	number = rand.Intn(101)
}

func checkNumeroAdivinado(random int, input int) (res int) {

	 if (input == random) {
		res = 0
	 } else if (input > random) {
		res = 1
	 } else {
		res = -1
	 }

	 return res

}


func checkIntentos (chec, intents int) {
	if (intents <= 5){
		fmt.Println("Sos bueno")
	} else if (intents > 15) {
		fmt.Println("Malo")
	} else {
		fmt.Println("Regular")
	}
}