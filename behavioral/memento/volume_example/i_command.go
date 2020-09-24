package main

type ICommand interface {
	GetValue() interface{}
}
