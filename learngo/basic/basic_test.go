// author: ashing
// time: 2020/4/19 3:26 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import "testing"

func TestTriangle(t *testing.T) {
	a, b := 3, 4
	got := triangle(a, b)
	want := 5
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
