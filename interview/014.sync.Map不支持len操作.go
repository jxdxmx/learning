package interview

import (
	"sync"
)

func F014() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	//fmt.Println(len(m)) // Invalid argument for the len function
}
