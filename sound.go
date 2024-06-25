package main

import (
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

func PlaySound(asBytes []byte) {
	_, data, _ := minimp3.DecodeFull(asBytes)
	context, _ := oto.NewContext(22050, 2, 2, 4096)

	defer context.Close()
	player := context.NewPlayer()
	defer player.Close()
	player.Write(data)
}
