package dollar

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
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
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://servicebus.caixa.gov.br/cotacoes/api/CotacaoDolar/BuscarCotacaoDolar", nil)
	if err != nil {
		return value, err
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		return value, err
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
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
