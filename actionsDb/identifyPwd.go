package actionsDb

import "fmt"

func IdentifyPwd(user,Pwd string)(B bool, err error) {
	//fmt.Println("确认密码是否正确!.......")

	B,err = QueryPwd(user,Pwd)
	fmt.Printf("用户名密码是否正确:%v\n\n",B)

	return
}
