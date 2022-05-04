package httpClient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"nft_object/library/helper"
	"nft_object/library/logge"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	Timeout = 5
)

// Get get request
// options[0]["data"]  map[string]string data参数
// options[0]["header"]  map[string]string header参数
// options[0]["timeout"] int timeout参数
// options[0]["basicAuth"] []string{"user","password"}
func Get(requestUrl string, options ...map[string]interface{}) (string, error) {
	c := ghttp.NewClient()
	timeout := Timeout
	nowTime := gtime.Now()
	if len(options) >= 1 {
		if _, ok := options[0]["data"]; ok {
			data := gconv.MapStrStr(options[0]["data"])
			q := url.Values{}
			for k, v := range data {
				q.Add(k, v)
			}
			if strings.Contains(requestUrl, "?") {
				requestUrl += "&" + q.Encode()
			} else {
				requestUrl += "?" + q.Encode()
			}
		}
		if _, ok := options[0]["header"]; ok {
			header := gconv.MapStrStr(options[0]["header"])
			for k, v := range header {
				c.SetHeader(k, v)
			}
		}
		if _, ok := options[0]["timeout"]; ok {
			timeout = gconv.Int(options[0]["timeout"])
		}
		if _, ok := options[0]["basicAuth"]; ok {
			basicAuth := gconv.SliceStr(options[0]["basicAuth"])
			if len(basicAuth) == 2 {
				c = c.BasicAuth(basicAuth[0], basicAuth[1])
			}
		}
	}
	c.SetTimeout(time.Second * time.Duration(timeout))
	var res string
	var err error
	level := "info"
	errInfo := ""
	if r, err := c.Get(requestUrl); err != nil {
		level = "error"
		errInfo = err.Error()
	} else {
		defer r.Body.Close()
		if r.StatusCode != http.StatusOK {
			level = "error"
			err = errors.New("request status " + gconv.String(r.StatusCode) + " | error:" + r.RawResponse())
			errInfo = err.Error()
		}
		res = r.ReadAllString()
	}
	useTime := gtime.Now().Sub(nowTime).Seconds()
	logge.Write("http_client", level, useTime, requestUrl, strings.ReplaceAll(errInfo, "\n", ""))
	return res, err
}

// Post post request
// options[0]["header"]  map[string]string header参数
// options[0]["timeout"] int timeout参数
// options[0]["basicAuth"] []string{"user","password"}
func Post(requestUrl string, params interface{}, options ...map[string]interface{}) (string, error) {
	c := ghttp.NewClient()
	timeout := Timeout
	nowTime := gtime.Now()
	if len(options) >= 1 {
		if _, ok := options[0]["header"]; ok {
			header := gconv.MapStrStr(options[0]["header"])
			for k, v := range header {
				c.SetHeader(k, v)
			}
		}
		if _, ok := options[0]["timeout"]; ok {
			timeout = gconv.Int(options[0]["timeout"])
		}
		if _, ok := options[0]["basicAuth"]; ok {
			basicAuth := gconv.SliceStr(options[0]["basicAuth"])
			if len(basicAuth) == 2 {
				c = c.BasicAuth(basicAuth[0], basicAuth[1])
			}
		}
	}
	c.SetTimeout(time.Second * time.Duration(timeout))
	var err error
	res := ""
	level := "info"
	errInfo := ""

	if r, err := c.Post(requestUrl, params); err != nil {
		errInfo = err.Error()
		level = "error"
	} else {
		defer r.Body.Close()
		if r.StatusCode != http.StatusOK {
			err = errors.New("request status " + gconv.String(r.StatusCode) + " | error:" + r.RawResponse())
			errInfo = err.Error()
			level = "error"
		}
		res = r.ReadAllString()
	}
	useTime := gtime.Now().Sub(nowTime).Seconds()
	logge.Write("http_client", level, useTime, requestUrl,
		params, strings.ReplaceAll(errInfo, "\n", ""))
	return res, err
}

// PostForm post form请求
// attachments参数是附件，多个传递如：map[string]string{"filename1":"/file/xxx.jpg","filename2":"/file/xxx.png",}
func PostForm(requestUrl string, params url.Values, header map[string]string,
	attachments ...map[string]string) (interface{}, error) {
	var req *http.Request
	var resp *http.Response
	var err error
	nowTime := gtime.Now()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range params {
		_ = writer.WriteField(k, v[0])
	}
	if len(attachments) > 0 {
		attachment := attachments[0]
		for k, v := range attachment {
			fw, _ := writer.CreateFormFile(k, v)
			fd, _ := os.Open(v)
			defer fd.Close()
			io.Copy(fw, fd)
		}
	}
	if err := writer.Close(); err != nil {
		return "", errors.New(err.Error())
	}

	req, err = http.NewRequest("POST", requestUrl, body)
	if err != nil {
		return nil, err
	}
	// 签名头部
	for k, v := range header {
		req.Header.Add(k, v)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 10,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Second,
				DualStack: true,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
			IdleConnTimeout:       5 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: time.Duration(Timeout) * time.Second,
	}
	resp, err = client.Do(req)
	useTime := gtime.Now().Sub(nowTime).Seconds()
	level := "info"
	errInfo := ""
	if err != nil {
		level = "error"
		errInfo = err.Error()
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			level = "error"
			err = errors.New("request status " + gconv.String(resp.StatusCode))
			errInfo = err.Error()
		}
		_, err = body.ReadFrom(resp.Body)
	}
	logge.Write("http_client", level, useTime, requestUrl,
		params, strings.ReplaceAll(errInfo, "\n", ""))
	return body.String(), err
}

// BacthRequest 并发请求接口示例
func BacthRequest(data []map[string]interface{}) (interface{}, error) {
	ret := []interface{}{}
	w := sync.WaitGroup{}
	for _, row := range data {
		w.Add(1)
		go func(d map[string]interface{}) {

			// header回调外部方法处理并返回参数应用请求
			h := map[string]string{}
			if d["handleHeader"] != nil {
				hv, _ := helper.Call(d["handleHeader"], d["name"])
				for _, v := range hv {
					h = v.Interface().(map[string]string)
				}
			}
			req, err := Post(d["url"].(string), d["params"], map[string]interface{}{"header": h})
			if err != nil {
				// 请求错误或者超时处理
				fmt.Println(err.Error(), "====post error")
				w.Done()
				return
			}
			ret = append(ret, req)

			if d["children"] != nil {
				sub := d["children"].([]map[string]interface{})

				for _, row2 := range sub {

					p2 := row2["params"].(string)
					pparams, err := gjson.DecodeToJson(req)
					if err != nil {
						fmt.Println(err.Error(), "==========error===========")
					}
					if row2["match"] != nil {
						m := row2["match"].(map[string]string)
						for k2, v2 := range m {
							p2 += "&" + k2 + "=" + pparams.GetString(v2)
						}
					}
					row2["params"] = p2
				}
				req, _ := BacthRequest(sub)
				ret = append(ret, req)
			}
			w.Done()
		}(row)
	}
	w.Wait()
	fmt.Println("do ok")
	return ret, nil
}
