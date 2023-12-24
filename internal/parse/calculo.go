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
// é preciso localizar a linha com o texto VALOR DEVIDO e depois voltar até a
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
