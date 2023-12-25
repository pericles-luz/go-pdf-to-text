package parse

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
)

func Linha(lines []string, page, line int, calculo *application.Calculo) error {
	linha := application.NewLinha()
	found, err := findLine(lines, page, line)
	if err != nil {
		return err
	}
	if err := findMesAno(found, linha); err != nil {
		return fmt.Errorf("erro ao buscar mês/ano: %w", err)
	}
	if err := findVencimentoBasico(found, linha); err != nil {
		return err
	}
	if err := findSoma(found, linha); err != nil {
		return err
	}
	if err := findPercentual(found, linha); err != nil {
		return err
	}
	if err := findValorDevido(found, linha); err != nil {
		return err
	}
	if err := findIndiceCorrecao(found, linha); err != nil {
		return err
	}
	if err := findValorCorrigido(found, linha); err != nil {
		return err
	}
	if err := findJurosMora(found, linha); err != nil {
		return err
	}
	if err := findValorJurosMora(found, linha); err != nil {
		return err
	}
	if err := findTotalDevido(found, linha); err != nil {
		return err
	}
	calculo.AddLinha(linha)
	return nil
}

func findLine(lines []string, page, count int) (string, error) {
	if page == 0 || count == 0 {
		return "", application.ErrMesAnoNaoExiste
	}
	re := regexp.MustCompile(`^(jan|fev|mar|abr|mai|jun|jul|ago|set|out|nov|dez)/\d{2}.*`)
	foundPage := 0
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "Mês/Ano") {
			foundPage++
		}
		if page > foundPage {
			continue
		}
		// se achar nova página antes de achar a linha desejada, para de procurar
		if foundPage > page {
			return "", application.ErrMesAnoNaoEncontrado
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

func findMesAno(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`(jan|fev|mar|abr|mai|jun|jul|ago|set|out|nov|dez)/\d{2}`)
	mesAno := re.FindString(line)
	if mesAno == "" {
		return application.ErrMesAnoNaoEncontrado
	}
	linha.SetMesAno(mesAno)
	return nil
}

func findVencimentoBasico(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 2 {
		return application.ErrVencimentoBasicoNaoEncontrado
	}
	linha.SetVencimentoBasico(values[1])
	return nil
}

func findSoma(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 8 {
		return application.ErrSomaNaoEncontrada
	}
	linha.SetSoma(values[len(values)-8])
	return nil
}

func findPercentual(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 7 {
		return application.ErrPercentualNaoEncontrado
	}
	linha.SetPercentual(values[len(values)-7])
	return nil
}

func findValorDevido(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 6 {
		return application.ErrValorDevidoNaoEncontrado
	}
	linha.SetValorDevido(values[len(values)-6])
	return nil
}

func findIndiceCorrecao(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 5 {
		return application.ErrIndiceCorrecaoNaoEncontrado
	}
	linha.SetIndiceCorrecao(values[len(values)-5])
	return nil
}

func findValorCorrigido(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 4 {
		return application.ErrValorCorrigidoNaoEncontrado
	}
	linha.SetValorCorrigido(values[len(values)-4])
	return nil
}

func findJurosMora(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 3 {
		return application.ErrJurosMoraNaoEncontrado
	}
	linha.SetJurosMora(values[len(values)-3])
	return nil
}

func findValorJurosMora(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 2 {
		return application.ErrValorJurosMoraNaoEncontrado
	}
	linha.SetValorJurosMora(values[len(values)-2])
	return nil
}

func findTotalDevido(line string, linha *application.Linha) error {
	re := regexp.MustCompile(`[\d.,%]+`)
	values := re.FindAllString(line, -1)
	if len(values) < 1 {
		return application.ErrTotalDevidoNaoEncontrado
	}
	linha.SetTotalDevido(values[len(values)-1])
	return nil
}
