package dollar

import (
	"fmt"
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

	values["compras_paraguai"], err = p.getValueComprasParaguai()
	p.checkErr(errs, err, "Compras Paraguai")

	values["salto_del_guaira"], err = p.getValueSaltoDelGuaira()
	p.checkErr(errs, err, "Salto Del Guaira")

	return values, errs
}

func (p *PurchaseValues) getValueDollarBase() (value float32, err error) {
	return NewBase().GetValue()
}

func (p *PurchaseValues) getValueComprasParaguai() (value float64, err error) {
	return NewComprasParaguai().GetValue()
}

func (p *PurchaseValues) getValueSaltoDelGuaira() (value float64, err error) {
	return NewSaltoDelGuaira().GetValue()
}

// Logs
func (p *PurchaseValues) checkErr(errs map[string]string, err error, msg string) {
	if err != nil {
		errs[msg] = err.Error()
		fmt.Println(fmt.Sprintf("fuction error: %s", msg), err)
	}
}
