package main

import (
	"fmt"
	"head-first-go-exercises/chapter11/gadget"
)

type Player interface {
	Play(song string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Not a tape recorder")
	}

}

func main() {
	TryOut(gadget.TapePlayer{})
	TryOut(gadget.TapeRecorder{})
}
