package vec2

type Vec2 struct {
	X int
	Y int
}

func Add(a Vec2, b Vec2) Vec2 {
	return Vec2{
		a.X + b.X,
		b.Y + b.Y,
	}
}
