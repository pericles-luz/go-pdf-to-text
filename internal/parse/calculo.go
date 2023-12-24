package parse

import (
	"regexp"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
)

func CalculoBase(lines []string, calculo *application.Calculo) error {
	processoNumero, err := findProcessoNumero(lines)
	if err != nil {
		return err
	}
	calculo.SetProcessoNumero(processoNumero)
	processoPrincipal, err := findProcessoPrincipal(lines)
	if err != nil {
		return err
	}
	calculo.SetProcessoPrincipal(processoPrincipal)

	ajuizamento, err := findAjuizamento(lines)
	if err != nil {
		return err
	}
	calculo.SetAjuizamento(ajuizamento)
	citacao, err := findCitacao(lines)
	if err != nil {
		return err
	}
	calculo.SetCitacao(citacao)
	calculos, err := findCalculos(lines)
	if err != nil {
		return err
	}
	calculo.SetCalculo(calculos)
	err = findExequente(lines, calculo)
	if err != nil {
		return err
	}
	return nil
}

func findProcessoNumero(lines []string) (string, error) {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && strings.Contains(line, "Processo nº") {
			re := regexp.MustCompile(`\d{7}-\d{2}.\d{4}.\d{1}.\d{2}.\d{4}`)
			matches := re.FindAllString(line, -1)
			return matches[0], nil
		}
	}
	return "", application.ErrProcessoNumeroNaoEncontrado
}

func findProcessoPrincipal(lines []string) (string, error) {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && strings.Contains(line, "(Principal:") {
			re := regexp.MustCompile(`\d{7}-\d{2}.\d{4}.\d{1}.\d{2}.\d{4}`)
			matches := re.FindAllString(line, -1)
			return matches[len(matches)-1], nil
		}
	}
	return "", application.ErrProcessoPrincipalNaoEncontrado
}

func findAjuizamento(lines []string) (string, error) {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && strings.Contains(line, "Data do Ajuizamento") {
			re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
			matches := re.FindAllString(line, -1)
			return matches[0], nil
		}
	}
	return "", application.ErrAjuizamentoNaoEncontrado
}

func findCitacao(lines []string) (string, error) {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && strings.Contains(line, "Data da Cita") {
			re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
			matches := re.FindAllString(line, -1)
			return matches[0], nil
		}
	}
	return "", application.ErrCitacaoNaoEncontrada
}

func findCalculos(lines []string) (string, error) {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && strings.Contains(line, "Data dos C") {
			re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
			matches := re.FindAllString(line, -1)
			return matches[0], nil
		}
	}
	return "", application.ErrCalculoNaoEncontrado
}

func findExequente(lines []string, calculo *application.Calculo) error {
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		line = utils.GetOnlyNumbers(line)
		if utils.ValidateCPF(line) {
			calculo.SetCpf(line)
			calculo.SetExequente(lines[i-1])
			calculo.SetIdUnica(lines[i+1])
			return nil
		}
	}
	return application.ErrExequenteNaoEncontrado
}

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
