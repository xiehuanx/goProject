package main

import (
	"fmt"
	//"interfaceVsExtend"
	//"utils"
)

type mtint int

//func init()  {
//	utils.CreateData()
//	fmt.Println("init 函数最先执行")
//}
//
//func inter()  {
//	wukong := interfaceVsExtend.LittleMonkey{
//		Monkey: interfaceVsExtend.Monkey{
//			Name : "小猴子",
//		},
//	}
//	wukong.Pashu()
//}



func main() {
	//human := entity.NewPerson(12, "谢环测试", 12)
	//fmt.Println(human.Age())
	//human.Ceshi()
	//fmt.Println(human.Age())
	//
	//xiehuan := entity.NewUser("xiehuan", 12, "测试")
	//fmt.Println(xiehuan)

	//inter()

	main2()

	func
}

func main2()  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误信息：", err)
		}
	}()

	f := addSuper()
	fmt.Println(f(1))
	fmt.Println(f(1))
	panic("Reece")
}

func getSum(n1 int, n2 int) (int, int, int)  {
	return n1 + n2, n1, n2;
}

func funcTest(function func(int, int) (int, int, int), n1 int, n2 int) (int, int, int) {
	return function(n1, n2)
}

func addSuper() func(int) int  {
	var n = 10
	return func(n2 int) int {
		// return n + n2
		n += n2
		return n
	}
}

func sliceDemo()  {
	// 切片不需要指定长度
	numbers := []int{0,1,2,3,4,5,6,7,8}
	fmt.Println(numbers)
	fmt.Printf("%T", numbers)

	datas  := [...]int8{1, 2, 3, 4, 5}
	fmt.Println(datas)
	datas[0] = 20
	datas = [5]int8{}
	fmt.Println(datas)
}

func mapDemo()  {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	//strings.Join()
	if (ok) {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}


}

