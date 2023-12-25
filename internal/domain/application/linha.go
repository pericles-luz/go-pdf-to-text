package application

import (
	"time"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
)

type Linha struct {
	mesAno           time.Time
	vencimentoBasico uint64
	soma             uint64
	percentual       uint64
	valorDevido      uint64
	indiceCorrecao   uint64
	valorCorrigido   uint64
	jurosMora        uint64
	valorJurosMora   uint64
	totalDevido      uint64
	tercoFerias      uint64
	// // anuênio, art 244 Lei 8112/90
	// art244Lei8112    uint64
	// // vantagens art 184,inciso II, lei 1711/52
	// art184Lei1711    uint64
}

func NewLinha() *Linha {
	return &Linha{}
}

func (l *Linha) MesAno() string {
	return l.mesAno.Format("01/2006")
}

func (l *Linha) VencimentoBasico() uint64 {
	return l.vencimentoBasico
}

func (l *Linha) TercoFerias() uint64 {
	return l.tercoFerias
}

func (l *Linha) Soma() uint64 {
	return l.soma
}

func (l *Linha) Percentual() uint64 {
	return l.percentual
}

func (l *Linha) ValorDevido() uint64 {
	return l.valorDevido
}

func (l *Linha) IndiceCorrecao() uint64 {
	return l.indiceCorrecao
}

func (l *Linha) ValorCorrigido() uint64 {
	return l.valorCorrigido
}

func (l *Linha) JurosMora() uint64 {
	return l.jurosMora
}

func (l *Linha) ValorJurosMora() uint64 {
	return l.valorJurosMora
}

func (l *Linha) TotalDevido() uint64 {
	return l.totalDevido
}

func (l *Linha) SetMesAno(mesAno string) {
	l.mesAno, _ = time.Parse("01/2006/02", mes[mesAno[0:3]]+"/19"+mesAno[4:6]+"/01")
}

func (l *Linha) SetVencimentoBasico(vencimentoBasico string) {
	l.vencimentoBasico = uint64(utils.StringToInt(utils.GetOnlyNumbers(vencimentoBasico)))
}

func (l *Linha) SetTercoFerias(tercoFerias string) {
	l.tercoFerias = uint64(utils.StringToInt(utils.GetOnlyNumbers(tercoFerias)))
}

func (l *Linha) SetSoma(soma string) {
	l.soma = uint64(utils.StringToInt(utils.GetOnlyNumbers(soma)))
}

func (l *Linha) SetPercentual(percentual string) {
	l.percentual = uint64(utils.StringToInt(utils.GetOnlyNumbers(percentual)))
}

func (l *Linha) SetValorDevido(valorDevido string) {
	l.valorDevido = uint64(utils.StringToInt(utils.GetOnlyNumbers(valorDevido)))
}

func (l *Linha) SetIndiceCorrecao(indiceCorrecao string) {
	l.indiceCorrecao = uint64(utils.StringToInt(utils.GetOnlyNumbers(indiceCorrecao)))
}

func (l *Linha) SetValorCorrigido(valorCorrigido string) {
	l.valorCorrigido = uint64(utils.StringToInt(utils.GetOnlyNumbers(valorCorrigido)))
}

func (l *Linha) SetJurosMora(jurosMora string) {
	l.jurosMora = uint64(utils.StringToInt(utils.GetOnlyNumbers(jurosMora)))
}

func (l *Linha) SetValorJurosMora(valorJurosMora string) {
	l.valorJurosMora = uint64(utils.StringToInt(utils.GetOnlyNumbers(valorJurosMora)))
}

func (l *Linha) SetTotalDevido(totalDevido string) {
	l.totalDevido = uint64(utils.StringToInt(utils.GetOnlyNumbers(totalDevido)))
}

func (l *Linha) Validate() error {
	if l.mesAno.IsZero() {
		return ErrMesAnoInvalido
	}
	if l.VencimentoBasico() == 0 {
		return ErrVencimentoBasicoInvalido
	}
	// if l.Soma() != l.VencimentoBasico()+l.TercoFerias() {
	// 	return ErrSomaInvalida
	// }
	if l.Percentual() == 0 {
		return ErrPercentualInvalido
	}
	if extract.MuitoDiferente(l.ValorDevido(), l.Soma()*l.Percentual()/10000) {
		return ErrValorDevidoInvalido
	}
	if l.IndiceCorrecao() == 0 {
		return ErrIndiceCorrecaoInvalido
	}
	if extract.MuitoDiferente(l.ValorCorrigido(), l.ValorDevido()*l.IndiceCorrecao()/100000000) {
		return ErrValorCorrigidoInvalido
	}
	if l.JurosMora() == 0 {
		return ErrJurosMoraInvalido
	}
	if extract.MuitoDiferente(l.ValorJurosMora(), l.ValorCorrigido()*l.JurosMora()/1000000) {
		return ErrValorJurosMoraInvalido
	}
	if extract.MuitoDiferente(l.TotalDevido(), l.ValorCorrigido()+l.ValorJurosMora()) {
		return ErrTotalDevidoInvalido
	}
	return nil
}
