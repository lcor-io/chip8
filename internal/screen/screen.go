package screen

const (
	DEFAULT_SCREEN_WIDTH  uint8 = 64
	DEFAULT_SCREEN_HEIGHT uint8 = 32
)

type pixel bool

type Screen interface {
	Clear()
	Render() string
	GetPixel(x int, y int) pixel
	SetPixel(x int, y int, val pixel)
}
