package httpclient

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

var HttpClient *http.Client

func Req(method string, baseURL string, url string, data []byte, setters ...func(*http.Request)) (
	body []byte, header http.Header, statusCode int) {

	fullUrl := baseURL + url
	request, err := http.NewRequest(method, fullUrl, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	if method == "POST" || method == "PUT" || method == "DELETE" {
		request.Header.Set("Content-Type", "application/json;charset=utf8;")
		// request.Header.Set("Content-Type", "multipart/form-data;charset=utf8;")
	}
	return reqInner(request, setters...)
}

func reqInner(request *http.Request, setters ...func(*http.Request)) (
	body []byte, header http.Header, statusCode int) {

	for _, setter := range setters {
		setter(request)
	}

	res, err := HttpClient.Do(request)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	header = res.Header
	statusCode = res.StatusCode
	return
}

// createHTTPClient for connection re-use
func CreateHTTPClient(interval int, proxyUrl string) *http.Client {

	timeout := time.Duration(interval) * time.Second

	if proxyUrl != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: proxy,

				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 50,
				IdleConnTimeout:     time.Duration(60) * time.Second,
			},
			Timeout: timeout,
		}
		return client
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     time.Duration(60) * time.Second,
		},
		Timeout: timeout,
	}
	return client

}

func Init() {
	if HttpClient != nil {
		return
	}
	HttpClient = CreateHTTPClient(30, "")

}
