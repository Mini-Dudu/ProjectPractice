package main

import (
	"LOL/HandleRequest"
	"LOL/actionsDb"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("欢迎来到LOL")
	//
	////请求的网址
	//url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"			//英雄列表url
	//
	////url = "https://game.gtimg.cn/images/lol/act/img/js/items/items.js"					//物品列表url
	////解析JSON类型数据后、得到并返回byte字节切片
	//data := parse.ParseJSON(url)
	//fmt.Println(string(*data))
	//
	////定义英雄列表结构体类型变量
	//hero_list := object.HeroDetailed{}
	//
	//err := json.Unmarshal(*data,&hero_list)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(hero_list)
	//打开数据库

	if actionsDb.Db == nil {
		err := actionsDb.OpenDb()
		if err != nil {
			fmt.Println("打开数据库失败!")
		}
	}
	//
	//////插入简略英雄信息
	////err = actionsDb.Insert_heroInfo(&hero_list)
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	////
	//////插入英雄详细信息
	////err = actionsDb.Insert_detailsInfo(&hero_list)
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	//
	//url = "https://game.gtimg.cn/images/lol/act/img/js/items/items.js"					//物品列表url
	//
	////解析JSON类型数据后、得到并返回byte字节切片
	//data = parse.ParseJSON(url)
	//
	//fmt.Println(string(*data))
	//fmt.Printf("%T\n",data)
	////定义物品信息结构体类型变量
	//itemList := object.AllItem{}
	//err = json.Unmarshal(*data,&itemList)
	//fmt.Println(itemList)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////插入物品信息
	//err = actionsDb.Insert_itemInfo(&itemList)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//
	//
	////https://lol.qq.com/data/info-defail.shtml?id=1		http://game.gtimg.cn/images/lol/act/img/js/hero/10.js
	//HeroDetailedSlice := []object.HeroDetailed{}
	//for i := 1; i < 888; i++ {
	//
	//	s := strconv.Itoa(i)
	//
	//	url := "http://game.gtimg.cn/images/lol/act/img/js/hero/" + s + ".js"					//英雄详细信息url
	//	//解析JSON类型数据后、得到并返回byte字节切片
	//	data := parse.ParseJSON(url)
	//	//fmt.Println(string(*data))
	//
	//	//定义英雄列表结构体类型变量
	//	HeroDeta := object.HeroDetailed{}
	//
	//	err := json.Unmarshal(*data,&HeroDeta)
	//	if err != nil {
	//		continue
	//	}
	//
	//	HeroDetailedSlice = append(HeroDetailedSlice,HeroDeta)
	//	//fmt.Println(HeroDeta.Hero.HeroId)
	//}
	//fmt.Println(HeroDetailedSlice[1].Skins[1].MainImg)
	////插入英雄背景信息
	//err = actionsDb.Insert_backstroy(&HeroDetailedSlice)
	//fmt.Println(err)
	//
	////插入英雄技能信息
	//err = actionsDb.Insert_skili(&HeroDetailedSlice)
	//fmt.Println(err)
	//
	////插入英雄皮肤信息
	//err = actionsDb.Insert_skin(&HeroDetailedSlice)
	//fmt.Println(err)
	//
	//url = "http://game.gtimg.cn/images/lol/act/img/js/hero/10.js"					//英雄详细信息url
	////解析JSON类型数据后、得到并返回byte字节切片
	//data = parse.ParseJSON(url)
	//fmt.Println(string(*data))
	//
	////定义英雄列表结构体类型变量
	//hero_list = object.HeroDetailed{}
	//
	//err = json.Unmarshal(*data,&hero_list)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//
	////for i := 0; i < len(hero_list.Hero); i++ {
	////	for j := 0; j < 1; j++{
	////		fmt.Println(hero_list.Hero[i].Title)
	////	}
	////	fmt.Println(hero_list.Hero[i])
	////}
	//fmt.Println(hero_list)

	http.Handle("/static/img/", http.StripPrefix("/static/img/", http.FileServer(http.Dir("./HTML/img/"))))
	http.HandleFunc("/",HandleRequest.ParseRegister)
	http.HandleFunc("/login",HandleRequest.Login)
	http.HandleFunc("/article.html",HandleRequest.Article)
	http.HandleFunc("/insufficiency.html",HandleRequest.Insufficiency)
	http.HandleFunc("/hero_details.html",HandleRequest.Hero_details)
	http.HandleFunc("/team.html",HandleRequest.Team)

	fmt.Println("服务已开启......")
	err := http.ListenAndServe(":9091",nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	//关闭数据库
	err = actionsDb.CloseDb()
	if err != nil {
		fmt.Println("关闭数据库失败!")
	}

}


