package main

import (
	"fmt"
	"tape/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func tryOut(device Player) {
	device.Play("One more time")
	device.Stop()
	recorder, ok := device.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player is not a recorder")
	}
}

func main() {
	var player Player
	player = gadget.TapeRecorder{}
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)
	tryOut(player)
	player = gadget.TapePlayer{}
	playList(player, mixtape)
	tryOut(player)
}
