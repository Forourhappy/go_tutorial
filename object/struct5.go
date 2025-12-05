package object

import (
	"fmt"
	"reflect"
)

type Car1 struct {
	name    string `tag:"차량명"`
	color   string `tag:"색상"`
	company string `tag:"제조사"`
}

func Struct5() {
	// 필드 태그 사용
	tag := reflect.TypeOf(Car1{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag.Get("tag"), tag.Field(i).Name, tag.Field(i).Type)
	}
}
