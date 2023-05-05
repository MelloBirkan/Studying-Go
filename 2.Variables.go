package main
import "fmt"

func main() {
	// Forma completa
	var product string = "T-shirt"
	// Forma curta
	var cost = 20
	// Forma mais curta
	quantity := 5

	fmt.Println("Product's value is:", product)
	fmt.Printf("Cost's type is: %T\n", cost)

	quantity = 19
	fmt.Println("The quantity of products is:", quantity)
}