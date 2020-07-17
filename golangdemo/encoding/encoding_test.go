package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestBase64(t *testing.T) {
	src := []byte("abcdefg")                     // 准备数据
	n := base64.StdEncoding.EncodedLen(len(src)) // 计算加密后数据的长度
	dst := make([]byte, n)                       // 创建容器
	base64.StdEncoding.Encode(dst, src)          // 加密数据
	fmt.Println(dst)                             // 打印结果：[89 87 74 106 90 71 86 109 90 119 61 61]
	fmt.Printf("%s\n", dst)                      // 转换为字符串： YWJjZGVmZw==

	fmt.Println(base64.StdEncoding.EncodeToString(src)) // 结果：YWJjZGVmZw=

	wd, _ := os.Getwd()                                   // 获取当前的工作目录
	file, _ := os.Create(wd + "/a.txt")                   // 创建一个文件，用来存储编码后的内容
	wClose := base64.NewEncoder(base64.StdEncoding, file) // 新建一个流式的编码器
	wClose.Write([]byte("abcdefg"))                       // 往编码器中写入数据
	wClose.Close()                                        // 关闭流，注意一定要手动调用Close方法，将缓存中还没写入到目标流的数据刷入到目标流里面
	// 最后，abcdefg经过base64编码后的内容写入到了a.txt文件里面
}

type Student struct {
	Name string
	Age  int
}

func TestJson(t *testing.T) {

	/*	func (m *RawMessage) MarshalJSON() ([]byte, error)
		MarshalJSON返回 *m 的json编码。
		func (m *RawMessage) UnmarshalJSON(data []byte) error
		UnmarshalJSON将*m设为data的一个拷贝*/
	var raw json.RawMessage
	j, err := raw.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(j))
	err = raw.UnmarshalJSON([]byte(`{"name":"lai"}`))
	if err != nil {
		log.Println(err)
	}
	j, err = raw.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(j))

	// HTMLEscape 函数将json编码的src中的<、>、&、U+2028和U+2029字符替换为\u003c、\u003e、\u0026、\u2028、\u2029转义字符串，
	// 以便json编码可以安全的嵌入HTML的<script>标签里。因为历史原因，网络浏览器不支持<script>标签中使用标准HTML转义，因此必须使用另一种json编码方案。
	src := []byte(`{"name":"<lai>"}`)
	buffer := bytes.NewBuffer(nil)
	json.HTMLEscape(buffer, src)
	fmt.Println(buffer.String())

	// Indent函数将json编码的调整缩进之后写入dst。每一个json元素/数组都另起一行开始，以prefix为起始，一或多个indent缩进（数目看）
	ss := []Student{
		{"lai", 12},
		{"ada", 15}}
	b, err := json.Marshal(ss)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)
	fmt.Println(string(b))
}

// 测试使用。
// go test JsonParser_test.go -v
func TestJson2(t *testing.T) {

	t.Log("here ..")
	jsonData := `{"NickName": "zhangsan", "Age": 26, "MaxMoney": 199.66}`

	t.Log(jsonData)

	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &jsonObj)

	if err != nil {
		panic(err)
		t.Error("Json字符串拼写错误")
		return
	}

	idx := 0
	len := len(jsonObj)
	t.Log(len)
	sql := "\nCREATE TABLE IF NOT EXISTS  `` ( \n"
	sql += "  `id` bigint(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT ,\n"
	for k, v := range jsonObj {
		t.Log("map key , value : ", k, v)
		idx += 1
		name := getColumnName(k)
		sql += "  `" + name + "` " + getColumnType(v) + "COMMENT \"\""
		if idx < len {
			sql += ","
		}
		sql += "\n"
	}
	sql += ") ENGINE=InnoDB  AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ; "
	t.Log(sql)

}

// 参考： https://blog.csdn.net/yang8023tao/article/details/53737458
// 返回 column 名称，解决驼峰问题
func getColumnName(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))

}

// 返回 column 类型，转换 int float 和 string类型。
func getColumnType(v interface{}) string {
	out := ""

	valueType := reflect.TypeOf(v)
	switch v := reflect.ValueOf(v); v.Kind() {
	case reflect.String:
		fmt.Println(v.String())
		out = "varchar(200)"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println(v.Int())
		out = "bigint(20)"
	case reflect.Float32, reflect.Float64:
		fmt.Println(v.Float())
		out = "bigint(20)"
	default:
		fmt.Printf("unhandled kind %s", v.Kind())
		out = "varchar(200)"
	}
	fmt.Println("map value type : ", valueType)

	return out
}

func TestXml(t *testing.T) {

}

func TestGob(t *testing.T) {

}

func TestBinary(t *testing.T) {

}
