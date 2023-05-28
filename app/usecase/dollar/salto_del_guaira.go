package dollar

import (
	"strconv"
	"strings"

	"github.com/piovani/wallet/infra/http"
)

type SaltoDelGuaira struct{}

func NewSaltoDelGuaira() *SaltoDelGuaira {
	return &SaltoDelGuaira{}
}

func (a *SaltoDelGuaira) GetValue() (value float64, err error) {
	content, err := http.NewHttp().GetToCrawler("http://mundialcambios.com.py/?branch=5&lang=pt")
	if err != nil {
		return value, err
	}

	init := strings.Index(content, "x BRL")
	init += 263
	end := init + 4

	valueString := strings.Replace(content[init:end], ",", ".", 1)
	value, err = strconv.ParseFloat(valueString, 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
