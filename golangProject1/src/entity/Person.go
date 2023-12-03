package entity

type Person struct {
	id int
	name string
	age  int
}


func (p *Person) String() string {
	p.age = 10
	panic(p)
}

func (p *Person) Ceshi()  {
	p.age = 10
}

type User struct{
	name    string
	age     int
	address string
}

func NewUser(name string, age int, address string) *User {
	return &User{name: name, age: age, address: address}
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) Age() int {
	return u.age
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) Address() string {
	return u.address
}

func (u *User) SetAddress(address string) {
	u.address = address
}

func NewPerson(id int, name string, age int) *Person {
	return &Person{id: id, name: name, age: age}
}

func (p *Person) Id() int {
	return p.id
}

func (p *Person) SetId(id int) {
	p.id = id
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p Person) jisuan()  {
	
}












