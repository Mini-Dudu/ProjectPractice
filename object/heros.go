package object

//英雄简略信息
type HeroInfo struct {
	HeroId string
	Name string
	Alias string
	Title string
	Roles []string
	IsWeekFree string
	Attack string
	Defense string
	Magic string
	Difficulty string
	SelectAudio string
	BanAudio string
	IsARAMweekfree string
	Ispermanentweekfree string
	ChangeLabel string

}

//英雄基本信息
type Hero struct {
	HeroId string
	Name string
	Alias string
	Title string
	Roles []string
	ShortBio string
	Camp string
	Attack string
	Defense string
	Magic string
	Difficulty string
	Hp string
	Hpperlevel string
	Mp string
	Mpperlevel string
	Movespeed string
	Armor string
	Armorperlevel string
	Spellblock string
	Spellblockperlevel string
	Attackrange string
	Hpregen string
	Hpregenperlevel string
	Mpregen string
	Mpregenperlevel string
	Crit string
	Critperlevel string
	Attackdamage string
	Attackdamageperlevel string
	Attackspeed string
	Attackspeedperlevel string
	Allytips []string
	Enemytips []string
	HeroVideoPath string
	IsWeekFree string
	DamageType string
	Style string
	DifficultyL string
	Damage string
	Durability string
	CrowdControl string
	Mobility string
	Utility string
	SelectAudio string
	BanAudio string
	ChangeLabel string
}

const(
	a =1
	b =2
	c =3
)
//英雄皮肤
type Skin struct {
	SkinId string
	HeroId string
	HeroName string
	HeroTitle string
	Name string
	Chromas string
	ChromasBelongId string
	IsBase string
	EmblemsName string
	Description string
	MainImg string
	IconImg string
	LoadingImg string
	VideoImg string
	SourceImg string
	VedioPath string
	SuitType string
	PublishTime string
	ChromaImg string
	ShortBio string
}

//英雄技能
type Spell struct {
	HeroId string
	SpellKey string
	Name string
	Description string
	AbilityIconPath string
	AbilityVideoPath string
	DynamicDescription string
	//Cost []string
	//Costburn string
	//Cooldown []string
	//Cooldownburn string
	//Range []string
}
type num struct {
	ID int
	NUM string
	AGE int


}

func mul(a,b int) int {
	var nn  = a * b
	return nn
}



//英雄详细信息
type HeroDetailed struct {
	Hero Hero
	Skins []Skin
	Spells []Spell
	Version string
	FileName string
	FileTime string


}
