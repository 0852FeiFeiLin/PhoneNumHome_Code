package active

import (
	"PhoneNumHome_Code/entity"
	"PhoneNumHome_Code/page_parse"
	"PhoneNumHome_Code/request"
	"fmt"
	"html/template"
	"net/http"
)

//跳转到首页查询页面
func ToIndex(w http.ResponseWriter,r *http.Request){
	files, err := template.ParseFiles("./view/index.html")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("跳转index页面失败")
		return
	}
	fmt.Println("跳转index页面成功")
	files.Execute(w,nil)

}
//手机号码查询功能逻辑
func Index(w http.ResponseWriter,r *http.Request){
	//解析用户输入表单中的内容
	err := r.ParseForm()
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("解析用户表单信息失败")
		return
	}
	//获取输入的值
	value := r.FormValue("phoneNo")
	fmt.Println(value)
	var Data entity.Data
	Data.PhoneNo = value

	//定义url，发起请求，返回响应数据
	APIURL := "http://apis.juhe.cn/mobile/get"
	KEY := "da4e36772073edd03728dc50a06f0653"
	//http://apis.juhe.cn/mobile/get?phone=13429667914&key=您申请的KEY&dtype=json、xml、string
	//定义url
	//不指定类型请求（方式1）
	//urlPath := APIURL + "?phone=" + Data.PhoneNo + "&key=" + KEY

	//指定类型请求（方式2）
	urlPath1 := APIURL + "?phone=" + Data.PhoneNo + "&key=" + KEY + "&dtype=" + "xml"
	//发起请求
	resp, err := request.Request("POST", urlPath1, nil)
	//解析返回的[]byte,
	//Data = page_parse.ParseJson(resp)  后续操作JSON解析的就直接把所有Data1换成Data
	Data1 := page_parse.ParseXML(resp)
	//判断返回码，为0成功
	if Data.Error_code != 0{
		fmt.Println("请求出现错误")
		return
	}
	fmt.Println("详情结果集",Data1.Result1)
	//让结果显示在页面的表格上
	files, err := template.ParseFiles("./view/index.html")
	//方法1：将结构体数据放入map里面，便于前端模板遍历输出
	data := map[string]interface{}{
		"Phone_No" : value,
		"Phone_Info" : Data1.Result1,
	}//注意事项：JSON对象只支持string类型作为key,所以编写map的时候，必须是map[string]T这种格式

	//发送结构体数据到模板里面
	err = files.Execute(w, data)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("跳转失败")
		return
	}
	/*//方法2：将结构体数据放入slice里面，便于前端模板遍历输出
	Result := []entity.Result{
		entity.Result{
			City: Data.Result1.City,
			Province : Data.Result1.Province,
			Zip : Data.Result1.Zip,
			Areacode: Data.Result1.Areacode,
			Company: Data.Result1.Company,

		},

	}
	err = files.Execute(w, Result)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("跳转失败")
		return
	}*/
}
