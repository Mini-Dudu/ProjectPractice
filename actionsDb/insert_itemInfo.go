package actionsDb

import "LOL/object"

func Insert_itemInfo (itemInfo *object.AllItem) error {
	//循环添加
	for i := 0; i < len(itemInfo.Items); i++ {
		var map_, type_ string

		if len(itemInfo.Items[i].Maps) != 0 {
			for j := 0; j < len(itemInfo.Items[i].Maps); j++ {
				map_ += itemInfo.Items[i].Maps[j] + "、"
			}
			//map_ = itemInfo.Items[i].Maps[0]
		}

		if len(itemInfo.Items[i].Types) != 0 {
			for j := 0; j < len(itemInfo.Items[i].Types); j++ {
				type_ += itemInfo.Items[i].Types[j] + "、"
			}
			//type_ = itemInfo.Items[i].Types[0]
		}

		_, err := Db.Exec("insert into item_Info values(?,?,?,?,?,?,?,?,?,?)",
			itemInfo.Items[i].ItemId,
			itemInfo.Items[i].Name,
			itemInfo.Items[i].IconPath,
			itemInfo.Items[i].Price,
			itemInfo.Items[i].Description,
			map_,
			itemInfo.Items[i].Plaintext,
			itemInfo.Items[i].Sell,
			itemInfo.Items[i].Total,
			type_,

		)
		if err != nil {
			return err
		}
	}

	return nil
}
