package main

import (
	"github.com/gvalkov/golang-evdev"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	dev, err := evdev.Open("/dev/input/by-id/usb-0801_0001-event-kbd")
	if err != nil {
		log.Fatalf("Open: %v", err)
	}
	log.Infof("%s\n", dev.String())

	// We grab the device so it doesnt write on stdout
	if err := dev.Grab(); err != nil {
		log.Fatalf("Grab: %v", err)
	}
	defer func() {
		if err := dev.Release(); err != nil {
			log.Fatalf("Grab: %v", err)
		}
	}()

	var card string
	var cnt int
	for {
		evt, err := dev.ReadOne()
		if err != nil {
			log.Fatalf("Read: %v", err)
		}
		if cnt == 0 || cnt == 1 {
			log.Printf("% X % X % X\n", evt.Code, evt.Type, evt.Value)
		}
		if evt.Code == evdev.KEY_KPENTER && evt.Type == evdev.EV_KEY && evt.Value == 0 {
			log.Infof(card)
			card = ""
			cnt = 0
		} else {
			if evt.Type == evdev.EV_KEY && evt.Value == 1 {
				card += constToString(evt.Code)
			}
			cnt++
		}

	}
}

// %A7027879101034555532           ^                          ^1218100?;7027879101034555532=12187991000030000?
// %B4585560018423366^SERPOUL VINCENT          /^23012211000000521000000?;4585560018423366=230122110000521?+01458556149343287496=7027020000000000032004000000020000023011=000000000000=000000000000=0=100432?
// %B4119110092802823^ /SERPOUL VINCENT         ^2505201000000000000000600000000?;4119110092802823=250520100000600001?

func constToString(c uint16) string {
	switch c {
	case evdev.KEY_1:
		return "1"
	case evdev.KEY_2:
		return "2"
	case evdev.KEY_3:
		return "3"
	case evdev.KEY_4:
		return "4"
	case evdev.KEY_5:
		return "5"
	case evdev.KEY_6:
		return "6"
	case evdev.KEY_7:
		return "7"
	case evdev.KEY_8:
		return "8"
	case evdev.KEY_9:
		return "9"
	case evdev.KEY_0:
		return "0"
	case evdev.KEY_A:
		return "A"
	case evdev.KEY_B:
		return "B"
	case evdev.KEY_C:
		return "C"
	case evdev.KEY_D:
		return "D"
	case evdev.KEY_E:
		return "E"
	case evdev.KEY_F:
		return "F"
	case evdev.KEY_G:
		return "G"
	case evdev.KEY_H:
		return "H"
	case evdev.KEY_I:
		return "I"
	case evdev.KEY_J:
		return "J"
	case evdev.KEY_K:
		return "K"
	case evdev.KEY_L:
		return "L"
	case evdev.KEY_M:
		return "M"
	case evdev.KEY_N:
		return "N"
	case evdev.KEY_O:
		return "O"
	case evdev.KEY_P:
		return "P"
	case evdev.KEY_Q:
		return "Q"
	case evdev.KEY_R:
		return "R"
	case evdev.KEY_S:
		return "S"
	case evdev.KEY_T:
		return "T"
	case evdev.KEY_U:
		return "U"
	case evdev.KEY_V:
		return "V"
	case evdev.KEY_W:
		return "W"
	case evdev.KEY_X:
		return "X"
	case evdev.KEY_Y:
		return "Y"
	case evdev.KEY_Z:
		return "Z"
	}
	return ""
}
