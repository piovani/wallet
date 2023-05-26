package controllers

type PurchaseValuesOut struct {
	DollarBase      any `json:"dollar_base"`
	ComprasParaguai any `json:"compras_paraguai"`
	SaltoDelGuaira  any `json:"salto_del_guaira"`
	Nomad           any `json:"nomad"`
	Caixa           any `json:"caixa"`
}

func NewPurchaseValuesOut(values map[string]any) *PurchaseValuesOut {
	return &PurchaseValuesOut{
		DollarBase:      values["dollar_base"],
		ComprasParaguai: values["compras_paraguai"],
		SaltoDelGuaira:  values["salto_del_guaira"],
		Nomad:           values["nomad"],
		Caixa:           values["caixa"],
	}
}
