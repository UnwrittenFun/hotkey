package main

import (
	"fmt"
	"log"

	"github.com/UnwrittenFun/hotkey"
)

func main() {
	hk := hotkey.NewListener()

	_, err := hk.CreateAndRegisterHotkey(hotkey.ModCtrl+hotkey.ModAlt+hotkey.ModShift, 'P', func() {
		fmt.Println("It was pressed :o")
	})
	if err != nil {
		log.Fatal(err)
	}

	hk.Listen()
}
