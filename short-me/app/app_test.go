package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ronething/short-me/app"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	expTime   = 60
	longURL   = "https://blog.ronething.cn"
	shortLink = "panda"
)

type storageMock struct {
	mock.Mock
}

var (
	app   main.App
	mockR *storageMock
)

func (s *storageMock) Shorten(url string, exp int64) (string, error) {

	var (
		args mock.Arguments
	)
	args = s.Called(url, exp)

	return args.String(0), args.Error(1)
}

func (s *storageMock) ShortLinkInfo(eid string) (interface{}, error) {

	var (
		args mock.Arguments
	)
	args = s.Called(eid)

	return args.Get(0), args.Error(1)
}

func (s *storageMock) UnShorten(eid string) (string, error) {

	var (
		args mock.Arguments
	)
	args = s.Called(eid)

	return args.String(0), args.Error(1)
}

func init() {
	app = main.App{}
	mockR = new(storageMock)
	app.Initialize(&main.Env{S: mockR})

}

type shortLinkResp struct {
	ShortLink string `json:"short_link"`
}

type shortLinkInfoResp struct {
	URL                 string `json:"url"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes"`
}

func TestCreateShortLink(t *testing.T) {
	var (
		jsonStr []byte
		req     *http.Request
		rw      *httptest.ResponseRecorder
		resp    shortLinkResp
		err     error
	)
	jsonStr = []byte(`{"url":"https://blog.ronething.cn", "expiration_in_minutes": 60}`)
	if req, err = http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonStr)); err != nil {
		t.Fatal("request error", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// once 表示只执行一次
	mockR.On("Shorten", longURL, int64(expTime)).Return(shortLink, nil).Once()

	rw = httptest.NewRecorder()
	app.Router.ServeHTTP(rw, req)

	if rw.Code != http.StatusCreated {
		t.Fatal("code unexcepted", rw.Code)
	}

	resp = shortLinkResp{}

	test := rw.Body.Bytes()

	test = test

	if err = json.NewDecoder(rw.Body).Decode(&resp); err != nil {
		t.Fatal("can not decode the resp")
	}

	if resp.ShortLink != shortLink {
		t.Fatal("unexcepted shortLink", resp.ShortLink)
	}

}

func TestGetShortLink(t *testing.T) {

	var (
		r             string
		req           *http.Request
		rw            *httptest.ResponseRecorder
		resp          shortLinkInfoResp
		shortLinkInfo shortLinkInfoResp
		err           error
	)

	shortLinkInfo = shortLinkInfoResp{URL: "https://blog.ronething.cn", ExpirationInMinutes: 60}

	r = fmt.Sprintf("/api/info?shortLink=%s", shortLink)

	if req, err = http.NewRequest("GET", r, nil); err != nil {
		t.Fatal("request error", err)
	}

	// once 表示只执行一次
	mockR.On("ShortLinkInfo", shortLink).Return(shortLinkInfo, nil).Once()

	rw = httptest.NewRecorder()
	app.Router.ServeHTTP(rw, req)

	if rw.Code != http.StatusOK {
		t.Fatal("code unexcepted", rw.Code)
	}

	resp = shortLinkInfoResp{}

	if err = json.NewDecoder(rw.Body).Decode(&resp); err != nil {
		t.Fatal("can not decode the resp", err)
	}

	if resp.URL != shortLinkInfo.URL {
		t.Fatal("unexcepted shortLink", resp.URL)
	}

}

func TestRedirect(t *testing.T) {

	var (
		r   string
		req *http.Request
		rw  *httptest.ResponseRecorder
		err error
	)
	r = fmt.Sprintf("/%s", shortLink)

	if req, err = http.NewRequest("GET", r, nil); err != nil {
		t.Fatal("request error", err)
	}

	// once 表示只执行一次
	mockR.On("UnShorten", shortLink).Return(longURL, nil).Once()

	rw = httptest.NewRecorder()
	app.Router.ServeHTTP(rw, req)

	if rw.Code != http.StatusTemporaryRedirect {
		t.Fatal("code unexcepted", rw.Code)
	}

}
