package client

import (
	"github.com/YuanJey/commonHttpClient/pkg/params"
	"net/http"
)

type HttpClient struct{}

func (c *HttpClient) Request(p *params.RequestConfig) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(p.Method, p.Url, p.Params())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", p.ContentType)
	for key, value := range p.Headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if p.PageConf.IsPage {
		p.PageConf.AddPage()
	}
	return resp, nil
}
