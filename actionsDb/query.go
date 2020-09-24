package actionsDb

import (
	"LOL/object"
	"fmt"
)


func QueryUser(user string) (B bool){
	//fmt.Println("query查询",user)
	rows,_ := Db.Query("select LOL_userName from lol_user where LOL_userName=?",user)

	for rows.Next() {
		var userName string
		//fmt.Println("写入前",len(userName))
		err := rows.Scan(&userName)
		//fmt.Println("写入后",len(userName))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(len(userName))
		if len(userName) == 0 {
			B = false
		}else {
			B = true
		}
	}
	return
}

func QueryPwd(user,Pwd string)(B bool, err error) {
	rows,_ := Db.Query("select LOL_userName, LOL_userPwd from lol_user where LOL_userName=?",user)

	for rows.Next() {
		var userName string
		var pwd string

		err = rows.Scan(&userName,&pwd)

		if Pwd != pwd {
			B = false
		}else if Pwd == pwd {
			B = true
		}
	}
	return
}

var ItemSlice []object.Adhguhg
//查询物品信息
func QueryItem()(B bool, err error) {
	rows,_ := Db.Query("select item_id, item_Name, item_IconPath, item_Price, item_Description, item_Plaintext, item_Sell, item_Total from item_info")		// where item_id<?",1030

	for rows.Next() {
		var itemInfos object.Adhguhg

		err = rows.Scan(&itemInfos.ItemId,&itemInfos.Name,&itemInfos.IconPath,&itemInfos.Price,&itemInfos.Description,&itemInfos.Plaintext,&itemInfos.Sell,&itemInfos.Total)
		ItemSlice = append(ItemSlice,itemInfos)
		//fmt.Println(itemInfos)

	}
	fmt.Println(ItemSlice[5])
	fmt.Println("物品长度",len(ItemSlice))
	B = true
	return
}

var HeroInfoSilce []object.HeroInfo

//查询英雄详细信息
func QueryHeroDetails()(B bool, err error) {
	rows,_ := Db.Query("SELECT h.hero_id,h.hero_alias,h.hero_name," +
		"H.hero_engName,H.hero_IsWeekFree,H.hero_Attack,H.hero_Defense,H.hero_Magic,H.hero_Difficulty,H.hero_SelectAudio," +
		"H.hero_BanAudio FROM hero_briefinfo h,(SELECT * FROM hero_detailsinfo) H where h.hero_id=H.hero_idDetails;")

	for rows.Next() {
		var heroInfos object.HeroInfo

		err = rows.Scan(&heroInfos.HeroId,&heroInfos.Name,&heroInfos.Alias,&heroInfos.Title,
			&heroInfos.IsWeekFree,&heroInfos.Attack,&heroInfos.Defense,&heroInfos.Magic,&heroInfos.Difficulty,
			&heroInfos.SelectAudio,&heroInfos.BanAudio)
		HeroInfoSilce = append(HeroInfoSilce,heroInfos)
		//fmt.Println(itemInfos)
		//fmt.Println(ItemSlice)
	}
	fmt.Println(len(HeroInfoSilce))
	//for i := 0; i < 10; i++ {
	//	fmt.Println(HeroInfoSilce[i])
	//}
	return
}


var HeroInfo []object.Skin
//查询英雄简略信息及main图片
func QueryHeroMainImg()(B bool, err error) {

		rows,_ := Db.Query("SELECT s.skin_name, s.skin_mainImg, bj.hero_backStory FROM (SELECT hero_id, skin_name, skin_mainImg FROM hero_skins WHERE skin_id=hero_id*1000) s,hero_backStory bj WHERE s.hero_id=bj.hero_id;")

		for rows.Next() {
			var heroInfos object.Skin

			err = rows.Scan(&heroInfos.Name,&heroInfos.MainImg,&heroInfos.ShortBio)
			HeroInfo = append(HeroInfo,heroInfos)
			//fmt.Println(heroInfos)
		}
	//fmt.Println(rows == nil)


	//for i := 0; i < 10; i++ {
	//	fmt.Println(HeroInfo[i])
	//}
	return
}