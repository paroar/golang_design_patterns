package main

type IProductRetriever interface {
	GetPrice() float32
	GetName() string
}
