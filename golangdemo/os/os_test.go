package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {

	defer func() { fmt.Println("d打印前") }()
	defer func() { fmt.Println("d打印中") }()
	defer func() { fmt.Println("d打印后") }()

	panic("执行异常")
}
func TestAllwaysUse(t *testing.T) {
	// os包中实现了平台无关的接口，设计向Unix风格，但是错误处理是go风格，当os包使用时，如果失败之后返回错误类型而不是错误数量．
	// 获取主机名
	fmt.Println(os.Hostname())

	// 获取当前目录
	fmt.Println(os.Getwd())

	// 获取用户ID
	fmt.Println(os.Getuid())

	// 获取有效用户ID
	fmt.Println(os.Geteuid())

	// 获取组ID
	fmt.Println(os.Getgid())

	// 获取有效组ID
	fmt.Println(os.Getegid())

	// 获取进程ID
	fmt.Println(os.Getpid())

	// 获取父进程ID
	fmt.Println(os.Getppid())

	// 获取环境变量的值
	fmt.Println(os.Getenv("GOPATH"))

	// 设置环境变量的值
	// os.Setenv("TEST", "test")

	// 改变当前工作目录
	// os.Chdir("C:/")
	fmt.Println(os.Getwd())

	// 创建文件
	f1, _ := os.Create("./1.txt")
	defer f1.Close()

	// 修改文件权限
	if err := os.Chmod("./1.txt", 0777); err != nil {
		fmt.Println(err)
	}

	// 修改文件所有者
	if err := os.Chown("./1.txt", 0, 0); err != nil {
		fmt.Println(err)
	}

	// 修改文件的访问时间和修改时间
	os.Chtimes("./1.txt", time.Now().Add(time.Hour), time.Now().Add(time.Hour))

	// 获取所有环境变量
	fmt.Println(strings.Join(os.Environ(), "\r\n"))

	// 把字符串中带${var}或$var替换成指定指符串
	fmt.Println(os.Expand("${1} ${2} ${3}", func(k string) string {
		mapp := map[string]string{
			"1": "111",
			"2": "222",
			"3": "333",
		}
		return mapp[k]
	}))

	// 创建目录
	os.Mkdir("abc", os.ModePerm)

	// 创建多级目录
	os.MkdirAll("abc/d/e/f", os.ModePerm)

	// 删除文件或目录
	os.Remove("abc/d/e/f")

	// 删除指定目录下所有文件
	os.RemoveAll("abc")

	// 重命名文件
	os.Rename("./2.txt", "./2_new.txt")

	// 判断是否为同一文件
	// unix下通过底层结构的设备和索引节点是否相同来判断
	// 其他系统可能是通过文件绝对路径来判断
	fs1, _ := f1.Stat()
	f2, _ := os.Open("./1.txt")
	fs2, _ := f2.Stat()
	fmt.Println(os.SameFile(fs1, fs2))

	// 返回临时目录
	fmt.Println(os.TempDir())
}

func TestFile(t *testing.T) {
	// 创建文件，返回一个文件指针
	f3, _ := os.Create("./3.txt")
	defer f3.Close()

	// 以读写方式打开文件，如果不存在则创建文件，等同于上面os.Create
	f4, _ := os.OpenFile("./4.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f4.Close()
	/*	打开标记：

		O_RDONLY：只读模式(read-only)
		O_WRONLY：只写模式(write-only)
		O_RDWR：读写模式(read-write)
		O_APPEND：追加模式(append)
		O_CREATE：文件不存在就创建(create a new file if none exists.)
		O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
		O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
		O_TRUNC：打开并清空文件
		至于操作权限perm，除非创建文件时才需要指定，不需要创建新文件时可以将其设定为０.虽然go语言给perm权限设定了很多的常量，但是习惯上也可以直接使用数字，如0666(具体含义和Unix系统的一致).
	*/ // 打开文件，返回文件指针
	f1, _ := os.Open("./1.txt")
	defer f1.Close()

	// 修改文件权限，类似os.chmod
	f1.Chmod(0777)

	// 修改文件所有者，类似os.chown
	f1.Chown(0, 0)

	// 返回文件的句柄，通过NewFile创建文件需要文件句柄
	fmt.Println(f1.Fd())

	// 从文件中读取数据
	buf := make([]byte, 128)
	// read每次读取数据到buf中
	for n, _ := f1.Read(buf); n != 0; n, _ = f1.Read(buf) {
		fmt.Println(string(buf[:n]))
	}

	// 向文件中写入数据
	for i := 0; i < 5; i++ {
		f3.Write([]byte("写入数据" + strconv.Itoa(i) + "\r\n"))
	}

	// 返回一对关联的文件对象
	// 从r中可以读取到从w写入的数据
	r, w, _ := os.Pipe()
	// 向w中写入字符串
	w.WriteString("写入w")
	buf2 := make([]byte, 128)
	// 从r中读取数据
	n, _ := r.Read(buf)
	fmt.Println(string(buf2[:n]))

	// 改变工作目录
	os.Mkdir("a", os.ModePerm)
	dir, _ := os.Open("a")
	// 改变工作目录到dir，dir必须为一个目录
	dir.Chdir()
	fmt.Println(os.Getwd())

	// 读取目录的内容，返回一个FileInfo的slice
	// 参数大于0，最多返回n个FileInfo
	// 参数小于等于0，返回所有FileInfo
	fi, _ := dir.Readdir(-1)
	for _, v := range fi {
		fmt.Println(v.Name())
	}

	// 读取目录中文件对象的名字
	names, _ := dir.Readdirnames(-1)
	fmt.Println(names)

	// 获取文件的详细信息，返回FileInfo结构
	fi3, _ := f3.Stat()
	// 文件名
	fmt.Println(fi3.Name())
	// 文件大小
	fmt.Println(fi3.Size())
	// 文件权限
	fmt.Println(fi3.Mode())
	// 文件修改时间
	fmt.Println(fi3.ModTime())
	// 是否是目录
	fmt.Println(fi3.IsDir())
}

func TestFileTwo(t *testing.T) {
	// func (f *File) Seek(offset int64, whence int) (ret int64, err error)
	// Seek设置下一次读或写操作的偏移量offset，根据whence来解析：
	// 0意味着相对于文件的原始位置，1意味着相对于当前偏移量，2意味着相对于文件结尾。它返回新的偏移量和错误（如果存在）。
	s := make([]byte, 10)
	file, _ := os.Open("tmp.txt")
	defer file.Close()
	file.Seek(-12, 2) // 从离最后位置12的地方开始
	n, _ := file.Read(s)
	fmt.Println(string(s[:n]))

}
func TestFileThree(t *testing.T) {
	// func (f *File) Stat() (fi FileInfo, err error)　　//返回文件描述相关信息，包括大小，名字等．等价于os.Stat(filename string)
	// func (f *File) Sync() (err error)                        //同步操作，将当前存在内存中的文件内容写入硬盘．
	// func (f *File) Truncate(size int64) error        //类似  os.Truncate(name, size),，将文件进行截断
	// func (f *File) Write(b []byte) (n int, err error)　　//将b中的数据写入f文件
	// func (f *File) WriteAt(b []byte, off int64) (n int, err error)　//将b中数据写入f文件中，写入时从offset位置开始进行写入操作
	// func (f *File) WriteString(s string) (ret int, err error)　　　//将字符串s写入文件中
	file, _ := os.Create("tmp.txt")
	a := "hellobyte"
	file.WriteAt([]byte(a), 10) // 在文件file偏移量10处开始写入hellobyte
	file.WriteString("string")  // 在文件file偏移量0处开始写入string
	file.Write([]byte(a))       // 在文件file偏移量string之后开始写入hellobyte，这个时候就会把开始利用WriteAt在offset为10处开始写入的hellobyte进行部分覆盖
	b := make([]byte, 20)
	file.Seek(0, 0) // file指针指向文件开始位置
	n, _ := file.Read(b)
	fmt.Println(string(b[:n])) // stringhellobytebyte，这是由于在写入过程中存在覆盖造成的
}
func TestFileFileModel(t *testing.T) {
	// FileModel 的方法主要用来进行判断和输出权限
	// func (m FileMode) IsDir() bool              //判断m是否是目录，也就是检查文件是否有设置的ModeDir位
	// func (m FileMode) IsRegular() bool　　//判断m是否是普通文件，也就是说检查m中是否有设置mode type
	// func (m FileMode) Perm() FileMode　　//返回m的权限位
	// func (m FileMode) String() string　　　　//返回m的字符串表示
	fd, err := os.Stat("tmp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fm := fd.Mode()
	fmt.Println(fm.IsDir())     // false
	fmt.Println(fm.IsRegular()) // true
	fmt.Println(fm.Perm())      // -rwxrwxrwx
	fmt.Println(fm.String())    // -rwxrwxrwx
}
