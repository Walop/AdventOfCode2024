package vec2

type Vec2 struct {
	X int
	Y int
}

func Add(a Vec2, b Vec2) Vec2 {
	return Vec2{
		a.X + b.X,
		a.Y + b.Y,
	}
}

func (v *Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func (v *Vec2) Substract(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

func (v *Vec2) Reverse() Vec2 {
	return Vec2{
		-v.X,
		-v.Y,
	}
}
