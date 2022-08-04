package pkg

type Service interface {
	Execute(*Data)
	SetNext(service Service)
}

type Data struct {
	GetSource bool
	Update    bool
}
