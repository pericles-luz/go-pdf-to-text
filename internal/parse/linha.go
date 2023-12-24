package parse

import (
	"regexp"
	"strings"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
)

func Linha(lines []string, page, line int, calculo *application.Calculo) error {
	linha := application.NewLinha()
	mesAno, err := findMesAno(lines, page, line)
	if err != nil {
		return err
	}
	if mesAno == "" {
		return application.ErrMesAnoNaoEncontrado
	}
	linha.SetMesAno(mesAno)
	vencimentoBasico, err := findVencimentoBasico(lines, page, line)
	if err != nil {
		return err
	}
	if vencimentoBasico == "" {
		return application.ErrVencimentoBasicoNaoEncontrado
	}
	linha.SetVencimentoBasico(vencimentoBasico)
	soma, err := findSoma(lines, page, line)
	if err != nil {
		return err
	}
	if soma == "" {
		return application.ErrSomaNaoEncontrada
	}
	linha.SetSoma(soma)
	valorDevido, err := findValorDevido(lines, page, line)
	if err != nil {
		return err
	}
	if valorDevido == "" {
		return application.ErrValorDevidoNaoEncontrado
	}
	linha.SetValorDevido(valorDevido)
	calculo.AddLinha(linha)
	return nil
}

func findMesAno(lines []string, page, count int) (string, error) {
	if page == 0 || count == 0 {
		return "", application.ErrMesAnoNaoExiste
	}
	re := regexp.MustCompile(`^(jan|fev|mar|abr|mai|jun|jul|ago|set|out|nov|dez)/\d{2}$`)
	foundPage := 0
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		if line == "Mês/Ano" {
			foundPage++
		}
		if page > foundPage {
			continue
		}
		// se achar nova página antes de achar a linha desejada, para de procurar
		if foundPage > page {
			return "", nil
		}
		if re.MatchString(line) {
			count--
		}
		if count == 0 {
			return line, nil
		}
	}
	return "", application.ErrMesAnoNaoEncontrado
}

func findVencimentoBasico(lines []string, page, count int) (string, error) {
	if page == 0 || count == 0 {
		return "", application.ErrVencimentoBasicoNaoEncontrado
	}
	foundPage := 0
	foundCount := 0
	mesAno := regexp.MustCompile(`^(jan|fev|mar|abr|mai|jun|jul|ago|set|out|nov|dez)/\d{2}$`)
	value := regexp.MustCompile(`,\d{2}$`)
	for i := 0; i < len(lines); i++ {
		if foundPage > page {
			return "", application.ErrVencimentoBasicoNaoEncontrado
		}
		line := strings.TrimSpace(lines[i])
		if mesAno.MatchString(line) {
			continue
		}
		if len(line) == 0 && page != foundPage {
			continue
		}
		if len(line) == 0 && foundCount > 0 {
			return "", application.ErrVencimentoBasicoNaoEncontrado
		}
		if line == "ADICIONAL 1/3 DE" {
			continue
		}
		if line == "FERIAS" {
			continue
		}
		if line == "VENCIMENTO BASICO" {
			foundPage++
			foundCount = 0
			continue
		}
		if value.MatchString(line) {
			foundCount++
		}
		if foundPage == page && foundCount == count {
			return line, nil
		}
	}
	return "", application.ErrVencimentoBasicoNaoEncontrado
}

func findSoma(lines []string, page, count int) (string, error) {
	if page == 0 || count == 0 {
		return "", application.ErrSomaNaoEncontrada
	}
	foundPage := 0
	foundCount := 0
	value := regexp.MustCompile(`,\d{2}$`)
	for i := 0; i < len(lines); i++ {
		if foundPage > page {
			return "", application.ErrSomaNaoEncontrada
		}
		line := strings.TrimSpace(lines[i])
		if strings.Contains(line, "%") {
			continue
		}
		if len(line) == 0 && page != foundPage {
			continue
		}
		if len(line) == 0 && foundCount > 0 {
			return "", application.ErrSomaNaoEncontrada
		}
		if line == "Soma" {
			foundPage++
			foundCount = 0
			if foundPage == page && page == 2 {
				return findSomaPage2(lines, i, count)
			}

			continue
		}
		if page != foundPage {
			continue
		}
		if value.MatchString(line) {
			foundCount++
		}
		if foundCount == count {
			return line, nil
		}
	}
	return "", application.ErrSomaNaoEncontrada
}

// quando a soma etiver na página 2, a busca deve ser feita de forma diferente
// é preciso localizar a linha com o texto "VALOR DEVIDO" e depois voltar até a
// segunda linha vazia. A partir daí, deve-se contar as linhas com valores
// até encontrar a linha desejada
func findSomaPage2(lines []string, index, count int) (string, error) {
	foundCount := 0
	foundText := 0
	value := regexp.MustCompile(`,\d{2}$`)
	for i := index; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "VALOR DEVIDO" {
			foundText = i
			break
		}
	}
	if foundText == 0 {
		return "", application.ErrSomaNaoEncontrada
	}
	spaces := 0
	start := 0
	for i := foundText; i > index; i-- {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			spaces++
		}
		if spaces == 2 {
			start = i + 1
			break
		}
	}
	for i := start; i < foundText; i++ {
		line := strings.TrimSpace(lines[i])
		if value.MatchString(line) {
			foundCount++
		}
		if foundCount == count {
			return line, nil
		}
	}
	return "", application.ErrSomaNaoEncontrada
}

func findValorDevido(lines []string, page, count int) (string, error) {
	if page == 2 {
		return findValorDevidoPage2(lines, page, count)
	}
	foundPage := 0
	value := regexp.MustCompile(`,\d{2}$`)
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		if line == "VALOR DEVIDO" {
			foundPage++
		}
		if page > foundPage {
			continue
		}
		// se achar nova página antes de achar a linha desejada, para de procurar
		if foundPage > page {
			return "", application.ErrValorDevidoNaoEncontrado
		}
		if value.MatchString(line) {
			count--
		}
		if count == 0 {
			return line, nil
		}
	}
	return "", application.ErrValorDevidoNaoEncontrado
}

// na página 2, a busca deve ser feita de forma diferente
// é preciso localizar a linha com o texto "Soma" e pesquisar os valores a partir
// da linha seguinte até encontrar a linha desejada
func findValorDevidoPage2(lines []string, page, count int) (string, error) {
	foundPage := 0
	value := regexp.MustCompile(`,\d{2}$`)
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		if line == "Soma" {
			foundPage++
		}
		if page > foundPage {
			continue
		}
		// se achar nova página antes de achar a linha desejada, para de procurar
		if foundPage > page {
			return "", application.ErrValorDevidoNaoEncontrado
		}
		if value.MatchString(line) {
			count--
		}
		if count == 0 {
			return line, nil
		}
	}
	return "", application.ErrValorDevidoNaoEncontrado
}