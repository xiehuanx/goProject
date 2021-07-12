package main

import (
	"encoding/json"
	"fmt"
)

type staudent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age string `json:"age"`
	Class string `json:"class"`
}

type student struct {
	id int
	name string
	age string
	class string

}

func (s *student) Id() int {
	return s.id
}

func (s *student) SetId(id int) {
	s.id = id
}

func (s *student) Name() string {
	return s.name
}

func (s *student) SetName(name string) {
	s.name = name
}

func (s *student) Age() string {
	return s.age
}

func (s *student) SetAge(age string) {
	s.age = age
}

func (s *student) Class() string {
	return s.class
}

func (s *student) SetClass(class string) {
	s.class = class
}

func main() {
	jsonByte, err := json.Marshal("测试测试")

	fmt.Println(string(jsonByte), err)
	xiehuan := staudent{
		12 ,
		"谢环",
		"13",
		"yiban",
	}

	xiehuanJson, _ := json.Marshal(xiehuan)
	fmt.Println(string(xiehuanJson))
	var test student
	test.SetName("ceshi")
	test.SetAge("12")
	test.SetId(12)
	test.SetClass("测啊hi")

	testJson, _ := json.Marshal(&test)
	fmt.Println(string(testJson))

	Deserialization := staudent{}

	err = json.Unmarshal(xiehuanJson, &Deserialization)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Deserialization)

}
