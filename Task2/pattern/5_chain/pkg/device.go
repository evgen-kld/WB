package pkg

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Println("Данные уже получены")
		d.Next.Execute(data)
		return
	}
	fmt.Println("Данные успешно получены")
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(service Service) {
	d.Next = service
}
