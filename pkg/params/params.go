package params

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type RequestConfig struct {
	Method      string                 `yaml:"method" json:"method,omitempty"`
	BodyType    int                    `yaml:"bodyType" json:"bodyType,omitempty"`
	Url         string                 `yaml:"url" json:"url,omitempty"`
	ContentType string                 `yaml:"contentType" json:"contentType,omitempty"`
	Headers     map[string]string      `yaml:"headers" json:"headers,omitempty"`
	Req         map[string]interface{} `yaml:"req" json:"req,omitempty"`
	PageConf    PageConfig             `yaml:"pageConf" json:"pageConf"`
}
type PageConfig struct {
	IsPage bool `json:"isPage,omitempty" yaml:"isPage"`
	//当前页
	Page      int    `json:"page,omitempty" yaml:"page"`
	PageField string `json:"pageField,omitempty" yaml:"pageField"`
	PageSize  int    `json:"pageSize,omitempty" yaml:"pageSize"`
	SizeField string `json:"sizeField,omitempty" yaml:"sizeField"`
}

func (p *PageConfig) AddPage() {
	if p.IsPage {
		p.Page++
	}
}
func (r *RequestConfig) Params() io.Reader {
	if r.Req == nil {
		return nil
	}
	switch r.BodyType {
	case BodyTypeQuery:
		//初始化url
		if strings.Contains(r.Url, "?") {
			r.Url = strings.Split(r.Url, "?")[0]
		}
		values := url.Values{}
		if r.PageConf.IsPage {
			for k, v := range r.Req {
				switch reflect.TypeOf(v).String() {
				case "string":
					if k == r.PageConf.PageField {
						values.Add(k, strconv.Itoa(r.PageConf.Page))
					} else if k == r.PageConf.SizeField {
						values.Add(k, strconv.Itoa(r.PageConf.PageSize))
					} else {
						values.Add(k, v.(string))
					}
				case "int":
					if k == r.PageConf.PageField {
						values.Add(k, strconv.Itoa(r.PageConf.Page))
					} else if k == r.PageConf.SizeField {
						values.Add(k, strconv.Itoa(r.PageConf.PageSize))
					} else {
						values.Add(k, strconv.Itoa(v.(int)))
					}
				}
			}
		} else {
			for k, v := range r.Req {
				switch reflect.TypeOf(v).String() {
				case "string":
					values.Add(k, v.(string))
				case "int":
					values.Add(k, strconv.Itoa(v.(int)))
				}
			}
		}
		r.Url = r.Url + "?" + values.Encode()
		return nil
	case BodyTypeJson:
		if r.Req != nil {
			if r.PageConf.IsPage {
				for k, v := range r.Req {
					switch reflect.TypeOf(v).String() {
					case "string":
						if k == r.PageConf.PageField {
							r.Req[k] = strconv.Itoa(r.PageConf.Page)
						} else if k == r.PageConf.SizeField {
							r.Req[k] = strconv.Itoa(r.PageConf.PageSize)
						}
					case "int":
						if k == r.PageConf.PageField {
							r.Req[k] = strconv.Itoa(r.PageConf.Page)
						} else if k == r.PageConf.SizeField {
							r.Req[k] = strconv.Itoa(r.PageConf.PageSize)
						}
					}
				}
				jsonStr, err := json.Marshal(r.Req)
				if err != nil {
					return nil
				}
				fmt.Println(string(jsonStr))
				return strings.NewReader(string(jsonStr))
			} else {
				jsonStr, err := json.Marshal(r.Req)
				if err != nil {
					return nil
				}
				return strings.NewReader(string(jsonStr))
			}
		}
	case BodyTypeFormUrlencoded:
		values := url.Values{}
		if r.PageConf.IsPage {
			for k, v := range r.Req {
				switch reflect.TypeOf(v).String() {
				case "string":
					if k == r.PageConf.PageField {
						values.Add(k, strconv.Itoa(r.PageConf.Page))
					} else if k == r.PageConf.SizeField {
						values.Add(k, strconv.Itoa(r.PageConf.PageSize))
					} else {
						values.Add(k, v.(string))
					}
				case "int":
					if k == r.PageConf.PageField {
						values.Add(k, strconv.Itoa(r.PageConf.Page))
					} else if k == r.PageConf.SizeField {
						values.Add(k, strconv.Itoa(r.PageConf.PageSize))
					} else {
						values.Add(k, strconv.Itoa(v.(int)))
					}
				}
			}
			return strings.NewReader(values.Encode())
		} else {
			for k, v := range r.Req {
				switch reflect.TypeOf(v).String() {
				case "string":
					values.Add(k, v.(string))
				case "int":
					values.Add(k, strconv.Itoa(v.(int)))
				}
			}
			return strings.NewReader(values.Encode())
		}
	case BodyTypeFormData:
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		if r.PageConf.IsPage {
			for k, v := range r.Req {
				if k == r.PageConf.PageField {
					r.Req[k] = strconv.Itoa(r.PageConf.Page)
				} else if k == r.PageConf.SizeField {
					r.Req[k] = strconv.Itoa(r.PageConf.PageSize)
				} else {
					r.Req[k] = v.(string)
				}
				//switch reflect.TypeOf(v).String() {
				//case "string":
				//	if k == r.PageConf.PageField {
				//		r.Req[k] = strconv.Itoa(r.PageConf.Page)
				//	} else if k == r.PageConf.SizeField {
				//		r.Req[k] = strconv.Itoa(r.PageConf.PageSize)
				//	} else {
				//		r.Req[k] = v.(string)
				//	}
				//case "int":
				//	if k == r.PageConf.PageField {
				//		r.Req[k] = strconv.Itoa(r.PageConf.Page)
				//	} else if k == r.PageConf.SizeField {
				//		r.Req[k] = strconv.Itoa(r.PageConf.PageSize)
				//	} else {
				//		r.Req[k] = strconv.Itoa(v.(int))
				//	}
				//}
				err := writer.WriteField(k, r.Req[k].(string))
				if err != nil {
					return nil
				}
			}
			err := writer.Close()
			if err != nil {
				return nil
			}
			return body
		} else {
			for k, v := range r.Req {
				err := writer.WriteField(k, v.(string))
				if err != nil {
					return nil
				}
			}
			err := writer.Close()
			if err != nil {
				return nil
			}
			return body
		}
	}
	return nil
}
