package main

import (
	"flag"
	"fmt"
)

func main() {

	//args := os.Args
	//for i, v := range args {
	//	fmt.Println(i, v)
	//}
	var name string
	var password string
	var age string
	flag.StringVar(&name, "n", "", "测试")
	flag.StringVar(&password, "p", "", "测试")
	flag.StringVar(&age, "a", "", "测试")
	flag.Parse()
	fmt.Println(name, password, age)
}
