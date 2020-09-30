package main

import "fmt"

type NameVisitor struct {
	ProductList string
}

func (n *NameVisitor) Visit(p IProductRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", n.ProductList, p.GetName())
}
