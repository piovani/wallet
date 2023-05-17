package dollar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Base struct{}

func NewBase() *Base {
	return &Base{}
}

type Exchange struct {
	Content []struct {
		Coin      string    `json:"moeda"`
		BuyValue  float32   `json:"valorCompra"`
		SaleValue float32   `json:"valorVenda"`
		Date      time.Time `json:"dataIndicador"`
		Type      string    `json:"tipoCotacao"`
	} `json:"conteudo"`
}

func (b *Base) GetValue() (value float32, err error) {
	res, err := http.Get("https://www.bcb.gov.br/api/servico/sitebcb/indicadorCambio")
	if err != nil {
		return value, err
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return value, err
	}

	var exchange Exchange
	if err = json.Unmarshal(content, &exchange); err != nil {
		return value, err
	}

	for i := 0; i < len(exchange.Content); i++ {
		if exchange.Content[i].Coin == "Dólar" && exchange.Content[i].Type == "Fechamento" {
			value = exchange.Content[i].BuyValue
			if exchange.Content[i].Type == "Fechamento" {
				break
			}
		}
	}

	if value == float32(0) {
		return value, fmt.Errorf("value not found")
	}

	return value, nil
}
