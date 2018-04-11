var pool = make(chan int, 10) // data communication, buffer size = 10
var done = make(chan bool) // synchronize status

func producer(num int) {
	for {
		msg:= rand.Int()
		pool <- msg // auto lock inside channel
	}
}

func consumer(num int) {

	// block when there is no data in pool
	// when sender sends data, weak the goroutine automatically
	for msg:= range pool {
		fmt.Println(msg)
	}

	done <- true // close the pool, done
}

func main() {
	producerCount:= 2
	consumerCount:= 4

	for i:= 0; i < producerCount; i++ {
		go producer(i) // start new gorountine
	}

	for i:= 0; i < consumerCount; i++ {
		go consumer(i)	// start new gorountine
	}

	// main goroutine blocks, when done, exits the program
	<-done
}