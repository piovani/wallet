package dollar

import (
	"fmt"
	"sync"
)

type PurchaseValues struct{}

func NewPurchaseValues() *PurchaseValues {
	return &PurchaseValues{}
}

func (p *PurchaseValues) Execute() (map[string]any, map[string]any) {
	values := make(map[string]any)
	errs := make(map[string]any)

	wg := sync.WaitGroup{}
	wg.Add(4)

	go p.getValueDollarBase(&wg, values, errs)
	go p.getValueComprasParaguai(&wg, values, errs)
	go p.getValueSaltoDelGuaira(&wg, values, errs)
	go p.getValueNomad(&wg, values, errs)

	wg.Wait()

	return values, errs
}

func (p *PurchaseValues) getValueDollarBase(wg *sync.WaitGroup, values, errs map[string]any) {
	prefix := "dollar_base"
	values[prefix], errs[prefix] = NewBase().GetValue()
	wg.Done()
}

func (p *PurchaseValues) getValueComprasParaguai(wg *sync.WaitGroup, values, errs map[string]any) {
	prefix := "compras_paraguai"
	values[prefix], errs[prefix] = NewComprasParaguai().GetValue()
	wg.Done()
}

func (p *PurchaseValues) getValueSaltoDelGuaira(wg *sync.WaitGroup, values, errs map[string]any) {
	prefix := "salto_del_guaira"
	values[prefix], errs[prefix] = NewSaltoDelGuaira().GetValue()
	wg.Done()
}

func (p *PurchaseValues) getValueNomad(wg *sync.WaitGroup, values, errs map[string]any) {
	prefix := "nomad"
	values[prefix], errs[prefix] = NewNomad().GetValue()
	wg.Done()
}

// Logs
func (p *PurchaseValues) checkErr(errs map[string]string, err error, msg string) {
	if err != nil {
		errs[msg] = err.Error()
		fmt.Println(fmt.Sprintf("fuction error: %s", msg), err)
	}
}
