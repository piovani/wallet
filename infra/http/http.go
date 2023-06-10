package http

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Http struct {
	client *http.Client
}

func NewHttp() *Http {
	return &Http{
		client: &http.Client{
			Timeout: time.Second * 2,
		},
	}
}

func (h *Http) Get(url string) (content []byte, err error) {
	res, err := h.getBase(url)
	if err != nil {
		return content, err
	}

	content, err = io.ReadAll(res.Body)
	if err != nil {
		return content, err
	}

	return content, nil
}

func (h *Http) GetToCrawler(url string) (content string, err error) {
	res, err := h.getBase(url)
	if err != nil {
		return content, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return content, err
	}

	return doc.Text(), nil
}

func (h *Http) getBase(url string) (res *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}

	res, err = h.client.Do(req)
	if err != nil {
		return res, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 && res.StatusCode != 202 {
		return res, fmt.Errorf("request failed")
	}

	return res, nil
}
