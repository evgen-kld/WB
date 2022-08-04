package pkg

import "fmt"

type UpdateData struct {
	Name string
	Next Service
}

func (upd *UpdateData) Execute(data *Data) {
	if data.Update {
		fmt.Println("Данные уже обработаны")
		upd.Next.Execute(data)
		return
	}
	fmt.Println("Данные успешно обработаны")
	data.Update = true
	upd.Next.Execute(data)
}

func (upd *UpdateData) SetNext(service Service) {
	upd.Next = service
}
