package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
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

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}
type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func TestXml(t *testing.T) {
	// 参考："github.com/beevik/etree"

	file, err := os.Open("servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

func TestGob(t *testing.T) {

}

func TestBinary(t *testing.T) {
	// 此包实现了对数据与byte之间的转换，以及varint的编解码。
	// 数据的byte序列化转换
	// func Read(r io.Reader, order ByteOrder, data interface{}) error
	/*
		参数列表：
		1）r  可以读出字节流的数据源
		2）order  特殊字节序，包中提供大端字节序和小端字节序
		3）data  需要解码成的数据
		返回值：error  返回错误
		功能说明：Read从r中读出字节数据并反序列化成结构数据。data必须是固定长的数据值或固定长数据的slice。
		从r中读出的数据可以使用特殊的 字节序来解码，并顺序写入value的字段。当填充结构体时，使用(_)名的字段讲被跳过。
	*/
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi)

	//Write(w io.Writer, order ByteOrder, data interface{}) error
	/*
		参数列表：
		1）w  可写入字节流的数据
		2）order  特殊字节序，包中提供大端字节序和小端字节序
		3）data  需要解码的数据
		返回值：error  返回错误
		功能说明：
		Write讲data序列化成字节流写入w中。data必须是固定长度的数据值或固定长数据的slice，或指向此类数据的指针。
		写入w的字节流可用特殊的字节序来编码。另外，结构体中的(_)名的字段讲忽略。
	*/
	buf = new(bytes.Buffer)
	pi = math.Pi

	err = binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.Bytes())

	//func Size(v interface{}) int
	/*
		参数列表：v  需要计算长度的数据
		返回值：int 数据序列化之后的字节长度
		功能说明：
		Size讲返回数据系列化之后的字节长度，数据必须是固定长数据类型、slice和结构体及其指针。
	*/
	var a int
	p := &a
	c := [10]int64{1}
	s := "adsa"
	bs := make([]byte, 10)

	fmt.Println(binary.Size(a))  // -1
	fmt.Println(binary.Size(p))  // -1
	fmt.Println(binary.Size(c))  // 80
	fmt.Println(binary.Size(s))  // -1
	fmt.Println(binary.Size(bs)) // 10

	//func PutUvarint(buf []byte, x uint64) int

	/*
		参数列表：
		1）buf  需写入的缓冲区
		2）x  uint64类型数字
		返回值：
		1）int  写入字节数。
		2）panic  buf过小。
		功能说明：
		PutUvarint主要是讲uint64类型放入buf中，并返回写入的字节数。如果buf过小，PutUvarint将抛出panic。
	*/

	u16 := 1234
	u64 := 0x1020304040302010
	sbuf := make([]byte, 4)
	buf2 := make([]byte, 10)

	ret := binary.PutUvarint(sbuf, uint64(u16))
	fmt.Println(ret, len(strconv.Itoa(u16)), sbuf)

	ret = binary.PutUvarint(buf2, uint64(u64))
	fmt.Println(ret, len(strconv.Itoa(u64)), buf)

	//func PutVarint(buf []byte, x int64) int

	/*
		参数列表：
		1）buf  需要写入的缓冲区
		2）x int64类型数字
		返回值：
		1）int  写入字节数
		2）panic  buf过小
		功能说明：
		PutVarint主要是讲int64类型放入buf中，并返回写入的字节数。如果buf过小，PutVarint将抛出panic。
	*/
	i16 := 1234
	i64 := -1234567890
	sbuf = make([]byte, 4)
	buf2 = make([]byte, 10)

	ret = binary.PutVarint(buf2, int64(i16))
	fmt.Println(ret, len(strconv.Itoa(i16)), sbuf)

	ret = binary.PutVarint(buf2, int64(i64))
	fmt.Println(ret, len(strconv.Itoa(i64)), buf)

	//func Uvarint(buf []byte) (uint64, int)
	/*
		参数列表：buf  需要解码的缓冲区
		返回值：
		1）uint64  解码的数据。
		2）int  解析的字节数。
		功能说明:
		Uvarint是从buf中解码并返回一个uint64的数据,及解码的字节数(>0)。如果出错,则返回数据0和一个小于等于0的字节数n,其意义为:
		1)n == 0: buf太小
		2)n < 0: 数据太大,超出uint64最大范围,且-n为已解析字节数
	*/
	sbuf = []byte{}
	buf4 := []byte{144, 192, 192, 132, 136, 140, 144, 16, 0, 1, 1}
	bbuf := []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, ret := binary.Uvarint(sbuf)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(buf4)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(bbuf)
	fmt.Println(num, ret)

	//func Varint(buf []byte) (int64, int)
	/*
		参数列表: buf  需要解码的缓冲区
		返回值:
		1) int64 解码的数据
		2) int  解析的字节数
		功能说明:
		Varint是从buf中解码并返回一个int64的数据,及解码的字节数(>0).如果出错,则返回数据0和一个小于等于0的字节数n,其意义为:
		1) n == 0: buf太小
		2) n < 0: 数据太大,超出64位,且-n为已解析字节数
	*/
	var sbuf2 []byte
	var buf7 []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
	var bbuf7 []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num2, ret := binary.Varint(sbuf2)
	fmt.Println(num, ret) //0 0

	num2, ret = binary.Varint(buf7)
	fmt.Println(num, ret) //580990878187261960 9

	num2, ret = binary.Varint(bbuf7)
	fmt.Println(num2, ret) //0 -11

	//func ReadUvarint(r io.ByteReader) (uint64, error)
	/*
		参数列表:
		返回值:
		1) uint64  解析出的数据
		2) error  返回的错误
		功能说明:
		ReadUvarint从r中解析并返回一个uint64类型的数据及出现的错误.
		功能说明:
		ReadUvarint从r中解析并返回一个uint64类型的数据及出现的错误.
	*/

	var sbuf []byte
	var buf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
	var bbuf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, err := binary.ReadUvarint(bytes.NewBuffer(sbuf))
	fmt.Println(num, err) //0 EOF

	num, err = binary.ReadUvarint(bytes.NewBuffer(buf))
	fmt.Println(num, err) //1161981756374523920 <nil>

	num, err = binary.ReadUvarint(bytes.NewBuffer(bbuf))
	fmt.Println(num, err) //4620746270195064848 binary: varint overflows a 64-bit integer
}
