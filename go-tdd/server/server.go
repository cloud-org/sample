// author: ashing
// time: 2020/4/21 12:52 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

//func PlayerServer(w http.ResponseWriter, r *http.Request) {
//	//w.Write([]byte("20"))
//	//fmt.Fprint(w, "20")
//
//	// 1234 len(12)  [2:] = 34
//	player := r.URL.Path[len("/players/"):]
//	fmt.Fprint(w, GetPlayerScore(player))
//}

//func GetPlayerScore(name string) string {
//	if name == "Pepper" {
//		return "20"
//	}
//
//	if name == "Floyd" {
//		return "10"
//	}
//
//	return ""
//}
