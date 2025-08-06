package main
import "fmt"

func main() { 
	var (
		 Promotores int = 10 
		 CantE int = 0
		 CA int = 0
		 CB int = 0
		 CC int = 0
		 )
		 
		 for i := 0; i < Promotores; i++ {

			fmt.Println("Ingrese la cantidad de excursiones: ")
			fmt.Scanln(&CantE)
			
			CategoriaProm := categoriaPromotor(CantE, CA, CB, CC)
			ImportePagar := importeApagar(CategoriaProm, CantE)

			fmt.Printf("El promotor es de categoria %s y debe cobrar %.2f\n", CategoriaProm, ImportePagar)

			}
		
		}


func categoriaPromotor(cantE int, CA int, CB int, CC int) string {

	if cantE < 10 {
		CA++
		return "A"
	} else if cantE >= 10 && cantE < 100 {
		CB++
		return "B"
	} else {
		CC++
		return "C"
	}
}
func importeApagar(categoria string, cantE int) float64 {
		switch categoria {
		case "A":
			return float64(cantE)*1000.0
		case "B":
			return float64(cantE)*1500.0
		case "C":
			return float64(cantE)*1900.0
		}
		return 0
}




