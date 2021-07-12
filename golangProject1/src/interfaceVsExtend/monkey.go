package interfaceVsExtend

import "fmt"

type Monkey struct {
	Name string
}

func (monkry *Monkey) Pashu()  {
	fmt.Println(monkry.Name, "这是一只")
}


type LittleMonkey struct {
	Monkey
}

func main()  {

}