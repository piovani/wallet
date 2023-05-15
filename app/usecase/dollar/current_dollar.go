package dollar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Exchange struct {
	Content []struct {
		Coin      string    `json:"moeda"`
		BuyValue  float64   `json:"valorCompra"`
		SaleValue float64   `json:"valorVenda"`
		Date      time.Time `json:"dataIndicador"`
		Type      string    `json:"tipoCotacao"`
	} `json:"conteudo"`
}

type Value struct {
	Value float64 `json:"divisa"`
}

type CurrentDollar struct{}

func NewCurrentDollar() *CurrentDollar {
	return &CurrentDollar{}
}

func (c *CurrentDollar) Execute() (float64, error) {
	value := float64(0)

	res, err := http.Get("https://www.bcb.gov.br/api/servico/sitebcb/indicadorCambio")
	if err != nil {
		return value, err
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return value, err
	}

	var v Value
	if err = json.Unmarshal(content, &v); err != nil {
		return value, err
	}

	fmt.Println(v)

	// for i := 0; i < len(exchange.Content); i++ {
	// 	if exchange.Content[i].Coin == "Dólar" && exchange.Content[i].Type == "Fechamento" {
	// 		value = exchange.Content[i].BuyValue
	// 		if exchange.Content[i].Type == "Fechamento" {
	// 			break
	// 		}
	// 	}
	// }

	return value, nil
}
