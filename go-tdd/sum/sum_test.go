// author: ashing
// time: 2020/4/16 3:53 下午
// mail: axingfly@gmail.com
// Less is more.

package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	//
	//t.Run("collection of any size", func(t *testing.T) {
	//	numbers := []int{1, 2, 3}
	//
	//	got := Sum(numbers)
	//	want := 6
	//
	//	if got != want {
	//		t.Errorf("got %d want %d given, %v", got, want, numbers)
	//	}
	//})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf("got %v want %v", got, want)
	//}
}

// compare two []int slice if equal or not
func Equal(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) { // 限定类型为 []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
