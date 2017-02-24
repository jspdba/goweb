#文章列表
[beego文档](https://beego.me/docs/mvc/controller/config.md)
##beego路由文档
[beego 路由文档](https://beego.me/docs/mvc/controller/router.md)
##git 设置代理|取消代理
    git config --global http.proxy 127.0.0.1:1080
    git config --global https.proxy 127.0.0.1:1080

    git config –global http.proxy http://user:password@10.167.32.133:8080
    git config –global http.proxy https://user:password@10.167.32.133:8080

#删除HTTP代理
    git config --system (或 --global 或 --local) --unset http.proxy
    git config --system (或 --global 或 --local) --unset https.proxy

    //git config --system --unset http.proxy
    //git config --system --unset https.proxy

    git config --global --unset http.proxy
    git config --global --unset https.proxy
##go 编码转换
    http://studygolang.com/articles/1712
##httplib
    https://beego.me/docs/module/httplib.md

##记录一个git错误
    ..\github.com\andybalholm\cascadia\parser.go:11:2: cannot find package "golang.org/x/net/html" in any of:
	C:\Go\src\golang.org\x\net\html (from $GOROOT)
	D:\zhongliang\go\goweb\src\golang.org\x\net\html (from $GOPATH)

    https://www.oschina.net/question/566882_212351
	从https://github.com/golang/net下载，然后把目录改成golang.org/x/net。然后，万事大吉。
    ps：有git的话可以直接 go get github.com/golang/net，没有的话自己手动下载放到src目录下即可
##go下载地址
    [go 下载地址](https://golang.org/dl/)
##go 安装
    [go 安装](https://golang.org/doc/install)
- tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
- vi /etc/profile
- export PATH=$PATH:/usr/local/go/bin
- 或者(如果是其它目录)
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOROOT/bin
- export GOPATH=/opt/go/webapp
##beego
###bee 工具的安装
- 您可以通过如下的方式安装 bee 工具：
- go get github.com/beego/bee   
    bee可执行文件默认存放在$GOPATH/bin里面，所以您需要把$GOPATH/bin添加到您的环境变量中，才可以进行下一步
## go 交叉编译
[参考](http://www.tuicool.com/articles/fyumIzn)
### 在Go根目录下的src目录，新建一个build.bat文件，并复制内容如下
    set CGO_ENABLED=0
    set GOROOT_BOOTSTRAP=C:/Go
    ::x86块
    set GOARCH=386
    set GOOS=windows
    call make.bat --no-clean
      
    set GOOS=linux
    call make.bat --no-clean
      
    set GOOS=freebsd
    call make.bat --no-clean
      
    set GOOS=darwin
    call make.bat --no-clean
    ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
      
    ::x64块
    set GOARCH=amd64
    set GOOS=linux
    call make.bat --no-clean
    ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
      
    ::arm块
    set GOARCH=arm
    set GOOS=linux
    call make.bat --no-clean
    ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
      
    set GOARCH=386
    set GOOS=windows
    go get github.com/nsf/gocode
    pause
###执行 build.bat
    完成后，在cmd命令行下依次执行：
    set GOOS=linux
    set GOPACH=amd64
    go build -o -x APPNAME main.go
    编译后的文件会出现在main.go相应的目录下。
#访问服务器
[web服务器](http://182.92.85.72:8888/link/edit)
#UI
[Flat ui-view](http://www.bootcss.com/p/flat-ui/)
#文章
[入门 elasticSearch](http://wiki.jikexueyuan.com/project/elasticsearch-definitive-guide-cn/)
#chrome强制https访问解决办法
    chrome://net-internals/#hsts
#go 并发编程 例子
[go 并发编程](http://studygolang.com/articles/2423)
## cookie
[js-cookie 文档](https://github.com/js-cookie/js-cookie)
##  Toastr
[Toastr](http://codeseven.github.io/toastr/)

#用到的包
##cron
    go get github.com/jakecoffman/cron
##cron job案例
[beego网站开发 定时执行任务](http://blog.csdn.net/u013401219/article/details/47278219)

#mysql 使用
##mysql 安装
- yum install mysql
- service mysqld start
- 修改密码
    1.用root 进入mysql后
    mysql>set password =password('wuchaofei1');
    mysql>flush privileges;
    
#主机 45.62.101.92
[主页](http://45.62.101.92)
    CREATE DATABASE IF NOT EXISTS beego DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
#悟空搜索模块
    go get 
##一个问题的解决
    exec: "gcc": executable file not found in %PATH%
    cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
##解决方案
[mingw-w64 下载地址](https://sourceforge.net/projects/mingw-w64/files/latest/download)
[mingw-w64 下载地址](http://www.mingw-w64.org/doku.php/download)

    自行安装mingw 64位，即可解决。注意要将bin目录添加到%PATH%环境变量。
    http://blog.csdn.net/mecho/article/details/24305369
    https://sourceforge.net/projects/mingw/files/latest/download?source=files
    https://sourceforge.net/projects/mingw-w64/files/latest/download
    https://sourceforge.net/projects/mingw-w64/
#go get
    悟空搜索
        go get -u -v github.com/huichen/wukong
    新浪微博Go语言SDK gobo
        //go get -u github.com/huichen/gobo
#新浪微博
    https://github.com/huichen/gobo
#蓝灯
[蓝灯 github 源码](https://github.com/getlantern/lantern)
[蓝灯 getlantern](https://www.getlantern.org/)
[蓝灯 github](https://github.com/getlantern/forum)
[蓝灯最新版下载地址](https://github.com/getlantern/forum/issues/833)
[蓝灯 网盘下载](https://ln.sync.com/dl/8d3e0f650#jm5ygm7p-qceg64ka-9pdwj8fh-vdzgsayz)


#服务器配置信息
- cd /opt
- tar xzvf web.tar.gz -C web
- cd web
- chmod+x web
- nohup ./web &
- cat nohup.out

#https
[go https 服务](http://studygolang.com/articles/2946)
[windows https 服务](http://www.cnblogs.com/developer-ios/p/6074665.html)

#加薪邮件范例
[加薪](http://shenqingshu.yjbys.com/baogao/89376.html)

##codis 资料
[Codis集群的搭建与使用 - GoogSQL - 博客园](http://www.cnblogs.com/xuanzhi201111/p/4425194.html)
[Codis 高可用负载均衡群集的搭建与使用 - 李惟忠的技术博客 - 51CTO技术博客](http://liweizhong.blog.51cto.com/1383716/1639918)
[github codis](https://github.com/CodisLabs/codis/blob/release3.2/doc/tutorial_zh.md)
D:\zhongliang\go\goweb\src\readMe.md
[ubuntu 安装go](http://blog.csdn.net/sunylat/article/details/50812998)
#ubuntu 安装go
1. 下载地址：
64位：https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
32位：https://storage.googleapis.com/golang/go1.6.linux-386.tar.gz
2. 解压缩到想放置GO语言的位置。
我放到了：“/usr/local/go”
3. 配置Ubuntu的环境变量
export GOROOT=/usr/local/go
export GOBIN=/usr/local/go/bin
export PATH=$PATH:$GOBIN
#java 安装
	http://www.cnblogs.com/a2211009/p/4265225.html
    sudo add-apt-repository ppa:webupd8team/java
    sudo apt-get update
    sudo apt-get install oracle-java8-installer
    sudo apt-get install oracle-java8-set-default
#release版本
https://github.com/CodisLabs/codis/releases
http://blog.csdn.net/dc_726/article/details/47052607
#go 下载
http://www.golangtc.com/download
## 无坑安装godep
http://studygolang.com/articles/7922

1. git clone https://github.com/golang/tools.git

##apt-get修改源
http://www.cnblogs.com/lyon2014/p/4715379.html

#redis 下载
https://github.com/antirez/redis-hashes/blob/master/README
##codis 文档
https://github.com/CodisLabs/codis/blob/release3.2/doc/tutorial_zh.md
http://www.cnblogs.com/xuanzhi201111/p/4425194.html
##etcd 安装
http://www.linuxdiyf.com/linux/18212.html
##codis 操作
http://www.cnblogs.com/softidea/p/5365640.html

