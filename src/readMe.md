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
