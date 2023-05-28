package dollar

import (
	"strconv"
	"strings"

	"github.com/piovani/wallet/infra/http"
)

type Nomad struct{}

func NewNomad() *Nomad {
	return &Nomad{}
}

func (n *Nomad) GetValue() (value any, err error) {
	content, err := http.NewHttp().GetToCrawler("https://www.nomadglobal.com/")
	if err != nil {
		return value, err
	}

	init := strings.Index(content, "Cotação agoraR$ ")
	end := strings.Index(content[init:], "\n")

	value, err = strconv.ParseFloat(content[init+18:init+end], 64)
	if err == nil {
		return value, err
	}

	return value, nil
}
