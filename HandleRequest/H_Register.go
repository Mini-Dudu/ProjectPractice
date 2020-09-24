package HandleRequest

import (
	"LOL/actionsDb"
	"html/template"
	"net/http"
)

//处理物品信息请求
func ParseRegister(writer http.ResponseWriter, request *http.Request) {

	str := "register.html"
	temp, err := template.ParseFiles("./HTML/" + str)

	if err != nil {
		errorTemp, _ := template.ParseFiles("./HTML/error.html")
		errorTemp.Execute(writer, err.Error())
		return
	}
	temp.Execute(writer, actionsDb.ItemSlice)
}
