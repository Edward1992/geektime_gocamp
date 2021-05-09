package week3

import (
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestErrGroup(t *testing.T) {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.baidu.com",
		"https://www.zhihu.com",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			defer func() {
				if resp != nil {
					resp.Body.Close()
				}
			}()
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				t.Log(string(body))
			}
			return err
		})
	}

	if err := g.Wait(); err == nil {
		t.Log("Successfully fetched all URLs.")
	}
}