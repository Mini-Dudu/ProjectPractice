package HandleRequest

import (
	"LOL/actionsDb"
	"fmt"
	"html/template"
	"net/http"
)

//处理登入求情
func Login(writer http.ResponseWriter,request *http.Request) {

	//1、解析所请求的参数
	err := request.ParseForm()
	if err != nil {
		err,_:= template.ParseFiles("./HTML/error.html")
		err.Execute(writer,"登入出错，请重试!")
		return
	}

	//得到用户登入信息
	userName := request.FormValue("username")
	userPwd := request.FormValue("password")

	fmt.Printf("用户名: %v\n",userName)
	//如果用户未输入账户和密码，则给出相应的提示
	if userName == "" || userPwd == "" {
		err,_:= template.ParseFiles("./HTML/error.html")
		err.Execute(writer,"用户名和密码不能为空!请输入后重试!")
		return
	}

	//template.ParseFiles("./HTML/home.html")

	//2、在数据库中查询用户名是否存在,存在返回true 否则返回false
	IsUser,err := actionsDb.IsUser(userName)
	if err != nil {
		fmt.Println("查询用户名出错")
		return
	}

	//3、用户名存在.....
	if IsUser {
		//进行密码对比 返回bool值
		//fmt.Println("============")
		identifyPwd,err := actionsDb.IdentifyPwd(userName,userPwd)
		if err != nil {
			fmt.Println("确认用户密码出错")
			fmt.Println(err.Error())

		}

		if identifyPwd {
			err,_:= template.ParseFiles("./HTML/home.html")
			err.Execute(writer,"确认密码,密码正确，........")
		}else {
			err,_:= template.ParseFiles("./HTML/error.html")
			err.Execute(writer,"用户名或密码错误，请从新输入正确的用户名和密码!")
		}

	}else {
		err,_:= template.ParseFiles("./HTML/error.html")
		err.Execute(writer,"用户名不存在，请注册后在进行此操作!")
	}



}
