package main

type Context struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         IState
}
