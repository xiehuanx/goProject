package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Id int
	Name string
	Age string
	Class string

}

func TestRefect01(b interface{})  {

	reflectType := reflect.TypeOf(b)
	fmt.Println("reflect 的类型是：", reflectType)
	reflectValue := reflect.ValueOf(b)
	fmt.Println("reflect 的值是：", reflectValue)
	//newValue := reflectValue.Int() +2
	//fmt.Println("操作之后 的值是：", newValue)
	valueInfer :=reflectValue.Interface()
	fmt.Println("valueInfer 的值是：", valueInfer)
	 ceshi := valueInfer.(int)
	fmt.Println("ceshi 的值是：", ceshi)
}

func TestRefect02(b interface{})  {

	reflectType := reflect.TypeOf(b)
	fmt.Println("reflect 的类型是：", reflectType)
	reflectValue := reflect.ValueOf(b)
	fmt.Println("reflect 的值是：", reflectValue)
	//reflectValue.
	valueInfer :=reflectValue.Interface()
	fmt.Println("valueInfer 的值是：", valueInfer)
	a := valueInfer.(student)
	fmt.Println(a)
	fmt.Println("kind 的值是：", reflectType.Kind())
	reflectValue.Kind()
}

func TestReflect03(b interface{})  {
	reflectValue := reflect.ValueOf(b)
	fmt.Println("value 的值是：", reflectValue)
	fmt.Println("Type 的值是：", reflectValue.Type())
	fmt.Println("Kind 的值是：", reflectValue.Kind())
	reflectInterface := reflectValue.Interface()
	 newValue := reflectInterface.(float64)
	fmt.Println("newValue 的值是：", newValue)
}

func TestReflect04(str string)  {
	stirs := reflect.ValueOf(str)
	stirs.Elem().SetString("测试")
	fmt.Println("就是：", stirs)
}

func main() {
	//TestRefect01(2)
	//xiehuan := student{
	//	12,
	//	"谢环",
	//	"测试",
	//	"班级",
	//}
	//
	//TestRefect02(xiehuan)
	//var ceshi float64 = 2.7
	//TestReflect03(ceshi)
	var str string = "12121"
	stirs := reflect.ValueOf(&str)
	stirs.Elem().SetString("测试")
	fmt.Println("就是：", str)
}
