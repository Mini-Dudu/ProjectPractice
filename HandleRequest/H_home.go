package HandleRequest

//import (
//	"html/template"
//	"net/http"
//)
//
////处理主页面请求
//func Parse1(writer http.ResponseWriter,request *http.Request){
//	temp,err := template.ParseFiles("./HTML/home.html")
//	if err != nil {
//		errorTemp,_ :=template.ParseFiles("./HTML/error.html")
//		errorTemp.Execute(writer,err.Error())
//		return
//	}
//	temp.Execute(writer,nil)
//}
