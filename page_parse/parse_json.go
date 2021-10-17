package page_parse

import (
	"PhoneNumHome_Code/entity"
	"encoding/json"
	"fmt"
)

//解析读取的响应数据
func ParseJson(resp []byte)(re entity.Data){
	var data entity.Data
	//反序列化
	jsonerr := json.Unmarshal(resp, &data)  //结果存在结构体data里面
	if jsonerr != nil {
		fmt.Println(jsonerr.Error())
		fmt.Println("json反序列化失败")
		return
	}
	fmt.Println(data)
	//返回解析好的数据，存在结构体里面
	return data
}