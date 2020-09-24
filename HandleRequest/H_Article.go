package HandleRequest

import (
	"LOL/actionsDb"
	"fmt"
	"html/template"
	"net/http"
)

//处理物品信息请求
func Article(writer http.ResponseWriter,request *http.Request){
	B,err :=actionsDb.QueryItem()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("是否有数据:",B)
	str := "article.html"
	temp,err := template.ParseFiles("./HTML/" + str)

	if err != nil {
		errorTemp,_ :=template.ParseFiles("./HTML/error.html")
		errorTemp.Execute(writer,err.Error())
		return
	}

	temp.Execute(writer,&actionsDb.ItemSlice)
}
