package test

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"testing"
)

func Test_md5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("wuchaofei1")) // 需要加密的字符串为 sharejs.com
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
}
