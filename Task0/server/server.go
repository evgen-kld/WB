package server

import (
	"Task0/cache"
	"fmt"
	"net/http"
)

func Start() {
	fmt.Println("Server start")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "server/templates/start_page.html")
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		a, err := cache.Cache[id]
		var str string
		if err == false {
			str = "Ничего не найдено"
		} else {
			str = fmt.Sprintf("Номер:%s ФИО:%s", a.TrackNumber, a.Delivery.Name)
		}
		// вывод данных по полученному id
		fmt.Fprintf(w, str)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
