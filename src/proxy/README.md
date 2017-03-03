# ytran
**这是一个反连socket5代理**

## 链接
## http://www.cnhonker.com

## Usage:
	-listen	 [port1] [port2]
	-connect [ip]	 [port]
	-help
## 编译:
    go build ytran.go
## 使用说明:
    $./ytran -listen 1080 9001
    $./ytran -connect 192.168.1.1 9001
    
    需要使用openssl生成ca.crt  ca.key  证书
    
    192.168.1.1:1080 就是socket5代理  漫游内网吧
#rfc1928
[rfc1928](http://blog.csdn.net/mycoolx/article/details/7496564)
#go socket编程
[go socket编程](http://www.cnblogs.com/leoncfor/p/5009263.html)
# s.go是一个socket5 tcp代理服务器