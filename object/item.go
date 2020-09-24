package object

type Adhguhg struct {
	ItemId string
	Name string
	IconPath string
	Price string
	Description string
	Maps []string
	Plaintext string
	Sell string
	Total string
	//into []string			//有些数据为[]string，有些又为string这里舍弃这个数据....
	//from string			//有些数据为[]string，有些又为string这里舍弃这个数据....
	//SuitHeroId string		//大多数物品没有此数据，这里也进行舍弃
	//Tag string			//大多数物品没有此数据，这里也进行舍弃
	Types []string
}
