package main

import "fmt"

type CreateProductA struct {
}
type CreateProductB struct {
}

// 定义了工厂的通用接口
type CreateProduct interface {
	Use() string
}

func (c *CreateProductA) Use() string {
	// Do something that A need to do
	return "product A"
}
func (c *CreateProductB) Use() string {
	// Do something that B need to do
	return "product B"
}

type Factory struct {
}

func (f *Factory) CreateProduct(arg string) CreateProduct {
	switch arg {
	case "A":
		return &CreateProductA{}
	case "B":
		return &CreateProductB{}
	default:
		return &CreateProductA{}
	}
}

func main() {
	f := Factory{}
	A := f.CreateProduct("A")
	B := f.CreateProduct("B")
	fmt.Println(A.Use())
	fmt.Println(B.Use())
}
