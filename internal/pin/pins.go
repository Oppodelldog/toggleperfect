package pin

type LedPin interface {
	High()
	Low()
}

type KeyPin interface {
	IsKeyPressed() bool
}
