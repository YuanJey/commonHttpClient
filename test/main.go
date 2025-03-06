package main

import (
	"fmt"
	"github.com/YuanJey/commonHttpClient/pkg/client"
	"github.com/YuanJey/commonHttpClient/pkg/params"
)

func main() {
	req := make(map[string]interface{})
	req["page"] = "0"
	req["size"] = "0"
	pageConfig := params.PageConfig{
		IsPage:    true,
		Page:      1,
		PageField: "page",
		PageSize:  100,
		SizeField: "size",
	}
	requestConfig := params.RequestConfig{
		Method:      "GET",
		BodyType:    4,
		Url:         "https://www.baidu.com",
		ContentType: "application/json",
		Headers:     nil,
		Req:         req,
		PageConf:    pageConfig,
	}
	httpClient := client.HttpClient{}
	for i := 0; i < 10; i++ {
		requestConfig.Params()
		httpClient.Request(&requestConfig)
		fmt.Println(requestConfig.Url)
	}
}
