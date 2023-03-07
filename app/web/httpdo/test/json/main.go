package main

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/andytyc/goutil/app/web/httpdo"
)

type Book struct {
	ID      int
	Name    string
	Content []byte
}

// 了解自定义json编译器，可以参考:
// https://blog.isayme.org/posts/issues-44/

// 了解base64编码
// https://www.ruanyifeng.com/blog/2008/06/base64.html

func main() {
	// ********** 常用 **********

	dataCommon()

	// ********** 特殊情况1 **********

	// 有时我们传入给 &httpdo.ResData{} 的 Data 是 []byte 时，比较特殊. 当用内置的json编译器时[]byte接口值会被进行base64编码为字符串
	dataWithByte()
}

func dataCommon() {
	// data 是 Book 结构体
	book := &Book{ID: 1, Name: "book-1", Content: []byte("hello world book-1")}

	res := httpdo.NewResData(httpdo.Code_10200, "", book, nil)

	// http服务一般会进行json字符串格式交互
	resData, err := json.Marshal(res)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("resData :", string(resData))

	// 可以看到http响应的json字符串, {如同postman调试}
	// 2023/03/07 15:37:09 resData : {"code":10200,"msg":"请求成功","data":{"ID":1,"Name":"book-1","Content":"aGVsbG8gd29ybGQgYm9vay0x"},"more_info":""}

	reqres := &httpdo.ResData{Data: &Book{}} // 正确示范, 注意: 若提前设置 Data 接口传入的值类型 &Book{}, 则下边断言Data类型就是: *Book 所以，这里需要指明类型，这样json反序列化直接解析
	// reqres := &httpdo.ResData{} // 错误示范, 注意: 若没有设置Data接口值类型，则下边断言Data类型就是: map[string]interface{}
	err = json.Unmarshal(resData, reqres)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("reqres :", reqres)

	switch reqres.Data.(type) {
	case string:
		log.Println("reqres is string")
	case map[string]interface{}:
		log.Println("reqres is map[string]interface{}")
	case *Book:
		log.Println("reqres is *Book")
	default:
		log.Println("unknown")
	}

	// reqresdata := reqres.Data.(map[string]interface{})
	// log.Println("reqresdata :", reqresdata)

	// 2023/03/07 15:42:53 reqres is *Book

	reqbook := reqres.Data.(*Book)
	log.Println("reqbook :", reqbook, "Content", string(reqbook.Content))

	// 2023/03/07 15:42:53 reqbook : &{1 book-1 [104 101 108 108 111 32 119 111 114 108 100 32 98 111 111 107 45 49]} Content hello world book-1
}

func dataWithByte() {
	// data 是[]byte

	// book 是响应的具体数据
	book := &Book{ID: 1, Name: "book-1", Content: []byte("hello world book-1")}

	// 内置json对象序列化
	data, err := json.Marshal(book)
	if err != nil {
		log.Panicln(err)
	}

	// data 传入[]byte 若不是 &Book
	res := httpdo.NewResData(httpdo.Code_10200, "", data, nil)

	// http服务一般会进行json字符串格式交互
	resData, err := json.Marshal(res)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("resData :", string(resData))

	// 可以看到http响应的json字符串, {如同postman调试}
	// 2023/03/07 14:57:45 resData : {"code":10200,"msg":"请求成功","data":"eyJJRCI6MSwiTmFtZSI6ImJvb2stMSIsIkNvbnRlbnQiOiJhR1ZzYkc4Z2QyOXliR1FnWW05dmF5MHgifQ==","more_info":""}

	// 会发现data存储的数据类似不是[]byte, 而是string, 并且注意末尾有俩==
	// 这里的data字符串被编码了, string是base64编码后的字符串
	// 这个原因是: 在进行响应时进行的json编译器（一般是内置的encoding/json）序列化时，对于interface接口值是[]byte时，转换为了string, 并且是base64编码字符串

	// 先将json字符串解析道结构体中
	reqres := &httpdo.ResData{}
	err = json.Unmarshal(resData, reqres)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("reqres :", reqres)

	// 2023/03/07 15:06:16 reqres : &{10200 请求成功 eyJJRCI6MSwiTmFtZSI6ImJvb2stMSIsIkNvbnRlbnQiOiJhR1ZzYkc4Z2QyOXliR1FnWW05dmF5MHgifQ== }

	switch reqres.Data.(type) {
	case string:
		log.Println("reqres is string")
	case []byte:
		log.Println("reqres is bytes")
	default:
		log.Println("unknown")
	}

	// 2023/03/07 15:33:50 reqres is string

	// 此时，因为之前响应时data是序列化后的[]byte, 而不是Book, 所以需要反序列化下 -- 这个也是此用例测试的目的

	reqbook := &Book{}

	// 假如, 直接转换类型使用反序列化，会失败报错，如下两种情况:

	// err = json.Unmarshal(reqres.Data.([]byte), reqbook) // panic: interface conversion: interface {} is string, not []uint8

	// 解读: 直接转换为 []byte 失败，因为 Data 存储的却是 string

	// err = json.Unmarshal([]byte(reqres.Data.(string)), reqbook) // panic: invalid character 'e' looking for beginning of value

	// 解读: 先转换为 string 在转换为 []byte 然后进行json反序列化失败，因为 字符串不是json字符串，而是base64字符串

	// 正确的解决办法
	//
	// 先将string进行base64解码到[]byte, 而不是直接进行[]byte()转换

	dataByte, err := base64.StdEncoding.DecodeString(reqres.Data.(string))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dataByte, reqbook)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("reqbook :", reqbook, "Content", string(reqbook.Content))

	// 如此，就可以正确解析了
	// 2023/03/07 15:15:01 reqbook : &{1 book-1 [104 101 108 108 111 32 119 111 114 108 100 32 98 111 111 107 45 49]} Content hello world book-1
}
