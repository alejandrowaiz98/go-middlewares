package http

type Http struct {
}

type HttpImplementation interface {
}

func New() HttpImplementation {

	return Http{}
}
