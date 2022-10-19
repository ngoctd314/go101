# How to Gracefully Close Channels

1. No easy and universal ways to check whether or not c channel is closed without modifying the status of the channel.
2. Closing a closed channel will panic, so it is dangerous to close a channel if the closer don't know whether or not the channel is closed.
3. Sending values to a closed channel will panic, so it is dangerous to send values to a channel if the senders don't know whether or not the channel is closed.

```go
func isClose(c <-chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

func main() {
	c := make(chan int)
	fmt.Println(isClose(c))
	close(c)
	fmt.Println(isClose(c))
}
```

## The channel closing principle

One general principle of using Go channels is don't close a channel from the receiver side and dont close a channel if the channel has multiple concurrent senders. In other words, we should only close a channel in a sender goroutine if the sender is the only sender of the channel.

## Solutions which close channels rudely

You can use the recover mechanism to prevent the possible panic from crashing your program.

```go
func SafeClose(ch chan T) (closed bool) {
    defer func(){
        if recover() != nil {
            closed = false
        }
    }()

    close(ch)
    return true
}

func SafeSend(ch chan T, value T) (closed bool) {
    defer func(){
        if recover() != nil {
            closed = true
        }
    }()

    ch <- value // panic if ch is closed
    return false // closed = false; return
}
```

## Solutions which close channels politely

```go
type MyChannel struct {
    C chan T
    once sync.Once
}

func NewMyChannel() *MyChannel {
    return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
    mc.once.Do(func() {
        close(mc.C)
    })
}
```

## Solutions which close channels gracefully

**1. M receivers, one sender, the sender says "no more sends" by closing the data channel**


Let the sender close the data channel when it doesn't want to send more.

```go
func sender(max int) <-chan int {
	c := make(chan int, max)
	go func() {
		for i := 0; i < max; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receiver(c <-chan int, max int) {
	wg := sync.WaitGroup{}
	wg.Add(max)
	for i := 0; i < max; i++ {
		// multiple receivers
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Println(v)
			}
		}()
	}
	wg.Wait()
}
```

**2. One receiver, N senders, the only receiver says "please stop sending more" by closing an additional signal channel**

This is a situation a little more complicated than the above one. We can't let the receiver close the data channel to stop data transferring, for doing this will break the channel closing principle. But we can let the receiver close an additional signal channel to notify senders to stop sending values

```go
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
```
As mentioned in the comments, for the additional signal channel, its sender is the receiver of the data channel. The additional signal channel is closed by its only sender, which holds the channel closing principle.

In this example, the channel dataCh is never closed. Yes, channels don't have to be closed. A channel will be eventually garbage collected if no goroutines reference it any more, whether it is closed or not. So the gracefulness of closing a channel here is not to close the channel.

