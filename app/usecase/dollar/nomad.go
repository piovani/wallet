package dollar

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Nomad struct{}

func NewNomad() *Nomad {
	return &Nomad{}
}

func (n *Nomad) GetValue() (value any, err error) {
	res, err := http.Get("https://www.nomadglobal.com/")
	if err != nil {
		return value, err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return value, fmt.Errorf("DEU RUIM")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return value, err
	}

	content := doc.Text()
	init := strings.Index(content, "Cotação agoraR$ ")
	end := strings.Index(content[init:], "\n")

	value, err = strconv.ParseFloat(content[init+18:init+end], 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
