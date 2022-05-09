package interview

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func F019() {
	var _httpCli = &http.Client{
		Timeout: time.Duration(15) * time.Second,
		Transport: &http.Transport{
			//MaxIdleConns:          1,
			MaxIdleConnsPerHost:   10,
			MaxConnsPerHost:       10,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	i := 0
	//go func() {
	//	for {
	//		i++
	//		resp, err := _httpCli.Get("http://127.0.0.1")
	//		if err != nil {
	//			fmt.Println(10000+i, err)
	//		} else {
	//			bs, _ := ioutil.ReadAll(resp.Body)
	//			fmt.Println(10000+i, string(bs))
	//		}
	//		time.Sleep(time.Second)
	//	}
	//}()
	for {
		i++
		resp, err := _httpCli.Get("http://127.0.0.1")
		if err != nil {
			fmt.Println(10000+i, err)
		} else {
			bs, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(i, string(bs))
		}
		time.Sleep(time.Second)
	}
}
