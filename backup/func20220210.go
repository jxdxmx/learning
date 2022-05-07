package backup

import "fmt"

func f20220210() {
	var p = People{
		age:  10,
		name: "xu meng",
	}
	defer func() {
		fmt.Println(1, p.name, " age:", p.age)
	}()
	defer func(p People) {
		fmt.Println(2, p.name, " age:", p.age)
	}(p)
	defer func(p *People) {
		fmt.Println(3, p.name, " age:", p.age)
	}(&p)
	p = People{
		age:  20,
		name: "小田",
	}

}

type People struct {
	age  int
	name string
}
