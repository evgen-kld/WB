package pkg

const (
	ServerType           = "server"
	PersonalComputerType = "pc"
)

type Computer interface {
	GetType() string
	PrintInfo()
}

func NewComputer(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	default:
		return nil
	}
}
