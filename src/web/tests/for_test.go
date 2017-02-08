package test

import (
	"testing"
	"github.com/astaxie/beego"
)

func Test_for(t *testing.T) {
	length:=1001
	max:=50

	itemcount:=length/max
	for i,start,end:=0,0,max;i<itemcount;i++{
		beego.Info(start,end)
		start=end
		end+=max
	}
	if(length%max!=0){
		beego.Info(length-length%max)
	}
}
