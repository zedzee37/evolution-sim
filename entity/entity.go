package entity

type Entity interface {
	Load()
	Tick(dt float64)
}

type Drawable interface {
	Draw(offsetX int32, offsetY int32)
}

type DrawableEntity interface {
	Entity
	Drawable
}
