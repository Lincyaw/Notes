package main

import "fmt"

type Factory struct {
}

func (f *Factory) CreateProduct(arg string) string {
	switch arg {
	case "A":
		return CreateProductA()
	case "B":
		return CreateProductB()
	default:
		return "empty"
	}
}
func CreateProductA() string {
	return "product A"
}
func CreateProductB() string {
	return "product B"
}
func main() {
	f := Factory{}
	fmt.Println(f.CreateProduct("A"))
	fmt.Println(f.CreateProduct("B"))
	fmt.Println(f.CreateProduct("C"))
}
