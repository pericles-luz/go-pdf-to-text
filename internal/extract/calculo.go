package extract

import (
	"time"
)

type Calculo struct {
	processoNumero     string
	processoPrincipal  string
	ajuizamento        time.Time
	citacao            time.Time
	calculo            time.Time
	exequente          string
	cpf                string
	idUnica            string
	tabela             []*Linha
	total              *Total
	desagio35          *Total
	totalAposDesagio35 *Total
}

func NewCalculo() *Calculo {
	return &Calculo{}
}

func (c *Calculo) SetProcessoNumero(processoNumero string) {
	c.processoNumero = processoNumero
}

func (c *Calculo) SetProcessoPrincipal(processoPrincipal string) {
	c.processoPrincipal = processoPrincipal
}

func (c *Calculo) SetAjuizamento(ajuizamento string) {
	c.ajuizamento, _ = time.Parse("02/01/2006", ajuizamento)
}

func (c *Calculo) SetCitacao(citacao string) {
	c.citacao, _ = time.Parse("02/01/2006", citacao)
}

func (c *Calculo) SetCalculo(calculo string) {
	c.calculo, _ = time.Parse("02/01/2006", calculo)
}

func (c *Calculo) SetExequente(exequente string) {
	c.exequente = exequente
}

func (c *Calculo) SetCpf(cpf string) {
	c.cpf = cpf
}

func (c *Calculo) SetIdUnica(idUnica string) {
	c.idUnica = idUnica
}

func (c *Calculo) AddLinha(l *Linha) error {
	for _, linha := range c.tabela {
		if linha.MesAno() == l.MesAno() {
			return ErrLinhaJaExistente
		}
	}
	c.tabela = append(c.tabela, l)
	return nil
}

func (c *Calculo) SetTotal(t *Total) {
	c.total = t
}

func (c *Calculo) SetDesagio35(t *Total) {
	c.desagio35 = t
}

func (c *Calculo) SetTotalAposDesagio35(t *Total) {
	c.totalAposDesagio35 = t
}

func (c *Calculo) ProcessoNumero() string {
	return c.processoNumero
}

func (c *Calculo) ProcessoPrincipal() string {
	return c.processoPrincipal
}

func (c *Calculo) Ajuizamento() string {
	return c.ajuizamento.Format("02/01/2006")
}

func (c *Calculo) Citacao() string {
	return c.citacao.Format("02/01/2006")
}

func (c *Calculo) Calculo() string {
	return c.calculo.Format("02/01/2006")
}

func (c *Calculo) Exequente() string {
	return c.exequente
}

func (c *Calculo) Cpf() string {
	return c.cpf
}

func (c *Calculo) IdUnica() string {
	return c.idUnica
}

func (c *Calculo) Tabela() []*Linha {
	return c.tabela
}

func (c *Calculo) Total() *Total {
	return c.total
}

func (c *Calculo) Desagio35() *Total {
	return c.desagio35
}

func (c *Calculo) TotalAposDesagio35() *Total {
	return c.totalAposDesagio35
}
