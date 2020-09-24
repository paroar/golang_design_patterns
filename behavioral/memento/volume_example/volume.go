package main

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}
