package service

import (
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
)

type Calculo struct {
	calculo *application.Calculo
	lines   []string
}

func NewCalculo() *Calculo {
	return &Calculo{}
}

func (c *Calculo) Parse(path string) error {

	if err := c.loadFile(path); err != nil {
		return err
	}
	c.calculo = application.NewCalculo()
	if err := parse.CalculoBase(c.lines, c.calculo); err != nil {
		return err
	}
	if err := parse.Desagio35(c.lines, c.calculo); err != nil {
		return err
	}
	if err := parse.TotalAposDesagio35(c.lines, c.calculo); err != nil {
		return err
	}
	if err := parse.Total(c.lines, c.calculo); err != nil {
		return err
	}
	if err := parse.Table(c.lines, c.calculo); err != nil {
		return err
	}
	if c.calculo.Total().TotalDevido() == 0 {
		return application.ErrCalculoNaoEncontrado
	}
	if extract.MuitoDiferente(c.calculo.TotalDevido(), c.calculo.Total().TotalDevido()) {
		return application.ErrCalculoInconsistente
	}
	if err := c.calculo.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *Calculo) loadFile(path string) error {
	if !utils.FileExists(path) {
		return application.ErrArquivoNaoEncontrado
	}
	lines, err := extract.ReadLinesFromFile(path)
	if err != nil {
		return err
	}
	c.lines = lines
	return nil
}
