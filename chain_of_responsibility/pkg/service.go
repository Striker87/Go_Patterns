package pkg

type Service interface {
	Execute(data *Data)
	SetNext(service Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}
