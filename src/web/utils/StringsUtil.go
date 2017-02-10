package utils

import (
	"strings"
	"unicode"
	"strconv"
	"github.com/henrylee2cn/mahonia"
)


//翻转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//翻转slice
func ReverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//提取包含的字符
func ContainsKeys(str string,keys []string) []string{
	ks:=[]string{}
	for _,v:=range str{
		if IsContainsAnyKey(string(v),keys){
			ks=append(ks,string(v))
		}
	}
	return ks
}

func IsContainsAnyKey(str string,arr []string) bool {
	return strings.ContainsAny(str, strings.Join(arr,"&"))
}

func ToDigits(str string,keyValueMap map[string]string) string{
	for key:=range keyValueMap{
		str=strings.Replace(str,key,keyValueMap[key],-1);
	}
	return str
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
func IsDigits(str string) bool{
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
type KeyValues struct {
	Keys []string
	Values map[string] string
}

func Convert2Digit(s string) string{
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
	s=ToDigits(s,keyValues.Values)
	return s
}

//补零算法
func addZero(str string,unitKeyValues *KeyValues) string{
	containsKeys:=ContainsKeys(str,unitKeyValues.Keys)
	containsKeys=ReverseSlice(containsKeys)

	lastMaxLen := 0
	for _,key:=range containsKeys{
		if index:=strings.LastIndex(str,key); index>-1{
			s1:=str[:index]
			s2:=str[index+len(key):]
			no,_:=strconv.Atoi(unitKeyValues.Values[key])
			if s2==""{
				s2 += strings.Repeat("零",no)
			}else{
				zeroAddedCount :=no-len([]rune(s2))
				if no < lastMaxLen{
					zeroAddedCount =lastMaxLen+no-len([]rune(s2))
				}
				if zeroAddedCount>0{
					s2 = strings.Repeat("零",zeroAddedCount)+s2
				}
			}
			if s1==""{
				s1="一"
			}
			str=s1+s2
			if no>=lastMaxLen{
				lastMaxLen=no
			}
		}
	}

	return str
}
//找出数字
func FindDigit(str string) string{

	if strings.LastIndex(str,"章")>-1{
		str = str[0:strings.LastIndex(str,"章")]
	}
	str = strings.Replace(str," ","",-1)
	str = strings.Replace(str,"第","",-1)
	return str
}

func ParseToGBK(s string) string{
	enc:=mahonia.NewEncoder("gbk")
	return enc.ConvertString(s)
}
