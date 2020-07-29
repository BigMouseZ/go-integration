package main

// config 是一个日志配置项
type Config struct {
	filepath string `conf:"file_path" db:"name"`
	filePath string `conf:"file_name"`
	maxSize  int64  `conf:"max_size"`
}

func pareConf() {
	//1.打开文件

}

func main() {
	// 2.读取文件内容
	// 3.一行一行读取内容，根据tag找结构体里面对应的字段
	// 4.找到了要赋值

}
