package main

type IVisitor interface {
	Visit(IProductRetriever)
}
