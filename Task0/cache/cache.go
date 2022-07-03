package cache

import (
	"Task0/JSON"
	"Task0/db"
	"fmt"
)

var Cache = make(map[string]JSON.StructJSON, 0)

// создание кэша
func init() {
	Cache = db.GetCashe(Cache)
	fmt.Println("Cache created!")
}

// добавление нового элемента в кэш
func AddToCache(elem JSON.StructJSON) (string, bool) {
	_, ok := Cache[elem.OrderUID]
	if ok {
		str := fmt.Sprintf("Элемент %s уже существует", elem.OrderUID)
		return str, false
	}
	Cache[elem.OrderUID] = elem
	str := fmt.Sprintf("Элемент %s успешно добавлен", elem.OrderUID)
	return str, true
}
