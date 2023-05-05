package usecase

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

type CurrentDollar struct{}

func NewCurrentDollar() *CurrentDollar {
	return &CurrentDollar{}
}

func (c *CurrentDollar) Execute() float64 {
	value := float64(0)

	res, err := http.Get("https://www.bcb.gov.br/api/servico/sitebcb/indicadorCambio")
	if err != nil {
		c.logError(err)
		return value
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		c.logError(err)
		return value
	}

	var exchange Exchange
	if err = json.Unmarshal(content, &exchange); err != nil {
		c.logError(err)
		return value
	}

	for i := 0; i < len(exchange.Content); i++ {
		if exchange.Content[i].Coin == "Dólar" && exchange.Content[i].Type == "Fechamento" {
			value = exchange.Content[i].BuyValue
			if exchange.Content[i].Type == "Fechamento" {
				break
			}
		}
	}

	if value == float64(0) {
		c.logError(fmt.Errorf("value not found"))
	}

	return value
}

func (c *CurrentDollar) logError(err error) {
	fmt.Println("current_dollar", err)
}
