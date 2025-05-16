package main

import "tape/gadget"

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

func main() {
	var player Player
	player = gadget.TapeRecorder{}
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)
	player = gadget.TapePlayer{}
	playList(player, mixtape)
}
