package main

import (
	"web/fetcher"
	"log"
)

func main_1() {
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


	/*user:=new(fetcher.User)
	user.LoadConfig("")
	log.Println(user.Config)*/




	//测试反射
	//var x int =1
	//log.Println("Type:",reflect.TypeOf(x))
	//fmt.Println("Kind:  ", reflect.ValueOf(x).Kind())
	//fmt.Println("Kind is Int? ", reflect.ValueOf(x).Kind() == reflect.Int)

	//fmt.Println(reflect.ValueOf(x).Interface().(int))
	//fmt.Println(reflect.ValueOf(&x).CanSet())

	/*var x float64 = 3.4
	p := reflect.ValueOf(&x) // 获取x的地址
	fmt.Println("settability of p: ", p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v: ", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)*/


	/*type T struct {
		A int
		B string
	}

	t := T{12, "skidoo"}

	s:=reflect.ValueOf(&t).Elem()
	typeOfT:=s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}*/
	/*type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))*/


	res:=decode("https://account.aliyun.com/login/login_aliyun.htm?st=1QlGkotSZAjRaMGbljDnivg&nCode=&rType=&params=%7B%22site%22%3A%226%22%2C%22umidToken%22%3A%22Y8afbabdec7e1ef24e8e4f15fa8d6bc38%22%2C%22ru%22%3A%22https%3A%2F%2Fyq.aliyun.com%2Fattachment%2Fdownload%2F%3Fspm%3D5176.100239.0.0.sgdqRI%26id%3D498%26do%3Dlogin%22%2C%22ft%22%3A%22yqclub%22%7D")
	log.Println(res)
	//预约抢车程序
	//fetcher.Km2Cars(fetcher.Daolu,"2017/03/12",fetcher.Night,false)
}


func decode(v string) string{
	if v!=""{
		return fetcher.Decode(v)
	}
	return ""
}
