package test

import (
	"testing"
	"web/service"
	"log"
)

func Test_Search(t *testing.T) {
	book:=service.Search(" 流氓艳遇记 ")
	log.Println(book.Name)
	log.Println(book.Maker)
	log.Println(book.Url)
	log.Println(book.Content)
}
