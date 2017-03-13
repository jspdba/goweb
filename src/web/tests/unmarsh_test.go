package test

import (
	"testing"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"web/fetcher"
)

func TestUnmarshal(t *testing.T) {
	d,err:=ReadAll("xueche.yueche.json")
	if err!=nil{
		log.Println(err.Error())
	}else{
		var timeSection fetcher.TimeSection
		json.Unmarshal(d,&timeSection)

		log.Println(timeSection.Code)
		log.Println(timeSection.Message)
		log.Println(timeSection)
	}
}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}