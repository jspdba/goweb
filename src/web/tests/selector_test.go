package test

import (
	"testing"
	"github.com/astaxie/beego"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"regexp"
	"strconv"
	"unicode"
)

func TestSelector(t *testing.T) {
	url:="http://www.biquge.tw/0_671/"
	selector:="#list > dl > dd"
	getUrlInfo(url,selector)
}
func getUrlInfo(url string,selector string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		beego.Error(err)
	}
	if doc != nil {
		doc.Find(selector).Each(func(i int, contentSelection *goquery.Selection) {
			title := contentSelection.Find("a").Text()
			log.Println("第", i + 1, "个帖子的标题：", title)
			log.Println("第", i + 1, "个帖子的章节：", getIndex(title))

			href, _ := contentSelection.Find("a").Attr("href")
			if !strings.HasPrefix(href, "http") {
				href = getHost(url) + href
				log.Println("第", i + 1, "个帖子的url：", href)

			}
		})
	}
}

func getHost(url string) string{
	url=strings.Replace(url,"http://","", -1)
	url=url[0:strings.Index(url,"/")]
	return "http://"+url
}

type KeyValues struct {
	Keys []string
	Values map[string] string
}

func Test_convert2Digit(t *testing.T) {
	s:="十万零一十"
	keyValues := &KeyValues{}
	keyValues.Keys = []string{"一","二","三","四","五","六","七","八","九","零"}
	keyValues.Values = map[string]string{
		"一":"1",
		"二":"2",
		"三":"3",
		"四":"4",
		"五":"5",
		"六":"6",
		"七":"7",
		"八":"8",
		"九":"9",
		"零":"0",
	}

	unitKeyValues:=&KeyValues{}
	unitKeyValues.Keys=[]string{"十","百","千","万"}
	unitKeyValues.Values=map[string]string{
		"十":"1",
		"百":"2",
		"千":"3",
		"万":"4",
	}
	s=addZero(s,unitKeyValues)
	log.Println(s)
}

//补零算法
func addZero(str string,unitKeyValues *KeyValues) string{
	//s:="十万零一千零十"
	for containsKey(str,unitKeyValues.Keys){
		for _,key:=range unitKeyValues.Keys{
			if index:=strings.LastIndex(str,key); index>-1{
				s1:=str[:index]
				s2:=str[index+len(key):]
				no,_:=strconv.Atoi(unitKeyValues.Values[key])
				if s2==""{
					s2 += strings.Repeat("零",no)
				}
				//应该补几个零
				str=s1+s2
			}
		}
	}
	return str
}

func containsKey(str string,arr []string) bool {
	return strings.ContainsAny(str, strings.Join(arr,"&"))
}
//截取字符位置
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
//判断是否是数字字符串
func isDigits(str string) bool{
	if str==""{
		return false
	}
	for _,r :=range str{
		if !unicode.IsDigit(r){
			return false
		}
	}
	return true
}

func toDigits(str string,keyValueMap map[string]string) string{
	for key:=range keyValueMap{
		str=strings.Replace(str,key,keyValueMap[key],-1);
	}
	return str
}

func getIndex(s string) int{
	var valid = regexp.MustCompile("[0-9]{1,5}")
	data := valid.FindAllStringSubmatch(s, -1)
	if len(data)>=0{
		if num,err :=strconv.Atoi(strings.Join(data[0],"")); err==nil{
			return num
		}
	}
	return -1
}