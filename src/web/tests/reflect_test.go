package test

import (
	"testing"
	"log"
	"reflect"
)

//测试 reflect 包怎么用

func TestReflect(t *testing.T) {
	var x int =1

	log.Println("Type:",reflect.TypeOf(x))
}
