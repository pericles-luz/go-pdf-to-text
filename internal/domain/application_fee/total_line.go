package application_fee

import (
	"log"
	"strings"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
)

type TotalLine struct {
	main     uint64
	interest uint64
	total    uint64
	fees     uint64
}

func NewTotalLine() *TotalLine {
	return &TotalLine{}
}

func (l *TotalLine) Main() uint64 {
	return l.main
}

func (l *TotalLine) Interest() uint64 {
	return l.interest
}

func (l *TotalLine) Total() uint64 {
	return l.total
}

func (l *TotalLine) Fees() uint64 {
	return l.fees
}

func (l *TotalLine) SetMain(main uint64) {
	l.main = main
}

func (l *TotalLine) SetInterest(interest uint64) {
	l.interest = interest
}

func (l *TotalLine) SetTotal(total uint64) {
	l.total = total
}

func (l *TotalLine) SetFees(fees uint64) {
	l.fees = fees
}

func (l *TotalLine) Validate() error {
	if l.Main() == 0 {
		return application.ErrTotalPrincipalInvalido
	}
	if l.Interest() == 0 {
		return application.ErrTotalJurosInvalido
	}
	if l.Total() == 0 {
		log.Println("Total serado em totalLine:", l.Total())
		return application.ErrTotalInvalido
	}
	if l.Fees() == 0 {
		return application.ErrTotalHonorarioInvalido
	}
	if err := l.ValidateSum(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (l *TotalLine) ValidateSum() error {
	if extract.MuitoDiferente(l.Total(), l.Main()+l.Interest()) {
		return application.ErrTotalNaoBate
	}
	if extract.MuitoDiferente(l.Fees(), l.Total()*10/100) {
		return application.ErrTotalHonorarioNaoBate
	}
	return nil
}

func (l *TotalLine) Parse(line string) error {
	line = strings.ReplaceAll(line, "-", "")
	line = strings.TrimSpace(line)
	if len(line) < 50 {
		return application.ErrLinhaInvalida
	}
	if !strings.HasPrefix(line, "TOTAL GERAL") {
		return application.ErrLinhaNaoEhTotalGeral
	}
	for strings.Contains(line, "    ") {
		line = strings.ReplaceAll(line, "    ", "   ")
	}
	values := strings.Split(line, "   ")
	if len(values) < 5 {
		return application.ErrLinhaInvalida
	}
	l.SetMain(extract.StringToInt(strings.TrimSpace(values[1])))
	l.SetInterest(extract.StringToInt(strings.TrimSpace(values[2])))
	l.SetTotal(extract.StringToInt(strings.TrimSpace(values[3])))
	l.SetFees(extract.StringToInt(strings.TrimSpace(values[4])))
	return l.Validate()
}
