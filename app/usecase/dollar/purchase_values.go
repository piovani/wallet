package dollar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type PurchaseValues struct{}

func NewPurchaseValues() *PurchaseValues {
	return &PurchaseValues{}
}

func (p *PurchaseValues) Execute() (map[string]any, map[string]string) {
	values := make(map[string]any)
	errs := make(map[string]string)
	var err error

	values["dollar_base"], err = p.getValueDollarBase()
	p.checkErr(errs, err, "Dolar Base")
	values["salto_del_guaira"], err = p.getValueSaltoDelGuaira()
	p.checkErr(errs, err, "Salto Del Guaira")
	// values["compras_paraguai"], err = p.getValueComprasParaguai()
	// p.checkErr(errs, err, "Compras Paraguai")

	return values, errs
}

type Exchange struct {
	Content []struct {
		Coin      string    `json:"moeda"`
		BuyValue  float64   `json:"valorCompra"`
		SaleValue float64   `json:"valorVenda"`
		Date      time.Time `json:"dataIndicador"`
		Type      string    `json:"tipoCotacao"`
	} `json:"conteudo"`
}

func (p *PurchaseValues) getValueDollarBase() (value float64, err error) {
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

	if value == float64(0) {
		return value, fmt.Errorf("value not found")
	}

	return value, nil
}

type ExchangeComprasParaguai struct {
	Content []struct {
		Coin      string    `json:"esp-cotation"`
		BuyValue  float64   `json:"valorCompra"`
		SaleValue float64   `json:"valorVenda"`
		Date      time.Time `json:"dataIndicador"`
		Type      string    `json:"tipoCotacao"`
	} `json:"conteudo"`
}

// type Exchange struct {
// 	Content []struct {
// 		Coin      string    `json:"moeda"`
// 		BuyValue  float64   `json:"valorCompra"`
// 		SaleValue float64   `json:"valorVenda"`
// 		Date      time.Time `json:"dataIndicador"`
// 		Type      string    `json:"tipoCotacao"`
// 	} `json:"conteudo"`
// }

// func (p *PurchaseValues) getValueComprasParaguai() (value float64, err error) {
// 	res, err := http.Get("https://www.comprasparaguai.com.br/")
// 	if err != nil {
// 		return value, err
// 	}

// 	content, err := io.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		return value, err
// 	}

// 	var exchange Exchange
// 	if err = json.Unmarshal(content, &exchange); err != nil {
// 		return value, err
// 	}

// 	fmt.Println(exchange)

// 	return float64(10), nil
// }

func (p *PurchaseValues) getValueSaltoDelGuaira() (value float64, err error) {
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

	fmt.Println(valueString)

	return value, nil
}

// Logs
func (p *PurchaseValues) checkErr(errs map[string]string, err error, msg string) {
	if err != nil {
		errs[msg] = err.Error()
		fmt.Println(fmt.Sprintf("fuction error: %s", msg), err)
	}
}
