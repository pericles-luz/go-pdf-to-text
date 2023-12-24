package parse

import (
	"strings"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
)

func Total(lines []string, calculo *application.Calculo) error {
	found := false
	total := application.NewTotal()
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line == "DESÁGIO 35%" {
			found = true
		}
		if !found {
			continue
		}
		total.SetValorCorrigido(strings.TrimSpace(lines[i-2]))
		break
	}
	if total.ValorCorrigido() == 0 {
		return application.ErrValorCorrigidoTotalNaoEncontrado
	}
	foundTotal, foundJurosMora := false, false
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			if foundTotal && total.TotalDevido() > 0 {
				foundJurosMora = true
			}
			continue
		}
		if line == "TOTAL" {
			foundTotal = true
			continue
		}
		if !foundTotal {
			continue
		}
		if total.TotalDevido() == 0 {
			total.SetTotalDevido(strings.TrimSpace(lines[i]))
			continue
		}
		if !foundJurosMora {
			continue
		}
		total.SetValorJurosMora(strings.TrimSpace(lines[i]))
		calculo.SetTotal(total)
		return nil
	}
	return application.ErrTotalNaoEncontrado
}

func Desagio35(lines []string, calculo *application.Calculo) error {
	found := false
	desagio35 := application.NewTotal()
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line == "DESÁGIO 35%" {
			found = true
		}
		if !found {
			continue
		}
		desagio35.SetValorCorrigido(strings.TrimSpace(lines[i+2]))
		desagio35.SetValorJurosMora(strings.TrimSpace(lines[i+4]))
		desagio35.SetTotalDevido(strings.TrimSpace(lines[i+6]))
		calculo.SetDesagio35(desagio35)
		return nil
	}
	return application.ErrDesagio35NaoEncontrado
}

func TotalAposDesagio35(lines []string, calculo *application.Calculo) error {
	found := false
	totalAposDesagio35 := application.NewTotal()
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line == "TOTAL APÓS DESÁGIO" {
			found = true
		}
		if !found {
			continue
		}
		totalAposDesagio35.SetValorCorrigido(strings.TrimSpace(lines[i+2]))
		totalAposDesagio35.SetValorJurosMora(strings.TrimSpace(lines[i+4]))
		totalAposDesagio35.SetTotalDevido(strings.TrimSpace(lines[i+6]))
		calculo.SetTotalAposDesagio35(totalAposDesagio35)
		return nil
	}
	return application.ErrTotalAposDesagio35NaoEncontrado
}
