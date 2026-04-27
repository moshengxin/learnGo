package main

import (
	"fmt"
	"os"
)

func main() {

	//str := "456"
	//os.WriteFile("/learnGo/test.txt", []byte(str), 0666) // 0666全员读写，严禁执行

	//file, _ := os.ReadFile("/learnGo/test.txt")
	//fmt.Println(string(file))
	//
	//createdFile, _ := os.Create("test1.txt")
	//fmt.Println("创建出来的文件名", createdFile.Name())
	//bytes := make([]byte, 3)
	//bytes[0] = 's'
	//bytes[1] = 's'
	//createdFile.Write(bytes)
	//readFileStr, _ := os.ReadFile("test1.txt")
	//fmt.Println("创建出来的文件内容", string(readFileStr))

	file, _ := os.Open("test.txt")
	ret, _ := file.Seek(1, 2)
	ret1, _ := file.Seek(1, 18888)
	fmt.Println(ret, ret1)

}
