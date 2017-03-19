package entities

type Entity interface {
	Type() string

	Handle() int
	SetHandle(handle int)

	Layer() string
	SetLayer(layer string)
}

type baseEntity struct {
	handle int    // 5
	layer  string // 8
}

func (e *baseEntity) Handle() int {
	return e.handle
}

func (e *baseEntity) SetHandle(handle int) {
	e.handle = handle
}

func (e *baseEntity) Layer() string {
	return e.layer
}

func (e *baseEntity) SetLayer(layer string) {
	e.layer = layer
}
