package application

import (
	"fmt"
	"time"

	"github.com/pericles-luz/go-base/pkg/utils"
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
	c.cpf = utils.CompleteWithZeros(utils.GetOnlyNumbers(cpf), 11)
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
	if c.ajuizamento.IsZero() {
		return ""
	}
	return c.ajuizamento.Format("02/01/2006")
}

func (c *Calculo) Citacao() string {
	if c.citacao.IsZero() {
		return ""
	}
	return c.citacao.Format("02/01/2006")
}

func (c *Calculo) Calculo() string {
	if c.calculo.IsZero() {
		return ""
	}
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

func (c *Calculo) Validate() error {
	if c.ProcessoNumero() == "" {
		return ErrProcessoNumeroInvalido
	}
	if c.ProcessoPrincipal() == "" {
		return ErrProcessoPrincipalInvalido
	}
	if c.Ajuizamento() == "" {
		return ErrAjuizamentoInvalido
	}
	if c.Citacao() == "" {
		return ErrCitacaoInvalido
	}
	if c.Calculo() == "" {
		return ErrDataCalculoInvalido
	}
	if c.Exequente() == "" {
		return ErrExequenteInvalido
	}
	if !utils.ValidateCPF(c.Cpf()) {
		return ErrCpfInvalido
	}
	if c.IdUnica() == "" {
		return ErrIdUnicaInvalido
	}
	if len(c.Tabela()) == 0 {
		return ErrTabelaInvalida
	}
	for _, linha := range c.Tabela() {
		if err := linha.Validate(); err != nil {
			return fmt.Errorf("linha inválida: %w", err)
		}
	}
	if c.Total() == nil {
		return ErrTotalInvalido
	}
	if c.Desagio35() == nil {
		return ErrDesagio35Invalido
	}
	if c.TotalAposDesagio35() == nil {
		return ErrTotalAposDesagio35Invalido
	}
	if err := c.Total().Validate(); err != nil {
		return fmt.Errorf("total inválido: %w", err)
	}
	if err := c.Desagio35().Validate(); err != nil {
		return fmt.Errorf("desagio 35 inválido: %w", err)
	}
	if err := c.TotalAposDesagio35().Validate(); err != nil {
		return fmt.Errorf("total apos desagio 35 inválido: %w", err)
	}
	if err := c.ValidateTotals(); err != nil {
		return fmt.Errorf("totais inválidos: %w", err)
	}
	return nil
}

func (c *Calculo) ValidateTotals() error {
	if c.Total().TotalDevido() != c.Desagio35().TotalDevido()+c.TotalAposDesagio35().TotalDevido() {
		return ErrTotalInvalido
	}
	if c.Total().ValorCorrigido() != c.Desagio35().ValorCorrigido()+c.TotalAposDesagio35().ValorCorrigido() {
		return ErrValorCorrigidoInvalido
	}
	if c.Total().ValorJurosMora() != c.Desagio35().ValorJurosMora()+c.TotalAposDesagio35().ValorJurosMora() {
		return ErrValorJurosMoraInvalido
	}
	return nil
}
