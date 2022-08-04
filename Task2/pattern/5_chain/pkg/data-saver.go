package pkg

import "fmt"

type SaveData struct {
	Name string
	Next Service
}

func (sd *SaveData) Execute(data *Data) {
	if !data.Update {
		fmt.Println("Данные не обработаны")
		return
	}
	fmt.Println("Данные успешно сохранены")
}

func (sd *SaveData) SetNext(service Service) {
	sd.Next = service
}
