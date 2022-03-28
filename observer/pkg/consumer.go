package pkg

type Consumer interface {
	Update(pubName string)
	GetName() string
}
