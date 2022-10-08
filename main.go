package main

import (
	"log"
	"math/rand"
	"time"
)

// Seat ...
type Seat int

// Bar ...
type Bar chan Seat

// ServeCustomerAtSeat ...
func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bar27x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar27x7); seatId++ {
		bar27x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		seat := <-bar27x7
		go bar27x7.ServeCustomerAtSeat(customerId, seat)
	}
}
