// author: ashing
// time: 2020/4/16 1:47 下午
// mail: axingfly@gmail.com
// Less is more.

package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("x", func(t *testing.T) {
		repeated := Repeat("x", 5)
		expected := "xxxxx"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

	t.Run("other", func(t *testing.T) {
		repeated := Repeat("s", 3)
		expected := "sss"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	str := Repeat("panda", 3)
	fmt.Println(str)
	// Output: pandapandapanda
}
