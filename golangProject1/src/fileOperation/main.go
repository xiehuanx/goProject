package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("测试")
	file , error := os.OpenFile("F:\\IDE.txt", os.O_RDWR,0600)

	//reder := bufio.NewReader(file)

	//reder.ReadString()


	defer file.Close()
	if error == nil {
		fmt.Println(file.Name())
		contentByte,err :=ioutil.ReadAll(file)
		if err == nil {
			fmt.Println(string(contentByte))
			_,errorY :=file.Write([]byte("测试测试写入文件"))
			if errorY == nil {

			}
		}

	}
}
