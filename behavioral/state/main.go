package main

func main() {
	start := StartState{}
	game := &Context{
		Next: &start,
	}

	for game.Next.executeState(game) {
	}

}
