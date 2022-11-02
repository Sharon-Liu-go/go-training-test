package main

import (
	"fmt"

	"github.com/Sharon-Liu-go/go-training-test/hello"
	colly "github.com/gocolly/colly/v2"
)

func main() {
	hello.HelloWorld()
	collyUseTemplate()
}

func collyUseTemplate() {
	count := 0
	// 創建採集器對象
	collector := colly.NewCollector()
	// 發起請求之前調用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("發起請求之前調用...")
	})
	// 請求期間發生錯誤,則調用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("請求期間發生錯誤,則調用:", err)
	})
	// 收到響應後調用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("收到響應後調用:", response.Request.URL)
	})
	//OnResponse如果收到的內容是HTML ,則在之後調用
	collector.OnHTML("img", func(element *colly.HTMLElement) {
			// 解析html內容
		if element.Attr("data-src") != "" { //查看要查詢的該2個網頁，其圖片網址都塞在data-src裡，因此解析對象用html裡有屬性data-src的，並且有值不要空的
			count++
			fmt.Println("顯示:", element.Attr("data-src"))
		}
	})

	// url：請求具體的地址
	err := collector.Visit("https://www.greenvines.com.tw")
	err1 := collector.Visit("https://www.greenvines.com.tw/products/scalp-clarifying-shampoo")

	if err != nil || err1 != nil {
		fmt.Println("具體錯誤:", err)
	}
	fmt.Println("count:", count)
}
