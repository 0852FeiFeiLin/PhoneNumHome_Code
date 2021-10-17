package entity

type Data struct {
	Resultcode string `json:"resultcode" xml:"resultcode"` 
	Reason string `json:"reason" xml:"reason"`
	Error_code int `json:"error_code" xml:"error_code"`

	//匿名字段实现继承
	Result1 Result `json:"result" xml:"result"`
	//用户输入的手机号码
	PhoneNo string `form:"phoneNo" `
}