// author: ashing
// time: 2020/4/16 11:20 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import "testing"

func TestHello(t *testing.T) {
	//got := Hello("Chris")
	//want := "Hello, Chris"
	//
	//if got != want {
	//	t.Errorf("got '%q' want '%q'", got, want)
	//}

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // 标记辅助函数
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		//if got != want {
		//	t.Errorf("got '%q' want '%q'", got, want)
		//}
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		//if got != want {
		//	t.Errorf("got '%q' want '%q'", got, want)
		//}
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Panda", "French")
		want := "Bonjour, Panda"
		assertCorrectMessage(t, got, want)
	})
}
