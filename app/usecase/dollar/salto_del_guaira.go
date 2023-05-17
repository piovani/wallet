package dollar

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SaltoDelGuaira struct{}

func NewSaltoDelGuaira() *SaltoDelGuaira {
	return &SaltoDelGuaira{}
}

func (a *SaltoDelGuaira) GetValue() (value float64, err error) {
	res, err := http.Get("http://mundialcambios.com.py/?branch=5&lang=pt")
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

	content := doc.Find("h3").Text()
	init := strings.Index(content, "x BRL")
	position := strings.Index(content[init:len(content)-1], "USD x")

	valueString := strings.Replace(content[init+9:init+position], ",", ".", 1)

	value, err = strconv.ParseFloat(valueString, 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
