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

func (p *PurchaseValues) Execute() (map[string]any, error) {
	var err error
	values := make(map[string]any)

	values["salto_del_guaira"], err = p.getValueSaltoDelGuaira()
	values["compras_paraguai"], err = p.getValueComprasParaguai()
	if err != nil {
		return values, err
	}

	return values, nil
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

func (p *PurchaseValues) getValueComprasParaguai() (value float64, err error) {
	res, err := http.Get("https://www.comprasparaguai.com.br/")
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

	fmt.Println(exchange)

	return float64(10), nil
}

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
