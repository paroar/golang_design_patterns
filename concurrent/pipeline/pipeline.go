package pipeline

func LaunchPipelines(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)

	result := <-thirdCh
	return result
}

func generator(max int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func power(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()

	return out
}
