package application_fee

import (
	"log"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
)

type Line struct {
	sequence uint16
	cpf      string
	name     string
	uniqueID string
	main     uint64
	interest uint64
	total    uint64
	fees     uint64
	status   string
}

func NewLine() *Line {
	return &Line{}
}

func (l *Line) Sequence() uint16 {
	return l.sequence
}

func (l *Line) CPF() string {
	return l.cpf
}

func (l *Line) Name() string {
	return l.name
}

func (l *Line) UniqueID() string {
	return l.uniqueID
}

func (l *Line) Main() uint64 {
	return l.main
}

func (l *Line) Interest() uint64 {
	return l.interest
}

func (l *Line) Total() uint64 {
	return l.total
}

func (l *Line) Fees() uint64 {
	return l.fees
}

func (l *Line) Status() string {
	return l.status
}

func (l *Line) SetSequence(sequence uint16) {
	l.sequence = sequence
}

func (l *Line) SetCPF(cpf string) {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	l.cpf = cpf
}

func (l *Line) SetName(name string) {
	l.name = name
}

func (l *Line) SetUniqueID(uniqueID string) {
	l.uniqueID = uniqueID
}

func (l *Line) SetMain(main uint64) {
	l.main = main
}

func (l *Line) SetInterest(interest uint64) {
	l.interest = interest
}

func (l *Line) SetTotal(total uint64) {
	l.total = total
}

func (l *Line) SetFees(fees uint64) {
	l.fees = fees
}

func (l *Line) SetStatus(status string) {
	l.status = status
}

func (l *Line) Validate() error {
	if l.CPF() == "" {
		return application.ErrCpfInvalido
	}
	if l.UniqueID() == "" {
		return application.ErrIdUnicaInvalido
	}
	if l.Main() == 0 {
		return application.ErrVencimentoBasicoInvalido
	}
	if l.Interest() == 0 {
		return application.ErrJurosMoraInvalido
	}
	if l.Total() == 0 {
		log.Println("total zerado em line para CPF", l.CPF())
		return application.ErrTotalInvalido
	}
	if l.Fees() == 0 {
		return application.ErrHonorarioInvalido
	}
	if err := l.ValidateSum(); err != nil {
		return err
	}
	return nil
}

func (l *Line) ValidateSum() error {
	if extract.MuitoDiferente(l.Total(), l.Main()+l.Interest()) {
		return application.ErrTotalNaoBate
	}
	if extract.MuitoDiferente(l.Fees(), l.Total()*10/100) {
		return application.ErrTotalHonorarioNaoBate
	}
	return nil
}

// 2     RONALDO ASSUNCAO JACOMINI                       176.757.826-15         1495275                         0,00               0,00           0,00                0,00   Não Consta CPF na Lista
func (l *Line) Parse(line string) error {
	line = strings.TrimSpace(line)
	for strings.Contains(line, "    ") {
		line = strings.ReplaceAll(line, "    ", "   ")
	}
	if len(line) < 50 {
		return application.ErrLinhaInvalida
	}
	values := strings.Split(line, "   ")
	if len(values) < 9 {
		return application.ErrLinhaInvalida
	}
	l.SetSequence(uint16(utils.StringToInt(strings.TrimSpace(values[0]))))
	l.SetName(strings.TrimSpace(values[1]))
	l.SetCPF(strings.TrimSpace(values[2]))
	l.SetUniqueID(strings.TrimSpace(values[3]))
	l.SetMain(extract.StringToInt(strings.TrimSpace(values[4])))
	l.SetInterest(extract.StringToInt(strings.TrimSpace(values[5])))
	l.SetTotal(extract.StringToInt(strings.TrimSpace(values[6])))
	l.SetFees(extract.StringToInt(strings.TrimSpace(values[7])))
	l.SetStatus(strings.TrimSpace(values[8]))
	return l.Validate()
}
