package main

import "fmt"

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI

	Build() Computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewComputerBuilder() ComputerBuilderI {
	return computerBuilder{}
}

func (b computerBuilder) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}
func (b computerBuilder) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}
func (b computerBuilder) MB(val string) ComputerBuilderI {
	b.mb = val
	return b
}

func (b computerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func main() {
	compBuilder := NewComputerBuilder()
	computer := compBuilder.CPU("core i3").RAM(8).MB("gigabyte").Build()
	fmt.Println(computer)

	officeCompBuilder := NewOfficeComputerBuilder()
	officeComp := officeCompBuilder.Build()
	fmt.Println(officeComp)
}

type officeComputerBuilder struct {
	computerBuilder
}

func NewOfficeComputerBuilder() ComputerBuilderI {
	return officeComputerBuilder{}
}

func (b officeComputerBuilder) Build() Computer {
	return Computer{
		CPU: "intel pentium II",
		RAM: 2,
		MB:  "gigabyte",
	}
}
