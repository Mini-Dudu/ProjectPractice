package HandleRequest

import (
	"LOL/actionsDb"
	"fmt"
	"html/template"
	"net/http"
)
//处理英雄详细信息请求
func Hero_details(writer http.ResponseWriter,request *http.Request){
	B,err := actionsDb.QueryHeroMainImg()
	fmt.Println(B)
	if err != nil {
		fmt.Println(err)
	}
	str := "hero_details.html"
	temp,err := template.ParseFiles("./HTML/" + str)

	if err != nil {
		errorTemp,_ :=template.ParseFiles("./HTML/error.html")
		errorTemp.Execute(writer,err.Error())
		return
	}
	temp.Execute(writer,actionsDb.HeroInfo)
}
