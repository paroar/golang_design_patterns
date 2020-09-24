package main

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}
