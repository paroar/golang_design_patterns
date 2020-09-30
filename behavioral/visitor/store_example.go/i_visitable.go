package main

type IVisitable interface {
	Accept(IVisitor)
}
