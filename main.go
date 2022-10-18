package main

import (
	"math/rand"
	"sync"
)

func main() {
	const max = 100000
	const numSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	dataCh := make(chan int)
	// stopCh is an additional signal channel
	// Its sender is the receiver of channel
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < numSenders; i++ {
		go func() {
			for {
				// The try-receive operation is to try
				// to exit the goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == max-1 {
				// The receiver of channel dataCh is also the sender of stopCh
				close(stopCh)
				return
			}
		}
	}()
}
