package dollar

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/piovani/wallet/infra/http"
)

type ResponseCaixa []struct {
	Day   string `json:"dia"`
	Value string `json:"valor"`
}

type Caixa struct{}

func NewCaixa() *Caixa {
	return &Caixa{}
}

func (c *Caixa) GetValue() (value any, err error) {
	content, err := http.NewHttp().Get("https://servicebus.caixa.gov.br/cotacoes/api/CotacaoDolar/BuscarCotacaoDolar")
	if err != nil {
		return value, err
	}

	var response ResponseCaixa
	if err = json.Unmarshal(content, &response); err != nil {
		return value, err
	}

	subString := response[0].Value[3:]
	valueString := strings.Replace(subString, ",", ".", 1)

	value, err = strconv.ParseFloat(valueString, 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
