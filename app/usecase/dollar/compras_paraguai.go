package dollar

import (
	"strconv"
	"strings"

	"github.com/piovani/wallet/infra/http"
)

type ComprasParaguai struct{}

func NewComprasParaguai() *ComprasParaguai {
	return &ComprasParaguai{}
}

func (c *ComprasParaguai) GetValue() (value any, err error) {
	content, err := http.NewHttp().GetToCrawler("https://www.comprasparaguai.com.br/imposto-dolar/")
	if err != nil {
		return value, err
	}

	init := strings.Index(content, "COTAÇÃO DO DÓLAR\n")
	final := strings.Index(content, "COTAÇÃO\n ")

	sub := content[init:final]
	init = strings.Index(sub, "R$ ")
	final = strings.Index(sub[init:], "\n")

	valueString := strings.Replace(sub[init+3:init+final], ",", ".", 1)

	value, err = strconv.ParseFloat(valueString, 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
