package main

import (
	"PhoneNumHome_Code/active"
	"fmt"
	"net/http"
)

/*
	手机号码归宿地查询：
		思路：
			1、用户输入手机号码
			2、发送请求，可指定返回数据格式xml、json、string
			3、响应数据解析，反序列化
			4、定义结构体接收响应数据
			5、前端定义表格接收
*/
func main() {
	http.HandleFunc("/",active.ToIndex)
	http.HandleFunc("/phoneNum",active.Index)

	err := http.ListenAndServe(":8099", nil)
	if err != nil{
		fmt.Println("启动服务器失败")
		fmt.Println(err.Error())
		return
	}

}