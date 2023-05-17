package dollar

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ComprasParaguai struct{}

func NewComprasParaguai() *ComprasParaguai {
	return &ComprasParaguai{}
}

func (c *ComprasParaguai) GetValue() (value float64, err error) {
	res, err := http.Get("https://www.comprasparaguai.com.br/imposto-dolar/")
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
	init := strings.Index(content, "R$")
	position := strings.Index(content[init:len(content)-1], "\n")

	valueString := strings.Replace(content[init+3:init+position], ",", ".", 1)

	value, err = strconv.ParseFloat(valueString, 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
