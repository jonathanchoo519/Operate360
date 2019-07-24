package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func SpiderPage() (string,error){
	url := "http://sd.360.cn/download_center.html"
	result ,err := HttpGet(url)
	if err != nil {
		panic(err)
		return "",err
	}

	//f,err1 := os.Create("update.html")
	//if err1 != nil{
	//	panic(err)
	//	return
	//}
	//
	//f.WriteString(result)
	//f.Close()

	dom,err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		panic(err)
		return "",err
	}

	var newDate string

	dom.Find("div[class*=virus]").Find("div.cls").Find("div[class*=fl]").Find("span[style*=margin]").Each(func(i int, selection *goquery.Selection) {
		newDate = selection.Text()

	})
	newDate = string( []rune(newDate)[12:22])
	return newDate,nil
}

func HttpGet(url string) (result string,err error){
	resp ,err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte,4096)
	for{
		n,err2 := resp.Body.Read(buf)
		if n==0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf)
	}
	return
}


//func SpiderPage(){
//	c := colly.NewCollector(
//		colly.AllowedDomains("sd.360.cn"),
//		colly.CacheDir("./download_center.html"),
//		)
//	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
//		fmt.Println(element.Text)
//	})
//
//}