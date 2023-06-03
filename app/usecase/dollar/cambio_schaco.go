package dollar

import (
	"strconv"
	"strings"

	"github.com/piovani/wallet/infra/http"
)

type CambioSchaco struct {
	url string
}

func NewCambioSchaco() *CambioSchaco {
	return &CambioSchaco{
		url: "https://www.cambioschaco.com.py/pt-br/",
	}
}

func (c *CambioSchaco) GetValue() (value any, err error) {
	content, err := http.NewHttp().GetToCrawler(c.url)
	if err != nil {
		return value, err
	}

	init := strings.Index(content, " Dólar   x    Real") + 52
	end := init + 4

	valueString := strings.Replace(content[init:end], ",", ".", 1)

	value, err = strconv.ParseFloat(valueString, 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
