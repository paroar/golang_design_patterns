package main

type IState interface {
	executeState(*Context) bool
}
