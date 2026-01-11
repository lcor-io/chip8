package screen

const (
	SCREEN_WIDTH  uint8 = 64
	SCREEN_HEIGHT uint8 = 32
)

type Pixel bool

type Screen interface {
	Clear()
	View() string
	GetPixel(x int, y int) Pixel
	SetPixel(x int, y int, val Pixel)
}
