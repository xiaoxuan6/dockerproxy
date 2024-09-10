package services

import (
	"encoding/json"
	"github.com/xiaoxuan6/dockerproxy/global"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup

	IndexService = new(indexService)
)

type Item struct {
	Url  string `json:"url"`
	Stat bool   `json:"stat"`
}

type indexService struct {
}

func (i indexService) FetchUrls() {
	resp, err := http.Get(os.Getenv("GIST_URL"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("获取数据失败：", err.Error())
		return
	}

	var urls []string
	_ = json.Unmarshal(b, &urls)

	global.Urls = urls
}

func (i indexService) FetchUrlsWithVerify() []Item {
	urls := make([]Item, 0)
	for _, url := range global.Urls {
		url := url
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &http.Client{
				Timeout: 5 * time.Second,
			}
			resp, err := client.Get(url)
			if err != nil {
				urls = append(urls, Item{
					url,
					false,
				})
				return
			}
			defer resp.Body.Close()

			target := true
			if resp.StatusCode != http.StatusOK {
				target = false
			}

			urls = append(urls, Item{
				url,
				target,
			})
		}()
	}

	wg.Wait()
	return urls
}
