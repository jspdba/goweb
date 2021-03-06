package test

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"regexp"
	"strconv"
	"web/utils"
	"web/models"
	"sync"
)

func TestSelector(t *testing.T) {
	url:="http://www.biquge.tw/0_671/"
	selector:="#list > dl > dd"
	chapters:= GetUrlInfo(url,selector)
	GetChapterContent(chapters)
}

func GetChapterContent(chapters  []*models.Chapter)  {
	threadsCount:=100
	lenChapters:=len(chapters)
	log.Print(lenChapters)

	itemCount:=lenChapters/threadsCount
	var wg sync.WaitGroup

	if (threadsCount>=lenChapters){
		itemCount=1
		for i:=0;i<lenChapters;i++{
			wg.Add(1)
			go DownloadOne(&wg,chapters[i])
		}
	}else{
		for start,i:=0,0;i<threadsCount;i++{
			end:=start+itemCount
			log.Println(start,end,lenChapters)
			wg.Add(1)
			if end==lenChapters-1{
				log.Print(start,":")
				go Download(&wg,chapters[start:])
			}else{
				log.Println(start,":",end)
				go Download(&wg,chapters[start:end])
			}
			start=end
		}
	}
	wg.Wait()
}

func Download(wg *sync.WaitGroup,chapters []*models.Chapter){
	defer wg.Done()
	for _,chapter:=range chapters{
		if chapter.Url!=""{
			//log.Print(chapter.Url)
			selector:="#content"
			content:=GetContent(chapter.Url,selector)
			//log.Print("进度=",(i+1)/len(chapters),content)
			chapter.Content=content
			//log.Print("下载内容=",content)
		}
	}
}
func DownloadOne(wg *sync.WaitGroup,chapter *models.Chapter){
	defer wg.Done()
	if chapter.Url!=""{
		//log.Print(chapter.Url)
		selector:="#content"
		content:=GetContent(chapter.Url,selector)
		//log.Print("进度=",(i+1)/len(chapters),content)
		chapter.Content=content
		//log.Print("下载内容=",content)
	}
}

func Test_fetchContent(t *testing.T) {
	url:="http://www.biquge.tw/0_671/4770943.html"
	selector:="#content"
	log.Print(GetContent(url,selector))
}

func GetContent(url string, selector string) string{
	doc, err := goquery.NewDocument(url)
	if err == nil {
		if doc != nil {
			content:= doc.Find(selector).Text()
			return content
		}
	}else{
		log.Print(err)
	}
	return ""
}

func GetUrlInfo(url string,selector string) []*models.Chapter{
	doc, err := goquery.NewDocument(url)

	chapters := []*models.Chapter{}
	if err == nil && doc!=nil {
		doc.Find(selector).EachWithBreak(func(i int, contentSelection *goquery.Selection) bool{
			title := contentSelection.Find("a").Text()
			//log.Println("第", i + 1, "个帖子的标题：", title)
			title=utils.Convert2Digit(title)
			index:=getIndex(title)
			//log.Println("第", i + 1, "个帖子的章节：", index)

			href, _ := contentSelection.Find("a").Attr("href")
			if !strings.HasPrefix(href, "http") {
				href = getHost(url) + href
				log.Println("第", i + 1, "个帖子的url：", href)
			}
			chapter:=&models.Chapter{
				Title:title,
				Index:index,
				Url:href,
			}
			chapters=append(chapters,chapter)

			if len(chapters)==200{
				return false;
			}
			return true
		})
	}
	return chapters
}

func getHost(url string) string{
	url=strings.Replace(url,"http://","", -1)
	url=url[0:strings.Index(url,"/")]
	return "http://"+url
}

func getIndex(s string) int{

	valid ,err:= regexp.Compile("[0-9]{1,5}")
	if !valid.MatchString(s){
		return -1
	}
	if err==nil{
		data := valid.FindAllStringSubmatch(s, -1)
		if len(data)>=0{
			if num,err :=strconv.Atoi(strings.Join(data[0],"")); err==nil{
				return num
			}
		}
	}
	return -1
}