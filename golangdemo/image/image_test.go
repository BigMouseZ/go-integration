package main

import "testing"

/*
image实现了基本的2D图片库。

基本接口叫作Image。图片的色彩定义在image/color包。

Image接口可以通过调用如NewRGBA和NewPaletted函数等获得；也可以通过调用Decode函数解码包含GIF、JPEG或PNG格式图像数据的输入流获得。解码任何具体图像类型之前都必须注册对应类型的解码函数。注册过程一般是作为包初始化的副作用，放在包的init函数里。因此，要解码PNG图像，只需在程序的main包里嵌入如下代码：

import _ "image/png"
_表示导入包但不使用包中的变量/函数/类型，只是为了包初始化函数的副作用。

*/
func Test(t *testing.T) {

}
