package main

var counter = func(n int) chan chan<- int {
	requests := make(chan chan<- int)

	go func() {
		for request := range requests {
			if request == nil {
				n++
			} else {
				request <- n
			}
		}
	}()

	return requests
}(0)

// implicitly converted to chan chan <- (chan <- int)

func main() {

}
