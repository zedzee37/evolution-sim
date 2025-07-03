package entity

type Entity interface {
	OnLoad()
	Tick(dt float64)
}

type Drawable interface {
	Draw(offsetX int32, offsetY int32)
}
