package main

import (
	"web/fetcher"
	"log"
)

func main() {
	//fetcher.Run()
	//km2Cars          := "http://gongjiao.xuechebu.com" + "/KM2/ClYyCars2?osversion=5.0.2&ossdk=21&filters%5Byyrq%5D=2017%E5%B9%B403%E6%9C%8817%E6%97%A5&imei=***&xxzh=62153539&appversion=2.0.0&filters%5Bxnsd%5D=15&version=2.0.0&filters%5BtrainType%5D=1&ipaddress=172.19.190.2&filters%5Bjlcbh%5D=&filters%5Bxxzh%5D=62153539&os=an"
	//km2Cars          := "filters[yyrq]=2017年03月17日"
	//log.Println(EncodeFilter("yyrq","2017年03月17日"))

	/*m:=map[string]string{
		yyrq%5D%3D0001%E5%B9%B401%E6%9C%8801%E6%97%A5
		"yyrq":"2017年03月17日",
		"xnsd":"15",
		"trainType":"1",
		"jlcbh":"",
		"xxzh":"62153539",
	}

	log.Println(fetcher.EncodeFilterMap(m))*/
	/*dateTimeSection := &fetcher.DateTimeSection{
		DateSection:"2017/03/16",
		TimeSection:[]string{
			fetcher.Pm,
			fetcher.Night,
		},
	}
	timeSection := fetcher.GetTimeSection()
	uiDatas := fetcher.FindUseTime(timeSection,dateTimeSection)

	for _,info:= range uiDatas{
		log.Println(info)
		log.Println(fetcher.WhichDay(timeSection,&info))
	}*/
	//log.Println(fetcher.ParseDate("2018/03/15"))



	//log.Println(url.QueryEscape("filters[yyrq]"))


	//fmt.Print(url.QueryUnescape("/KM2/ClYyAddByMutil?trainType=6&osversion=5.0.2&ossdk=21&imei=***&xxzh=62153539&appversion=2.0.0&isJcsdYyMode=5&version=2.0.0&jlcbh=&ipaddress=172.19.190.2&os=an&params=30051.2017%E5%B9%B403%E6%9C%8811%E6%97%A5.58."))
	///KM2/ClYyAddByMutil?trainType=6&osversion=5.0.2&ossdk=21&imei=***&xxzh=62153539&appversion=2.0.0&isJcsdYyMode=5&version=2.0.0&jlcbh=&ipaddress=172.19.190.2&os=an&params=30051.2017年03月11日.58.
	//fmt.Println(fetcher.BuildParams())


	//fetcher.Km2Cars(fetcher.Daolu,"2017/03/12",fetcher.Night,false)
	user:=new(fetcher.User)
	user.LoadConfig("")
	log.Println(user.Config)
}


