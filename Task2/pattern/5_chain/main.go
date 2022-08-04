package main

import "Task2/pattern/5_chain/pkg"

func main() {
	device := &pkg.Device{Name: "device1"}
	dataUpdater := &pkg.UpdateData{Name: "updater1"}
	dataSaver := &pkg.SaveData{}
	device.SetNext(dataUpdater)
	dataUpdater.SetNext(dataSaver)
	data := &pkg.Data{}
	device.Execute(data)
}
