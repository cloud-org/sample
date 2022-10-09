// author: ashing
// time: 2020/4/18 1:21 上午
// mail: axingfly@gmail.com
// Less is more.

package racer

import (
	"fmt"
	"net/http"
	"time"
)

func httpGetTime(url string) time.Duration {
	s := time.Now()
	resp, _ := http.Get(url)
	d := time.Since(s)
	if resp.StatusCode == 200 {
		fmt.Println(url, " status code is 200")
	}

	return d
}

//func Racer(a, b string) (winner string) {
//	aDuration := httpGetTime(a)
//
//	bDuration := httpGetTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//
//	return b
//}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
