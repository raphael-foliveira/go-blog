package interfaces

type Repository[T interface{}] interface {
	Find(...string) ([]*T, error)
	FindOne(id int64) (*T, error)
	Create(*T) (*T, error)
	Update(*T) (*T, error)
	Delete(id int64) (int64, error)
}
