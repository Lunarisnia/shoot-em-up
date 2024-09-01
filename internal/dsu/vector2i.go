package dsu

type Vector2i struct {
	X int32
	Y int32
}

func (v Vector2i) Subtract(y Vector2i) Vector2i {
	return Vector2i{
		X: v.X - y.X,
		Y: v.Y - y.Y,
	}
}

func (v Vector2i) Add(y Vector2i) Vector2i {
	return Vector2i{
		X: v.X + y.X,
		Y: v.Y + y.Y,
	}
}

func (v Vector2i) MultiplyVector(y Vector2i) Vector2i {
	return Vector2i{
		X: v.X * y.Y,
		Y: v.Y * y.Y,
	}
}

func (v Vector2i) MultiplyScalar(y int32) Vector2i {
	return Vector2i{
		X: v.X * y,
		Y: v.Y * y,
	}
}
