package pkg

type Service interface {
	GetData(user string) ([]string, error)
}
