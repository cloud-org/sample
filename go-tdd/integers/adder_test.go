// author: ashing
// time: 2020/4/16 1:22 下午
// mail: axingfly@gmail.com
// Less is more.

package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add 2 and 2 ", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("add any", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
