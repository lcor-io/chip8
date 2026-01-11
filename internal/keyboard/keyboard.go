package keyboard

import "time"

const (
	KEYBOARD_SIZE = 16
	DEBOUNCE_TIME = 10 * time.Millisecond
)

type Keyboard struct {
	Keys [16]bool
}
