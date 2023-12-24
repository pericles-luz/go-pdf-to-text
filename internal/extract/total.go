package extract

import "github.com/pericles-luz/go-base/pkg/utils"

type Total struct {
	valorCorrigido uint64
	valorJurosMora uint64
	totalDevido    uint64
}

func NewTotal() *Total {
	return &Total{}
}

func (t *Total) ValorCorrigido() uint64 {
	return t.valorCorrigido
}

func (t *Total) ValorJurosMora() uint64 {
	return t.valorJurosMora
}

func (t *Total) TotalDevido() uint64 {
	return t.totalDevido
}

func (t *Total) SetValorCorrigido(valorCorrigido string) {
	t.valorCorrigido = uint64(utils.StringToInt(utils.GetOnlyNumbers(valorCorrigido)))
}

func (t *Total) SetValorJurosMora(valorJurosMora string) {
	t.valorJurosMora = uint64(utils.StringToInt(utils.GetOnlyNumbers(valorJurosMora)))
}

func (t *Total) SetTotalDevido(totalDevido string) {
	t.totalDevido = uint64(utils.StringToInt(utils.GetOnlyNumbers(totalDevido)))
}

func (t *Total) Validate() error {
	if t.ValorCorrigido() == 0 {
		return ErrValorCorrigidoInvalido
	}
	if t.ValorJurosMora() == 0 {
		return ErrValorJurosMoraInvalido
	}
	if t.TotalDevido() != t.ValorCorrigido()+t.ValorJurosMora() {
		return ErrTotalDevidoInvalido
	}
	return nil
}
