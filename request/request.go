package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//发起请求
func Request(method string ,url string,body io.Reader) (re []byte,err error){
	//http://apis.juhe.cn/mobile/get?phone=13429667914&key=您申请的KEY
	//创建客户端
	client := http.Client{
		//设置超时时间
		Timeout: 30 * time.Second,
	}
	//创建请求
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("请求创建失败")
		return
	}
	//客户端发起请求
	response, err := client.Do(request)
	if response.StatusCode != 200{
		fmt.Println("创建客户端请求失败")
		fmt.Println(err.Error())
		return
	}
	//返回的响应进行处理
	fmt.Println("获得响应啦？")
	//读取响应体数据
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	//返回响应体数据
	return bytes,nil
}
