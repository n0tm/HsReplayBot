package service

type Loader interface {
	Load() (interface{}, error)
}
