package page_parse

import (
	"PhoneNumHome_Code/entity"
	"encoding/xml"
	"fmt"
)

//解析返回的xml数据格式数据
func ParseXML(resp []byte) (re entity.Data){
	var data entity.Data
	err := xml.Unmarshal(resp, &data)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("xml格式数据解析失败")
		return
	}
	fmt.Println("xml数据格式解析。。。。")
	fmt.Println(data)
	return data
}
