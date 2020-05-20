package app

const defaultButtonRadius = 15

type button struct {
	radius   float64
	position pos
	symbol   string
	// symbolPos is relative to position
	symbolPos  pos
	caption    string
	captionPos pos
}

func (b button) SetRadius(v float64) button {
	b.radius = v

	return b
}
func (b button) SetX(v float64) button {
	b.position.x = v

	return b
}
func (b button) SetY(v float64) button {
	b.position.y = v

	return b
}

func buttonDown() button {
	return button{
		radius: defaultButtonRadius,
		symbol: "▼",
		symbolPos: pos{
			x: -9,
			y: +8,
		},
	}
}
func buttonUp() button {
	return button{
		radius: defaultButtonRadius,
		symbol: "▲",
		symbolPos: pos{
			x: -9,
			y: +6,
		},
	}
}
func buttonLeft() button {
	return button{
		radius: defaultButtonRadius,
		symbol: "◄",
		symbolPos: pos{
			x: -11,
			y: +7,
		},
	}
}
func buttonRight() button {
	return button{
		radius: defaultButtonRadius,
		symbol: "►",
		symbolPos: pos{
			x: -7,
			y: +7,
		},
	}
}
