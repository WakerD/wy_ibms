package test

import (
	"testing"
	"strings"
	"fmt"
	"strconv"
)

var a ="G"
//测试变量的全局和局部
func Test_Local_scope(t *testing.T) {
	n()
	m()
	n()
}

func n()  {
	print(a)
}

func m()  {
	a:="O"
	print(a)
}

//测试分割字符串
func Test_Split(t *testing.T) {
	str:="It's s/n/ow/ing in hunan today"
	result:=strings.Split(str,"/")
	fmt.Println("分割字符串：",result)
}

//测试string转int
func Test_Strconv(t *testing.T) {

	fmt.Printf("The size of ints is:%b/n",strconv.IntSize)
}