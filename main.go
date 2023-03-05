package main

import (
	"fmt"
)

func main() {
	// Preparamos un SimpleCoffee

	emptyCup := &Cup{}
	coffee := &Coffe{beberage: emptyCup}
	PrettyPrint(coffee)

	// Ahora al cafe anterior le agregamos leche de almendras
	coffeeWithMilk := &VeganMilk{beberage: coffee}
	PrettyPrint(coffeeWithMilk)

	// Y obviamente no nos olvidemos del azucar!
	coffeWithSugar := &Sugar{beberage: coffeeWithMilk}
	PrettyPrint(coffeWithSugar)

	// Mejor preparemos un te.
	tea := &Tea{beberage: emptyCup}
	PrettyPrint(tea)

	// Agregamos leche
	teaWithMilk := &Milk{beberage: tea}
	PrettyPrint(teaWithMilk)

	// Ahh pero soy vegano y lo quiero con azucar de mascabo
	teaWithVeganMilk := &VeganMilk{beberage: tea}
	teaWithVeganMilkAndBrownSugar := &BrownSugar{beberage: teaWithVeganMilk}
	PrettyPrint(teaWithVeganMilkAndBrownSugar)
}

func PrettyPrint(beberage Beberage) {

	fmt.Printf("** %v - Price: %v ** \n", beberage.GetDescription(), beberage.GetCost())
}

// Esta es la interface que comparten todos los decoradores
type Beberage interface {
	GetCost() float64
	GetDescription() string
}
