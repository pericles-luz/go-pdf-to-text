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

func FindExecutionNumber(lines []string) (string, error) {
	return findProcessoNumero(lines)
}

func FindMainNumber(lines []string) (string, error) {
	return findProcessoPrincipal(lines)
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
			calculo.SetCpf(clearColon(line))
			calculo.SetExequente(clearColon(lines[i-1]))
			calculo.SetIdUnica(clearColon(lines[i+1]))
			return nil
		}
	}
	return application.ErrExequenteNaoEncontrado
}

// limpa textos no estilo label:  texto e retorna apenas o texto
func clearColon(text string) string {
	colonPosition := strings.Index(text, ":")
	if colonPosition == -1 {
		return strings.TrimSpace(text)
	}
	if colonPosition == len(text)-1 {
		return ""
	}
	return strings.TrimSpace(text[colonPosition+1:])
}
