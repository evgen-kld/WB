package pkg

import "fmt"

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    4,
		Memory:  128,
		Monitor: true,
	}
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintInfo() {
	fmt.Println(pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
