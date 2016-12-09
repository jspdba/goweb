package test

import (
	"testing"
	"io/ioutil"
	"strings"
	"text/template"
	"os"
)

func Test_genFrame(t *testing.T) {
	dir:="D:\\zhongliang\\repos\\mobile\\html5\\prime\\prime-static\\dingqigou"
	toFile:=dir+"\\"+"right.html"
	if files,err:=ListDir(dir,".html"); err==nil{
		parseTemplate(files,toFile)
	}
}

func parseTemplate(path []string,target string){
	if tmpl, err := template.ParseFiles("D:\\zhongliang\\go\\goweb\\src\\web\\tests\\frame.tpl");err==nil{
		if f,err:=os.OpenFile(target,os.O_CREATE,0664);err == nil {
			err = tmpl.Execute(f, path)
			if err != nil {
				panic(err)
			}
		}else{
			panic(err)
		}
	}else{
		panic(err)
	}
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			if strings.HasPrefix(fi.Name(),"frame."){
				continue
			}
			if strings.HasPrefix(fi.Name(),"left."){
				continue
			}
			if strings.HasPrefix(fi.Name(),"right."){
				continue
			}
			files = append(files, fi.Name())
		}
	}
	return files, nil
}