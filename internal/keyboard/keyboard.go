package keyboard

import "time"

const (
	KEYBOARD_SIZE = 16
	DEBOUNCE_TIME = 10 // In ms
)

type Keyboard struct {
	Keys [16]bool

	keyDebounce [16]uint
	keyTicker   *time.Ticker
}

func New() *Keyboard {

	keyboard := Keyboard{}

	keyboard.keyTicker = time.NewTicker(time.Millisecond)

	go func() {
		for range keyboard.keyTicker.C {
			for i := range keyboard.keyDebounce {
				if keyboard.keyDebounce[i] > 0 {
					keyboard.keyDebounce[i]--
				} else {
					keyboard.Keys[i] = false
				}
			}

		}
	}()

	return &keyboard
}
