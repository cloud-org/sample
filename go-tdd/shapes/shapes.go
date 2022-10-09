// author: ashing
// time: 2020/4/16 5:13 下午
// mail: axingfly@gmail.com
// Less is more.

package shapes

import (
	"fmt"
	"math"
	"strconv"
)

type Shape interface {
	Area() float64
}

// 长方形
type Rectangle struct {
	Width  float64 // field
	Height float64
}

// 圆形
type Circle struct {
	Radius float64
}

// 三角形
type Triangle struct {
	Base   float64 // 底
	Height float64 // 高
}

func Perimeter(r Rectangle) float64 {
	//return 0
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	//return 0
	return r.Width * r.Height
}

func str2float(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func (c Circle) Area() float64 {
	f := math.Pi * c.Radius * c.Radius
	s := fmt.Sprintf("%.2f", f) // 保留两位小数
	f = str2float(s)
	return f
}

func (t Triangle) Area() float64 {
	//return t.Base * t.Height * 0.5
	return t.Base * t.Height / 2
}
