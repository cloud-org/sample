// author: ashing
// time: 2020/4/16 5:12 下午
// mail: axingfly@gmail.com
// Less is more.

package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %.2f want %.2f", got, want)
	//	}
	//}
	//
	//t.Run("rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{12, 6}
	//	want := 72.0
	//	checkArea(t, rectangle, want)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	want := 314.16
	//	checkArea(t, circle, want)
	//})

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.16},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {

		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})

		//got := tt.shape.Area()
		//if got != tt.want {
		//	t.Errorf("got %.2f want %.2f", got, tt.want)
		//}
	}
}
