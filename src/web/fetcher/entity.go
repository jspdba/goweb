//返回结果的结构体
package fetcher

import (
	"time"
)

type ResultData struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Type    int         `json:"type"`
}

const (
	Am    = "812" //"07:00--11:30"
	Pm    = "15"  //"12:30--17:00"
	Night = "58"  //"17:30--21:00"

	Daolu      = "1"
	Daocheruku = "6"
)

//学车时间
type DateTimeSection struct {
	DateSection string
	TimeSection []string
}

//时间选择
type TimeSection struct {
	Data    Section `json:"data"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
	//Type    int     `json:"type"`
}

type Section struct {
	XnsdList []Xnsd
	YyrqList []Yyrq
	QsList   []interface{}
	UIDatas  []UIData
}

type Xnsd struct {
	Xnsd     string
	XnsdName string
}
type Yyrq struct {
	Yyrq        string
	DisplayWeek string
	DisplayYyrq string
}

type UIData struct {
	Yyrq       string
	YyrqXH     int
	Xnsd       string
	XnsdName   string
	QsName     string
	Qsid       string
	SL         int
	KS         int
	IsBpked    bool
	IsBpked_SK int
	IsCreate   bool
	YyClInfo   string
}

//选择车（教练）
type Car struct {
	Data    Pager  `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Pager struct {
	Result []Info
	Total  int
}
type Info struct {
	JLCBH string
	CNBH  string
	YT    string
	JLYXM string
}

type SortedMap struct{
	Keys []string
	FilterKeys []string
	Values map[string]interface{}
}
//增加 key
func(this *SortedMap) appendKey(keys ...string) {
	this.Keys=append(this.Keys,keys...)
}
//增加 FilterKey
func(this *SortedMap) appendFilterKey(keys ...string) {
	this.FilterKeys=append(this.FilterKeys,keys...)
}

//是否是 Filterkey
func(this *SortedMap) IsFilterKey(key string)(result bool){
	if this.FilterKeys==nil || len(this.FilterKeys)==0{
		return false
	}
	result = false
	for _,k:=range this.FilterKeys{
		if k==key{
			result = true
			return result
		}
	}
	return
}


func ParseDate(in string) string{
	tm2, _ := time.Parse("2006/01/02", in)
	return tm2.Format("2006年01月02日")
}

//配置文件读取
type GongjiaoConf struct {
	Username string		`json:"username"`
	Password string		`json:"password"`
	PasswordMd5 string	`json:"passwordMd5"`
	Imei string		`json:"imei"`
}