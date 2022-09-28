package ex9_4

func RunGoroutinePipeline(
	in chan int,
	out chan<- int,
	size int,
) {
	var pipe, nextPipe chan int
	pipe = in
	for i := 0; i < size; i++ {
		nextPipe = make(chan int)
		go func(in <-chan int, out chan<- int) {
			n := <-in
			out <- n
		}(pipe, nextPipe)
		pipe = nextPipe
	}

	result := <-nextPipe
	out <- result
}
