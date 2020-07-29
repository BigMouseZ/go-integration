package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// config 是一个日志配置项，字段名大写
type Config struct {
	Filepath string `conf:"file_path" db:"name"`
	FilePath string `conf:"file_name"`
	MaxSize  int64  `conf:"max_size"`
}

func pareConf(fileName string, result interface{}) (err error) {
	// 0.前提条件，result必须是一个ptr
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	if t.Kind() != reflect.Ptr {
		err = errors.New("result必须是一个指针")
		return
	}
	// result不是结构体也不行
	tElem := t.Elem()
	if tElem.Kind() != reflect.Struct {
		err = errors.New("result必须是结构体指针")
		return
	}
	// 1.打开文件
	data, err := ioutil.ReadFile(fileName) // data:[]byte
	if err != nil {
		err = fmt.Errorf("打开配置文件%s失败", fileName)
		return err
	}
	// 2.读取文件内容,按照行分割，得到一个行的切片
	lineSlice := strings.Split(string(data), "/r/n")
	fmt.Println(lineSlice)
	// 一行行解析
	for index, line := range lineSlice {
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// 忽略空行和注释
			continue
		}
		// 解析是不是正经的配置项（判断是不是有等号）
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 {
			err = fmt.Errorf("第%d行语法错误", index+1)
		}
		// 按照=分割每一行，左边是key ，右边是value
		key := line[:equalIndex]
		value := line[equalIndex+1:]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if len(value) == 0 {
			err = fmt.Errorf("第%d行语法错误", index+1)
		}

		// 利用反射给resul赋值
		// 遍历结构体每个字段和key比较，匹配上了就把value赋值
		for i := 0; i < tElem.NumField(); i++ {
			field := tElem.Field(i)
			tag := field.Tag.Get("conf")
			if key == tag {
				// 匹配上了就把value赋值
				// 拿到每个字段的类型
				fieldType := field.Type
				fmt.Println(fieldType)
				switch fieldType.Kind() {
				case reflect.String:
					fieldValue := v.Elem().FieldByName(field.Name) // 根据字段名找到对应的字段值
					fieldValue.SetString(value)                    // 将配置文件中读取的value设置给结构体
				case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
					value64, _ := strconv.ParseInt(value, 10, 64)
					v.Elem().Field(i).SetInt(value64)
				}
			}
		}

	}

	return

}

func main() {
	// 2.读取文件内容
	// 3.一行一行读取内容，根据tag找结构体里面对应的字段
	// 4.找到了要赋值 D:\go_projects\go-integration\parseconfig\xxx.conf

	var c = &Config{} // 用来存储读取的数据
	err := pareConf("D:\\go_projects\\go-integration\\parseconfig\\xxx.conf", c)
	if err != nil {
		fmt.Println(err)
	}

}
