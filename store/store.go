package store

type Store interface {
	Save()
	Get()
	Delete()
}
