package HandleRequest

import (
	"html/template"
	"net/http"
)

//处理注册页面请求
func Insufficiency(writer http.ResponseWriter,request *http.Request){

	str := "insufficiency.html"
	temp,err := template.ParseFiles("./HTML/" + str)

	if err != nil {
		errorTemp,_ :=template.ParseFiles("./HTML/error.html")
		errorTemp.Execute(writer,err.Error())
		return
	}
	temp.Execute(writer,nil)
}
