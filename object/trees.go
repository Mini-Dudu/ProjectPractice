package object

type Trees struct {
	Header string
	Tags []string
}
//SELECT hero_engName,hero_Roles,hero_IsWeekFree,hero_Attack,hero_Defense,hero_Magic,hero_Difficulty,hero_SelectAudio,hero_BanAudio FROM hero_detailsinfo;
//select h.hero_id,h.hero_alias,h.hero_name FROM hero_briefinfo h;
//(SELECT H.hero_engName,H.hero_Roles,H.hero_IsWeekFree,H.hero_Attack,H.hero_Defense,H.hero_Magic,H.hero_Difficulty,H.hero_SelectAudio,H.hero_BanAudio FROM hero_detailsinfo H)
//select * FROM hero_briefinfo h,
//(SELECT * FROM hero_detailsinfo) H WHERE h.hero_id=H.hero_id;

/*

测试：
SELECT h.hero_id,h.hero_alias,h.hero_name,H.hero_engName,H.hero_Roles,H.hero_IsWeekFree,H.hero_Attack,H.hero_Defense,H.hero_Magic,H.hero_Difficulty,H.hero_SelectAudio,H.hero_BanAudio
FROM hero_briefinfo h,(SELECT * FROM hero_detailsinfo) H
where h.hero_id=H.hero_id;

SELECT s.hero_id, s.skin_name, s.skin_mainImg, bj.hero_backStory
FROM (SELECT hero_id, skin_name, skin_mainImg FROM hero_skins WHERE emblem_name="base") s,hero_backStory bj WHERE s.hero_id=bj.hero_id;


有效：
SELECT h.hero_id,h.hero_alias,h.hero_name," +
		"H.hero_engName,H.hero_IsWeekFree,H.hero_Attack,H.hero_Defense,H.hero_Magic,H.hero_Difficulty,H.hero_SelectAudio," +
		"H.hero_BanAudio FROM hero_briefinfo h,(SELECT * FROM hero_detailsinfo) H where h.hero_id=H.hero_idDetails;




 */