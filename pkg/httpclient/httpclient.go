package httpclient

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type HttpClient struct {
	baseUrl       string
	headers       map[string]string
	allowInsecure bool
}

type HeaderOptions map[string]string

func httpRequest(session HttpClient, method, route string, body io.Reader) *http.Request {
	base := session.baseUrl
	if match, _ := regexp.MatchString("http[s]?\\:\\/\\/", session.baseUrl); !match {
		base = "https://" + base
	}

	url := fmt.Sprintf("%s%s", base, route)
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		msg := fmt.Sprintf("Failed to create new HTTP Request: %d", err)
		panic(msg)
	}

	for key, val := range session.headers {
		req.Header.Add(key, val)
	}

	return req

}

func (s *HttpClient) Get(route string) (http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: s.allowInsecure},
	}
	client := http.Client{Transport: tr}
	method := http.MethodGet
	request := httpRequest(*s, method, route, nil)

	res, err := client.Do(request)
	if err != nil {
		msg := fmt.Errorf("failed HTTP Request: %s - %s - %s: Error %w", method, s.baseUrl, route, err)
		return http.Response{}, msg
	}
	return *res, nil
}

func (s *HttpClient) Post(route string, body io.Reader) (http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: s.allowInsecure},
	}
	client := http.Client{Transport: tr}

	method := http.MethodPost
	request := httpRequest(*s, method, route, body)

	res, err := client.Do(request)

	if err != nil {
		msg := fmt.Errorf("failed HTTP Request: %s - %s - %s: Error %w", method, s.baseUrl, route, err)
		return http.Response{}, msg
	}

	return *res, nil

}

func AddHeader(client *HttpClient, key, value string) {
	client.headers[key] = value
}

/*
Build and return a new HttpClient.

@params:

	baseUrl string
	token string
	headers? map[string]string
*/
func NewHttpClient(baseUrl string, allowInsecure bool, headerOpts ...HeaderOptions) HttpClient {
	h := HttpClient{baseUrl: baseUrl, headers: HeaderOptions{}}
	h.baseUrl = baseUrl
	h.allowInsecure = allowInsecure
	for _, headers := range headerOpts {
		for k, v := range headers {
			h.headers[k] = v
		}
	}
	return h
}
