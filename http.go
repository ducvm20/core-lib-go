package main

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

type Req struct {
	Url     string
	IsGet   bool
	Headers map[string]string
	Params  map[string]string
	Body    map[string]interface{}
	Timeout int
	Proxy   bool
}

func Request(req *Req) (*resty.Response, error) {
	if req != nil {
		req.Headers["Content-Type"] = "application/json"
		client := resty.New()
		client.SetTimeout(time.Second * time.Duration(req.Timeout))
		if req.Proxy {
			client.SetProxy("http://proxy.hcm.fpt.vn:80")
		}
		makeReq := client.R().SetHeaders(req.Headers).SetQueryParams(req.Params).
			SetBody(req.Body)
		if req.IsGet {
			return makeReq.Get(req.Url)
		}
		return makeReq.Post(req.Url)
	}
	return nil, errors.New("Request is null")
}
