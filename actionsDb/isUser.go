package actionsDb

import "fmt"

func IsUser(userName string) (B bool,err error) {
	//fmt.Println("判断用户名是否存在'!")

	B = QueryUser(userName)
	fmt.Println("用户名是否存在:",B)
	//QueryItem()
	//QueryHeroDetails()
	//QueryHeroMainImg()
	return B,nil
}

