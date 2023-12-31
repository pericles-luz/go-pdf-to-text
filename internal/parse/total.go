package parse

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
)

func Total(lines []string, calculo *application.Calculo) error {
	found := false
	total := application.NewTotal()
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "TOTAL       ") {
			found = true
		}
		if !found {
			continue
		}
		if err := readTotal(line, total); err != nil {
			return fmt.Errorf("total inválido: %w", err)
		}
		calculo.SetTotal(total)
		return nil
	}
	return application.ErrTotalNaoEncontrado
}

func Desagio35(lines []string, calculo *application.Calculo) error {
	found := false
	desagio35 := application.NewTotal()
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "DESÁGIO 35%") {
			found = true
		}
		if !found {
			continue
		}
		if err := readTotal(line, desagio35); err != nil {
			return fmt.Errorf("deságio 35 inválido: %w", err)
		}
		calculo.SetDesagio35(desagio35)
		return nil
	}
	return application.ErrDesagio35NaoEncontrado
}

func TotalAposDesagio35(lines []string, calculo *application.Calculo) error {
	found := false
	totalAposDesagio35 := application.NewTotal()
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "TOTAL APÓS DESÁGIO") {
			found = true
		}
		if !found {
			continue
		}
		if err := readTotal(line, totalAposDesagio35); err != nil {
			return fmt.Errorf("total após deságio 35 inválido: %w", err)
		}
		calculo.SetTotalAposDesagio35(totalAposDesagio35)
		return nil
	}
	return application.ErrTotalAposDesagio35NaoEncontrado
}

func readTotal(line string, total *application.Total) error {
	values := regexp.MustCompile(`[\d.,]+`)
	if len(line) < 40 {
		fmt.Println("toal não encontrado em:", line)
		return application.ErrTotalInvalido
	}
	matches := values.FindAllString(line[40:], -1)
	if len(matches) != 3 {
		return application.ErrTotalInvalido
	}
	total.SetValorCorrigido(utils.GetOnlyNumbers(matches[0]))
	total.SetValorJurosMora(utils.GetOnlyNumbers(matches[1]))
	total.SetTotalDevido(utils.GetOnlyNumbers(matches[2]))
	return nil
}
