package transport

import (
	"io/ioutil"
	"net/http"
	"time"
)

func newCustomHttpClient() *http.Client {
	tran := http.Transport{
		DisableKeepAlives:     true,
		MaxIdleConns:          5,
		MaxIdleConnsPerHost:   0,
		MaxConnsPerHost:       0,
		IdleConnTimeout:       2 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Transport: &tran,
		Timeout:   5 * time.Second,
	}
}

func GetRSSFeed(url string) ([]byte, error) {
	resp, err := newCustomHttpClient().Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}