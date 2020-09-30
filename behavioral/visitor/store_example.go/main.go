package main

import "fmt"

func main() {
	products := make([]IVisitable, 2)
	products[0] = &Rice{
		Product: Product{
			Price: 32.0,
			Name:  "Some rice",
		},
	}
	products[1] = &Pasta{
		Product: Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}

	//Print the sum of prices
	priceVisitor := &PriceVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)

	//Print the product list
	nameVisitor := &NameVisitor{}

	for _, n := range products {
		n.Accept(nameVisitor)
	}

	fmt.Printf("Product list: %s", nameVisitor.ProductList)
}
