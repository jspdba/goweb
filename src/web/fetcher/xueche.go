//公交驾校自动约车程序，已完成
package fetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/juju/persistent-cookiejar"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"path"
)

const (
	//userAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36"
	userAgent      = "android_gongjiao;v2.0.0;"
	connection     = "keep-alive"
	pragma         = "no-cache"
	xhr            = "XMLHttpRequest"
	acceptEncoding = ""
)

var (
	logger = Logger{Enabled: true}
	host   = "api.xuechebu.com"
	//1=道路训练 6=倒车入库
	trainType = "1"
	xxzh      = "62153539"
	jlcbh     = ""

	osversion  = "5.0.2"
	ossdk      = "21"
	appversion = "2.0.0"
	version    = "2.0.0"
	ipaddress  = "172.19.190.2"

	AuthError = errors.New("验证返回错误")

	//自动下单
	Auto = true

)

type MyTransport struct {
	tr        http.RoundTripper
	BeforeReq func(req *http.Request)
	AfterReq  func(resp *http.Response, req *http.Request)
}

func MyNewTransport(tr http.RoundTripper) *MyTransport {
	t := &MyTransport{}
	if tr == nil {
		tr = http.DefaultTransport
	}
	t.tr = tr
	return t
}

func (t *MyTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t.BeforeReq(req)
	resp, err = t.tr.RoundTrip(req)
	if err != nil {
		return
	}
	t.AfterReq(resp, req)
	return
}

type User struct {
	host   string
	Client *http.Client
	Config *GongjiaoConf
}

func (this *User) Init() {
	cookieJar, _ := cookiejar.New(nil)
	tr := MyNewTransport(nil)
	tr.AfterReq = func(resp *http.Response, req *http.Request) {
	}
	tr.BeforeReq = func(req *http.Request) {
	}

	if this.Client == nil {
		this.Client = &http.Client{
			Transport: tr,
			Jar:       cookieJar,
		}
	}
	this.LoadConfig("")
}

func (this *User) SaveCookie() {
	this.Client.Jar.(*cookiejar.Jar).Save()
}

func (this *User) Login() bool {
	loginUrl       := "http://api.xuechebu.com" + "/usercenter/userinfo/login?osversion=5.0.2&ossdk=21&passwordmd5="+this.Config.PasswordMd5+"&imei="+this.Config.Imei+"&username="+this.Config.Imei+"&appversion=2.0.0&version=2.0.0&ipaddress=172.19.190.2&os=an"
	resp, err := this.Get(loginUrl)
	if err != nil {
		logger.Error("登录失败：" + err.Error())
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("登录失败，StatusCode = %d", resp.StatusCode))
		return false
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取响应内容失败：%s", err.Error())
		return false
	}
	this.SaveCookie()
	logger.Info("相应内容=%s", content)
	return true
}

func newHTTPHeaders(isXhr bool) http.Header {
	headers := make(http.Header)
	headers.Set("Accept", "*/*")
	headers.Set("User-Agent", userAgent)
	headers.Set("Host", host)
	headers.Set("Connection", connection)
	headers.Set("Accept-Encoding", acceptEncoding)
	//headers.Set("Origin", "http://www.zhihu.com")
	//headers.Set("Pragma", pragma)
	if isXhr {
		headers.Set("X-Requested-With", xhr)
	}
	return headers
}

// Get 发起一个 GET 请求，自动处理 cookies
func (this *User) Get(url string) (*http.Response, error) {
	logger.Info("GET %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("NewRequest failed with URL: %s", url)
		return nil, err
	}

	req.Header = newHTTPHeaders(false)
	return this.Client.Do(req)
}

// Post 发起一个 POST 请求，自动处理 cookies
func (this *User) Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	logger.Info("POST %s, %s", url, bodyType)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	headers := newHTTPHeaders(false)
	headers.Set("Content-Type", bodyType)
	req.Header = headers
	return this.Client.Do(req)
}

//学车记录
func (this *User) Recode() {
	host = "gongjiao.xuechebu.com"
	path := "/KM2/GetXyXlZtInfo?xybh=5100908488&osversion=5.0.2&ossdk=21&imei="+this.Config.Imei+"&appversion=2.0.0&version=2.0.0&ipaddress=172.19.190.2&os=an"
	resp, err := this.Get("http://" + host + path)

	if err != nil {
		logger.Error("学车记录地址不存在：" + err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("查询学车记录失败，StatusCode = %d", resp.StatusCode))
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取响应内容失败：%s", err.Error())
		return
	}

	logger.Info("相应内容=%s", content)
	return
}

//科目2可约时间列表
func (this *User) Km2TimeSectionList() (timeSection *TimeSection) {
	host = "gongjiao.xuechebu.com"
	path := "/KM2/ClYyTimeSectionUIQuery2?trainType=" + trainType + "&osversion=" + osversion + "&ossdk=" + ossdk + "&imei=" + this.Config.Imei + "&xxzh=" + xxzh + "&appversion=" + appversion + "&version=" + version + "&jlcbh=" + jlcbh + "&ipaddress=" + ipaddress + "&os=an"
	resp, err := this.Get("http://" + host + path)

	if err != nil {
		logger.Error("预约列表地址不存在：" + err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("查询预约列表失败，StatusCode = %d", resp.StatusCode))
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取响应内容失败：%s", err.Error())
		return
	}
	logger.Info("返回内容=%s", content)

	var ts TimeSection
	err = json.Unmarshal(content, &ts)
	if err != nil {
		err = errors.New("unmarshal fail: " + string(content) + ", " + err.Error())
		return
	}
	timeSection = &ts
	return timeSection
}

//科目2车列表
func (this *User) Km2CarsList() (timeSection *TimeSection) {
	resp, err := this.Get(this.GetUrl("km2cars"))

	if err != nil {
		logger.Error("预约列表地址不存在：" + err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("查询预约列表失败，StatusCode = %d", resp.StatusCode))
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取响应内容失败：%s", err.Error())
		return
	}
	logger.Info("返回内容=%s", content)

	err = json.Unmarshal(content, timeSection)
	if err != nil {
		err = errors.New("unmarshal fail: " + string(content) + ", " + err.Error())
	}
	return
}

func CheckResp(resp *http.Response, err error) (err1 error, s []byte) {
	if err != nil {
		logger.Error("Get Error：" + err.Error())
		return err, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("State Error，StatusCode = %d", resp.StatusCode))
		return AuthError, nil
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取响应内容失败：%s", err.Error())
		return err, nil
	}
	logger.Info("相应内容=%s", content)
	return nil, content
}

func (this *User) getModuleList() {
	resp, err := this.Get(this.GetUrl("moduleList"))
	CheckResp(resp, err)
}

func (this *User) getBindDevice() {
	resp, err := this.Get(this.GetUrl("bindDevice"))
	CheckResp(resp, err)
}

func (this *User) setbadingstuinfo() {
	resp, err := this.Get(this.GetUrl("setbadingstuinfo"))
	CheckResp(resp, err)
	this.SaveCookie()
}

//返回
func (this *User) getCars(uidata UIData, day string) (err1 error, cars *Car) {
	resp, err := this.Get(this.BuildParamUrl(uidata, day))
	if err != nil {
		return err, nil
	}
	err, content := CheckResp(resp, err)
	if err != nil {
		return err, nil
	}
	var c Car
	err = json.Unmarshal(content, &c)
	if err != nil {
		err = errors.New("unmarshal fail: " + string(content) + ", " + err.Error())
		return err, nil
	}
	cars = &c
	return nil, cars
}

//创建带参数的请求地址
func(this *User) BuildParamUrl(uidata UIData, day string) (p string) {
	//p := "http://gongjiao.xuechebu.com" + "/KM2/ClYyCars2?osversion="+osversion+"&ossdk="+ossdk+"&filters[yyrq]=2017年03月17日&imei="+imei+"&xxzh="+xxzh+"&appversion="+appversion+"&filters[xnsd]=15&version=2.0.0&filters[trainType]="+trainType+"&ipaddress="+ipaddress+"&filters[jlcbh]=&filters[xxzh]=62153539&os=an"
	p = "http://gongjiao.xuechebu.com" + "/KM2/ClYyCars2?"
	p += EncodeSortedMap(&SortedMap{
		Keys:       []string{"osversion", "ossdk", "yyrq", "imei", "xxzh", "appversion", "xnsd", "version", "trainType", "ipaddress", "jlcbh", "filters%5Bxxzh%5D", "os"},
		FilterKeys: []string{"yyrq", "xnsd", "trainType", "jlcbh"},
		Values: map[string]interface{}{
			"osversion":         osversion,
			"ossdk":             ossdk,
			"yyrq":              ParseDate(day),
			"imei":             this.Config.Imei ,
			"xxzh":              xxzh,
			"appversion":        appversion,
			"xnsd":              uidata.Xnsd,
			"version":           version,
			"trainType":         trainType,
			"ipaddress":         ipaddress,
			"jlcbh":             jlcbh,
			"filters%5Bxxzh%5D": xxzh,
			"os":                "an",
		},
	})
	return
}

func (this *User) GetUrl(stype string) string {
	getModuleListBySchoolCode := "http://api.xuechebu.com" + "/school/GetModuleListBySchoolCode?osversion=5.0.2&ossdk=21&imei="+this.Config.Imei+"&appversion=2.0.0&schoolcode=340800035&version=2.0.0&ipaddress=172.19.190.2&os=an"
	bindDevice := "http://xcbapi.xuechebu.com" + "/jpushapi/BindDevice?osversion=5.0.2&ossdk=21&imei="+this.Config.Imei+"&appversion=2.0.0&version=2.0.0&ipaddress=172.19.190.2&devicetype=3&registrationId=160a3797c8096eda063&os=an"
	setbadingstuinfo := "http://gongjiao.xuechebu.com" + "/Student/setbadingstuinfo?id_type=1&xybh=5100908488&jgid=124001&password="+this.Config.Password

	switch stype {
	case "moduleList":
		this.host = "api.xuechebu.com"
		return getModuleListBySchoolCode
	case "bindDevice":
		this.host = "xcbapi.xuechebu.com"
		return bindDevice
	case "setbadingstuinfo":
		this.host = "gongjiao.xuechebu.com"
		return setbadingstuinfo
	case "km2cars":
		this.host = "gongjiao.xuechebu.com"
		//km2Cars := "http://gongjiao.xuechebu.com" + "/KM2/ClYyCars2?osversion="+osversion+"&ossdk="+ossdk+"&filters[yyrq]=2017年03月17日&imei="+imei+"&xxzh="+xxzh+"&appversion="+appversion+"&filters[xnsd]=15&version=2.0.0&filters[trainType]="+trainType+"&ipaddress="+ipaddress+"&filters[jlcbh]=&filters[xxzh]=62153539&os=an"
		km2Cars := "http://gongjiao.xuechebu.com" + "/KM2/ClYyCars2?osversion=" + osversion + "&ossdk=" + ossdk + "&imei=" + this.Config.Imei + "&xxzh=" + xxzh + "&appversion=" + appversion + "&version=" + version + "&ipaddress=" + ipaddress + "&os=an"
		return km2Cars
	default:
		this.host = "api.xuechebu.com"
		loginUrl       := "http://api.xuechebu.com" + "/usercenter/userinfo/login?osversion=5.0.2&ossdk=21&passwordmd5="+this.Config.PasswordMd5+"&imei="+this.Config.Imei+"&username="+this.Config.Imei+"&appversion=2.0.0&version=2.0.0&ipaddress=172.19.190.2&os=an"
		return loginUrl
	}
}

func EncodeFilterMap(m map[string]interface{}) (content string) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			content += url.QueryEscape("filters["+k+"]="+v.(string)) + "&"
		}
		content = content[:len(content)-1]
	}
	return
}

func EncodeFilterWithSortedMap(m *SortedMap) (content string) {
	if m != nil && len(m.Keys) > 0 {
		for _, k := range m.Keys {
			content += url.QueryEscape("filters["+strings.ToLower(k)+"]="+m.Values[k].(string)) + "&"
		}
		content = content[:len(content)-1]
	}
	return
}

func EncodeSortedMap(m *SortedMap) (content string) {
	if m != nil && len(m.Keys) > 0 {
		for _, k := range m.Keys {
			if m.IsFilterKey(k) {
				content += url.QueryEscape("filters["+k+"]") + "=" + url.QueryEscape(m.Values[k].(string)) + "&"
			} else {
				content += (k + "=" + m.Values[k].(string)) + "&"
			}
		}
		content = content[:len(content)-1]
	}
	return
}

func EncodeFilter(key, value string) (content string) {
	return url.QueryEscape("filters[" + key + "]=" + value)
}

//读取文件
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

//获取时间选择列表
func GetTimeSection() (ts *TimeSection) {
	wd, _ := os.Getwd()
	d, err := ReadAll(filepath.Join(wd, "src\\web", "xueche.yueche.json"))
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		json.Unmarshal(d, &ts)
	}
	return ts
}
func FindUseTime(ts *TimeSection, dateTimeSections ...*DateTimeSection) (uidata []UIData) {
	if ts == nil {
		log.Println("nil timesection")
		return
	}
	if ts.Code == 0 {
		uidata = []UIData{}
		for _, v := range ts.Data.UIDatas {
			if v.SL == 0 {
				//如果不可约继续
				continue
			}

			for _, dateTimeSection := range dateTimeSections {
				if strings.Index(v.Yyrq, dateTimeSection.DateSection) > -1 {
					if dateTimeSection.TimeSection == nil || len(dateTimeSection.TimeSection) == 0 {
						uidata = append(uidata, v)
						continue
					}
					for _, timeItem := range dateTimeSection.TimeSection {
						if strings.Index(v.Xnsd, timeItem) > -1 {
							uidata = append(uidata, v)
						}
					}
				}
			}

		}
		return
	}
	log.Println("无法解析时间列表", ts.Message)
	return
}

//查看是星期几
func WhichDay(timeSection *TimeSection, uidata *UIData) (today string) {
	for _, yyrq := range timeSection.Data.YyrqList {
		if yyrq.Yyrq == uidata.Yyrq {
			today = yyrq.DisplayWeek + WhichTime(uidata.Xnsd)
			return
		}
	}
	return
}

func WhichTime(time string) string {
	switch time {
	case Am:
		return "晚上"
	case Pm:
		return "下午"
	case Night:
		return "晚上"
	default:
		return ""
	}
}
func Run() {
	user := new(User)
	user.Init()
	//user.Login()
	//user.Recode()

	//user.getModuleList()
	//user.getBindDevice()
	//如果登录失效这里需要重新绑定一次
	//user.setbadingstuinfo()
	//科目2时间列表
	user.Km2TimeSectionList()
	//科目2约车列表
	user.Km2CarsList()
}

//查询约车信息 "2017/03/11"
func Km2Cars(typ, day, section string, bindFirst bool) {
	user := new(User)
	user.Init()

	if bindFirst {
		//如果登录失效这里需要重新绑定一次
		user.setbadingstuinfo()
	}

	switch typ {
	case Daolu:
		trainType = "1"
	case Daocheruku:
		trainType = "6"
	default:
		trainType = "1"
	}

	for {
		//科目2时间列表
		timeSection := user.Km2TimeSectionList()

		//查看此时间是否可用--开始
		dateTimeSection := &DateTimeSection{
			DateSection: day,
			TimeSection: []string{
				section,
			},
		}

		uiDatas := FindUseTime(timeSection, dateTimeSection)
		//查看此时间是否可用--结束

		//发起请求，预约此时间
		for _, info := range uiDatas {
			log.Println(WhichDay(timeSection, &info))
			err, cars := user.getCars(info, day)
			if err != nil {
				logger.Error("获取约车列表信息失败:%s", err.Error())
			}
			ok := user.DoYueChe(cars, day, section)

			if ok {
				return
			}
		}

		time.Sleep(60 * time.Second)
	}

}

func SwitchInput() string {
	var selected string
	fmt.Print(color.CyanString("请选择教练："))
	fmt.Scanf("%s", &selected)
	return selected
}

//解码
func Decode(u string) string {
	r, _ := url.QueryUnescape(u)
	return r
}

//预约教练
func (this *User) RequestYueche(car Info, day, section string) {
	host = "gongjiao.xuechebu.com"
	path := "/KM2/ClYyAddByMutil?"

	params := url.Values{}
	params.Add("trainType", trainType)
	params.Add("osversion", osversion)
	params.Add("ossdk", ossdk)
	params.Add("imei", this.Config.Imei)
	params.Add("xxzh", xxzh)
	params.Add("appversion", appversion)
	params.Add("isJcsdYyMode", "5")
	params.Add("version", version)
	params.Add("jlcbh", jlcbh)
	params.Add("ipaddress", ipaddress)
	params.Add("os", "an")
	params.Add("params", strings.Join([]string{
		car.CNBH, ParseDate(day), section, "",
	}, "."))
	//params.Add("params","30051.2017年03月11日.58.")
	uri := "http://" + host + path + params.Encode()

	resp, err := this.Get(uri)
	if err != nil {
		logger.Error("%s", err.Error())
		return
	}
	err, _ = CheckResp(resp, err)
	if err != nil {
		logger.Error("%s", err.Error())
		return
	}
}

//循环约车
func (this *User) DoYueChe(cars *Car, day, section string) (ok bool) {
	if cars == nil || cars.Data.Total == 0 {
		logger.Info("没有查询到约车列表信息")
		return
	}
	if cars.Code != 0 {
		logger.Info("获取失败:%s" + cars.Message)
		return
	}
	for {
		time.Sleep(200 * time.Microsecond)
		for i, k := range cars.Data.Result {
			logger.Info("序号:%d),教练：%s", i, k.JLYXM)
		}
		logger.Info("q=退出")

		var in = "0"
		if Auto == false {
			in = SwitchInput()
		}

		if "q" == strings.ToLower(in) {
			return
		}

		i, err := strconv.Atoi(in)
		if err != nil {
			continue
		}
		if i >= 0 && i < len(cars.Data.Result) {
			logger.Info("选择:%s,教练姓名：%s,%s,%s,%s", in, cars.Data.Result[i].JLYXM, cars.Data.Result[i].CNBH, cars.Data.Result[i].JLCBH, cars.Data.Result[i].YT)
			this.RequestYueche(cars.Data.Result[i], day, section)
			in = "q"
			return true
		}
	}
}


func (this *User) LoadConfig(cfg string) {
	if cfg==""{
		dir,e:=os.Getwd()
		if e!=nil{
			return
		}
		cfg=path.Join(dir,"src/web/gongjiaojiaxiao.json")
	}

	fd, err := os.Open(cfg)
	if err != nil {
		panic("无法打开配置文件 config.json: " + err.Error())
	}
	defer fd.Close()

	config := new(GongjiaoConf)
	err = json.NewDecoder(fd).Decode(&config)
	if err != nil {
		panic("解析配置文件出错: " + err.Error())
	}
	//读取配置文件，且放到user中
	this.Config = config
}