package extract

import "errors"

var (
	mes = map[string]string{
		"jan": "01",
		"fev": "02",
		"mar": "03",
		"abr": "04",
		"mai": "05",
		"jun": "06",
		"jul": "07",
		"ago": "08",
		"set": "09",
		"out": "10",
		"nov": "11",
		"dez": "12",
	}
	ErrLinhaJaExistente         = errors.New("linha ja existente")
	ErrMesAnoInvalido           = errors.New("mes/ano invalido")
	ErrVencimentoBasicoInvalido = errors.New("vencimento basico invalido")
	ErrSomaInvalida             = errors.New("soma invalida")
	ErrPercentualInvalido       = errors.New("percentual invalido")
	ErrValorDevidoInvalido      = errors.New("valor devido invalido")
	ErrValorCorrigidoInvalido   = errors.New("valor corrigido invalido")
	ErrValorJurosMoraInvalido   = errors.New("valor juros mora invalido")
	ErrTotalDevidoInvalido      = errors.New("total devido invalido")
	ErrIndiceCorrecaoInvalido   = errors.New("indice correcao invalido")
	ErrJurosMoraInvalido        = errors.New("juros mora invalido")
)
