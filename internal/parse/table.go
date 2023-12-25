package parse

import (
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
)

func Table(lines []string, calculo *application.Calculo) error {
	page := 1
	line := 1
	for {
		err := Linha(lines, page, line, calculo)
		if err == application.ErrMesAnoNaoEncontrado {
			if line == 1 {
				return nil
			}
			page++
			line = 1
			continue
		}
		if err != nil {
			return err
		}
		line++
	}
}
