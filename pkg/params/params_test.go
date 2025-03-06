package params

import (
	"fmt"
	"testing"
)

func TestRequestConfig_Params(t *testing.T) {
	req := make(map[string]interface{})
	req["page"] = "0"
	req["size"] = "0"
	req["name"] = "knight"
	pageConfig := PageConfig{
		IsPage:    true,
		Page:      1,
		PageField: "page",
		PageSize:  100,
		SizeField: "size",
	}
	requestConfig := RequestConfig{
		Method:      "GET",
		BodyType:    4,
		Url:         "https://www.baidu.com",
		ContentType: "application/json",
		Headers:     nil,
		Req:         req,
		PageConf:    pageConfig,
	}
	params1 := requestConfig.Params()
	fmt.Println(params1)
	requestConfig.PageConf.AddPage()
	params2 := requestConfig.Params()
	fmt.Println(params2)
}
