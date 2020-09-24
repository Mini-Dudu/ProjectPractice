package HandleRequest

import (
	"html/template"
	"net/http"
)

//处理团队详情信息请求
func Team(writer http.ResponseWriter,request *http.Request){

	str := "team.html"
	temp,err := template.ParseFiles("./HTML/" + str)

	if err != nil {
		errorTemp,_ :=template.ParseFiles("./HTML/error.html")
		errorTemp.Execute(writer,err.Error())
		return
	}
	temp.Execute(writer,nil)
}
