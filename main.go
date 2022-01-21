package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

/**
第一步，爬取下载视频
第二部，打开解析视频
**/
func main() {
	getVideo()
}
func getVideo() {
	outputDir, _ := os.Getwd()
	// fmt.Println(outputDir)
	c := colly.NewCollector()
	c.OnRequest(func(req *colly.Request) {
		fmt.Println(req.URL)
	})
	c.OnResponse(func(resp *colly.Response) {
		resp.Save(outputDir + "/" + resp.FileName())
		// fmt.Println(resp.Body)
	})
	c.Visit("file://" + outputDir + "/video/lol.mp4")
	c.Wait()
}
