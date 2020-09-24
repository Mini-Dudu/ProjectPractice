package actionsDb

import (
	"LOL/object"
)
//添加英雄简略信息
func Insert_heroInfo (herolist *object.Hero_list) error {

	//循环添加
	for i := 0; i < len(herolist.Hero); i++ {
		_, err := Db.Exec("insert into hero_briefInfo values(?,?,?)",
			//fmt.Println("哈哈哈",herolist.Hero[0].HeroId)
			herolist.Hero[i].HeroId,
			herolist.Hero[i].Name,
			herolist.Hero[i].Title,

		)

		if err != nil {
			return err
		}
	}

	return nil
}
//添加英雄详细
func Insert_detailsInfo (detailsInfo *object.Hero_list) error {

	//循环添加
	for i := 0; i < len(detailsInfo.Hero); i++ {
		_, err := Db.Exec("insert into hero_detailsInfo (hero_id, hero_engName, hero_Roles, " +
			"hero_IsWeekFree, hero_Attack, hero_Defense, hero_Magic, hero_Difficulty, hero_SelectAudio, " +
			"hero_BanAudio, hero_IsARAMweekfree, hero_Ispermanentweekfree, hero_ChangeLabel) " +
			"values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
			detailsInfo.Hero[i].HeroId,
			detailsInfo.Hero[i].Alias,
			detailsInfo.Hero[i].Roles[0],
			detailsInfo.Hero[i].IsWeekFree,
			detailsInfo.Hero[i].Attack,
			detailsInfo.Hero[i].Defense,
			detailsInfo.Hero[i].Magic,
			detailsInfo.Hero[i].Difficulty,
			detailsInfo.Hero[i].SelectAudio,
			detailsInfo.Hero[i].BanAudio,
			detailsInfo.Hero[i].IsARAMweekfree,
			detailsInfo.Hero[i].Ispermanentweekfree,
			detailsInfo.Hero[i].ChangeLabel,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

//添加英雄背景故事
func Insert_backstroy (HeroDetailedSlice *[]object.HeroDetailed) (err error) {
	//循环添加

	//fmt.Println((*HeroDetailedSlice)[0].Hero.ShortBio)
	for i := 0; i < len(*HeroDetailedSlice); i++ {
		_, err := Db.Exec("insert into hero_backstory (hero_id, hero_backStory) values(?,?)",
			(*HeroDetailedSlice)[i].Hero.HeroId,
			(*HeroDetailedSlice)[i].Hero.ShortBio,
		)

		if err != nil {
			return err
		}
	}
	return
}

//添加英雄技能信息
func Insert_skili (HeroDetailedSlice *[]object.HeroDetailed) (err error) {
	//循环添加

	//fmt.Println((*HeroDetailedSlice)[0].Hero.ShortBio)
	for i := 0; i < len(*HeroDetailedSlice); i++ {
		for j := 0; j < len((*HeroDetailedSlice)[i].Spells); j ++ {
			_, err := Db.Exec("insert into hero_skilintro (hero_id, hero_skillName, hero_skilPA_AC, hero_skillIntro) values(?,?,?,?)",
				(*HeroDetailedSlice)[i].Spells[j].HeroId,
				(*HeroDetailedSlice)[i].Spells[j].Name,
				(*HeroDetailedSlice)[i].Spells[j].SpellKey,
				(*HeroDetailedSlice)[i].Spells[j].Description,

			)

			if err != nil {
				return err
			}
		}

	}
	return
}

//添加英雄皮肤信息
func Insert_skin (HeroDetailedSlice *[]object.HeroDetailed) (err error) {
	//循环添加

	//fmt.Println((*HeroDetailedSlice)[0].Hero.ShortBio)
	for i := 0; i < len(*HeroDetailedSlice); i++ {
		for j := 0; j < len((*HeroDetailedSlice)[i].Skins); j ++ {
			_, err := Db.Exec("insert into hero_Skins (hero_id, skin_id, skin_name, emblem_name, skin_description, skin_mainImg, skin_iconImg, skin_loadingImg, skin_videoImg, skin_sourceImg) values(?,?,?,?,?,?,?,?,?,?)",
				(*HeroDetailedSlice)[i].Skins[j].HeroId,
				(*HeroDetailedSlice)[i].Skins[j].SkinId,
				(*HeroDetailedSlice)[i].Skins[j].Name,
				(*HeroDetailedSlice)[i].Skins[j].EmblemsName,
				(*HeroDetailedSlice)[i].Skins[j].Description,
				(*HeroDetailedSlice)[i].Skins[j].MainImg,
				(*HeroDetailedSlice)[i].Skins[j].IconImg,
				(*HeroDetailedSlice)[i].Skins[j].LoadingImg,
				(*HeroDetailedSlice)[i].Skins[j].VideoImg,
				(*HeroDetailedSlice)[i].Skins[j].SourceImg,
			)

			if err != nil {
				return err
			}
		}

	}
	return
}