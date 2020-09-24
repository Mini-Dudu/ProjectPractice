package parse

import (
	//"LOL/object"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//解析JSON类型数据后、得到并返回byte字节切片
func ParseJSON(url string) *[]byte {

	//模拟客户端
	client := http.Client{}

	//添加请求
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//发送请求
	res,err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//获取byte切片
	Data,err := ioutil.ReadAll(res.Body)//得到的是json格式的数据
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(string(Data))
	//fmt.Println(Data)
	//Data_list := object.Hero_list{}
	//
	//err = json.Unmarshal(Data,&Data_list)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	//return &Data_list
	//
	return &Data
}



