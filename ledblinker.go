package main

import (
	"time"
"fmt"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

func main() {
	embd.InitGPIO()
defer embd.CloseGPIO()

    ledPin,_ :=  embd.NewDigitalPin(18)
    ledPin.SetDirection(embd.Out)
    pinStatus := false

    tick := time.Tick(time.Millisecond * 1000)

    for {
        select {
	    case <- tick:
		if(pinStatus){
			ledPin.Write(embd.Low)
			pinStatus = false
		}else{
			ledPin.Write(embd.High)
			pinStatus = true
		}
		fmt.Println("tick!")
        }
    }
}
