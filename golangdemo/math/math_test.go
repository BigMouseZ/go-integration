package main

import (
	"fmt"
	"math"
	"testing"
)

func TestOne(t *testing.T) {
	num := math.Abs(float64(2302.2))
	fmt.Printf("type:%T value:%#v\n", num, num)
	i := 1
	fmt.Println(math.Abs(float64(i))) // 取到绝对值
	fmt.Println(math.Ceil(3.8))       // 向上取整
	fmt.Println(math.Floor(3.6))      // 向下取整
	fmt.Println(math.Mod(11, 3))      // 取余数 11%3 效果一样
	fmt.Println(math.Modf(3.22))      // 取整数跟小数
	fmt.Println(math.Pow(3, 2))       // X 的 Y次方
	fmt.Println(math.Pow10(3))        // 10的N次方
	fmt.Println(math.Sqrt(8))         // 开平方
	fmt.Println(math.Cbrt(8))         // 开立方
	fmt.Println(math.Pi)              // π

}

func TestAbs(t *testing.T) {
	fmt.Println(math.Abs(2.999312323132141665374)) // 2.9993123231321417
	fmt.Println(math.Abs(2.999312323132141465374)) // 2.9993123231321412

}

func TestConst(t *testing.T) {

	/*	取值极限：
		    MaxInt8   = 1<<7 - 1
		    MinInt8   = -1 << 7
		    MaxInt16  = 1<<15 - 1
		    MinInt16  = -1 << 15
		    MaxInt32  = 1<<31 - 1
		    MinInt32  = -1 << 31
		    MaxInt64  = 1<<63 - 1
		    MinInt64  = -1 << 63
		    MaxUint8  = 1<<8 - 1
		    MaxUint16 = 1<<16 - 1
		    MaxUint32 = 1<<32 - 1
		    MaxUint64 = 1<<64 - 1*/
	fmt.Println(math.MaxInt8) // 127
	fmt.Println(math.MinInt8) // -128

	/*
	             E   = 2.71828182845904523536028747135266249775724709369995957496696763 // A001113
	   	    Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // A000796
	   	    Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // A001622
	   	    Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // A002193
	   	    SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // A019774
	   	    SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // A002161
	   	    SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // A139339
	   	    Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // A002162
	   	    Log2E  = 1 / Ln2
	   	    Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // A002392
	   	    Log10E = 1 / Ln10
	   	    MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	   	    SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)
	   	    MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	   	    SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324  // 1 / 2**(1023 - 1 + 52)*/

	fmt.Println(math.Pi) // -128
}

func TestMaxAndMin(t *testing.T) {
	fmt.Println(math.Max(1000, 200)) // 1000
	fmt.Println(math.Min(1000, 200)) // 200

	fmt.Println(math.Dim(1000, 2000)) // 0
	fmt.Println(math.Dim(1000, 200))  // 800
}

/*
三角函数
正弦函数，反正弦函数，双曲正弦，反双曲正弦

func Sin(x float64) float64
func Asin(x float64) float64
func Sinh(x float64) float64
func Asinh(x float64) float64
一次性返回sin,cos

func Sincos(x float64) (sin, cos float64)
余弦函数，反余弦函数，双曲余弦，反双曲余弦

func Cos(x float64) float64
func Acos(x float64) float64
func Cosh(x float64) float64
func Acosh(x float64) float64
正切函数，反正切函数，双曲正切，反双曲正切

func Tan(x float64) float64
func Atan(x float64) float64 和 func Atan2(y, x float64) float64
func Tanh(x float64) float64
func Atanh(x float64) float64
幂次函数
func Cbrt(x float64) float64 //立方根函数
func Pow(x, y float64) float64 // x的幂函数
func Pow10(e int) float64 // 10根的幂函数
func Sqrt(x float64) float64 // 平方根
func Log(x float64) float64 // 对数函数
func Log10(x float64) float64 // 10为底的对数函数
func Log2(x float64) float64 // 2为底的对数函数
func Log1p(x float64) float64 // log(1 + x)
func Logb(x float64) float64 // 相当于log2(x)的绝对值
func Ilogb(x float64) int // 相当于log2(x)的绝对值的整数部分
func Exp(x float64) float64 // 指数函数
func Exp2(x float64) float64 // 2为底的指数函数
func Expm1(x float64) float64 // Exp(x) - 1
特殊函数
func Inf(sign int) float64 // 正无穷
func IsInf(f float64, sign int) bool // 是否正无穷
func NaN() float64 // 无穷值
func IsNaN(f float64) (is bool) // 是否是无穷值
func Hypot(p, q float64) float64 // 计算直角三角形的斜边长
类型转化函数
func Float32bits(f float32) uint32 // float32和unit32的转换
func Float32frombits(b uint32) float32 // uint32和float32的转换
func Float64bits(f float64) uint64 // float64和uint64的转换
func Float64frombits(b uint64) float64 // uint64和float64的转换
其他函数
func Abs(x float64) float64 // 绝对值函数
func Ceil(x float64) float64 // 向上取整
func Floor(x float64) float64 // 向下取整
func Mod(x, y float64) float64 // 取模
func Modf(f float64) (int float64, frac float64) // 分解f，以得到f的整数和小数部分
func Frexp(f float64) (frac float64, exp int) // 分解f，得到f的位数和指数
func Max(x, y float64) float64 // 取大值
func Min(x, y float64) float64 // 取小值
func Dim(x, y float64) float64 // 复数的维数
func J0(x float64) float64 // 0阶贝塞尔函数
func J1(x float64) float64 // 1阶贝塞尔函数
func Jn(n int, x float64) float64 // n阶贝塞尔函数
func Y0(x float64) float64 // 第二类贝塞尔函数0阶
func Y1(x float64) float64 // 第二类贝塞尔函数1阶
func Yn(n int, x float64) float64 // 第二类贝塞尔函数n阶
func Erf(x float64) float64 // 误差函数
func Erfc(x float64) float64 // 余补误差函数
func Copysign(x, y float64) float64 // 以y的符号返回x值
func Signbit(x float64) bool // 获取x的符号
func Gamma(x float64) float64 // 伽玛函数
func Lgamma(x float64) (lgamma float64, sign int) // 伽玛函数的自然对数
func Ldexp(frac float64, exp int) float64 // value乘以2的exp次幂
func Nextafter(x, y float64) (r float64) //返回参数x在参数y方向上可以表示的最接近的数值，若x等于y，则返回x
func Nextafter32(x, y float32) (r float32) //返回参数x在参数y方向上可以表示的最接近的数值，若x等于y，则返回x
func Remainder(x, y float64) float64 // 取余运算
func Trunc(x float64) float64 // 截取函数*/
