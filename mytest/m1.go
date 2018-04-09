package main


type Animal interface{
	GetID() int
	SetID(id int)
}

type People struct{
	id int
	name string
}

func (p *People) GetID() int {
	return p.id
}

func (p *People) SetID(id int) {
	p.id = id
}

func main() {
	var a Animal
	p := People{id:1, name:"aaa"}
	a = &p
	_ = a
}
