package main

func main() {
	exmain()

	// // With go routine
	// channel := make(chan int)
	// go func() { channel <- 35 }()
	// fmt.Println(<-channel)

	// // With buffer: Don't use this
	// channel := make(chan int, 1)
	// channel <- 35
	// fmt.Println(<-channel)

	// Using channel
	// channel := make(chan int)

	// go channelSend(channel)

	// channelReceive(channel)

}

// func channelSend(cs chan<- int) {
// 	cs <- 35
// }

// func channelReceive(cr <-chan int) {
// 	fmt.Println("O valor recebido do canal é:", <-cr)
// }
