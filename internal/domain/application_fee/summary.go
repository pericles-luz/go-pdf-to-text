package application_fee

import (
	"fmt"
	"log"
	"strings"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
)

type Summary struct {
	localExecutionNumber string
	executionNumber      string
	mainProcess          string
	table                []*Line
	total                *TotalLine
}

func NewSummary() *Summary {
	return &Summary{}
}

func (s *Summary) ExecutionNumber() string {
	return s.executionNumber
}

func (s *Summary) LocalExecutionNumber() string {
	return s.localExecutionNumber
}

func (s *Summary) MainProcess() string {
	return s.mainProcess
}

func (s *Summary) Table() []*Line {
	return s.table
}

func (s *Summary) SetExecutionNumber(executionNumber string) {
	s.executionNumber = executionNumber
}

func (s *Summary) SetMainProcess(mainProcess string) {
	s.mainProcess = mainProcess
}

func (s *Summary) SetLocalExecutionNumber(localExecutionNumber string) {
	s.localExecutionNumber = localExecutionNumber
}

func (s *Summary) AddToTable(line *Line) error {
	if err := line.Validate(); err != nil {
		return err
	}
	for _, l := range s.table {
		if l.Sequence() == line.Sequence() {
			return application.ErrLinhaJaExistente
		}
	}
	s.table = append(s.table, line)
	return nil
}

func (s *Summary) SetTotal(total *TotalLine) error {
	if err := total.Validate(); err != nil {
		return fmt.Errorf("total invalido em summary: %w", err)
	}
	s.total = total
	return nil
}

func (s *Summary) Validate() error {
	if s.ExecutionNumber() == "" {
		return application.ErrNumeroExecucaoInvalido
	}
	if s.MainProcess() == "" {
		return application.ErrProcessoPrincipalInvalido
	}
	if len(s.Table()) == 0 {
		return application.ErrTabelaInvalida
	}
	if s.total == nil {
		log.Println("total zerado em summary")
		return application.ErrTotalInvalido
	}
	if err := s.total.Validate(); err != nil {
		return err
	}
	return s.CalculateTotal().Validate()
}

func (s *Summary) Parse(lines []string) error {
	if len(lines) < 3 {
		return application.ErrLinhaInvalida
	}
	number, err := parse.FindExecutionNumber(lines)
	if err != nil {
		return err
	}
	s.SetExecutionNumber(number)
	main, err := parse.FindMainNumber(lines)
	if err != nil {
		return err
	}
	s.SetMainProcess(main)
	localExecutionNumber, err := parse.FindLocalExecutionNumber(lines)
	if err != nil {
		return err
	}
	s.SetLocalExecutionNumber(localExecutionNumber)
	totalLine := NewTotalLine()
	for _, source := range lines {
		if err := totalLine.Parse(source); err == nil {
			return s.SetTotal(totalLine)
		}
		line := NewLine()
		if err := line.Parse(source); err == nil {
			s.AddToTable(line)
		}
	}
	return nil
}

func (s *Summary) IsFeesFile(lines []string) bool {
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower("Planilha Consolidada - Honor")) {
			return true
		}
	}
	return false
}

func (s *Summary) HasLines() bool {
	for _, line := range s.Table() {
		if line.Useble() {
			return true
		}
	}
	return false
}

func (s *Summary) RemoveFromTable(line *Line) {
	for index, l := range s.table {
		if l.CPF() == line.CPF() {
			s.remove(index)
			return
		}
	}
}

func (s *Summary) remove(index int) {
	s.table = append(s.table[:index], s.table[index+1:]...)
}

func (s *Summary) CalculateTotal() *TotalLine {
	totalLine := NewTotalLine()
	for _, line := range s.Table() {
		if !line.Useble() {
			continue
		}
		totalLine.Add(line)
	}
	return totalLine
}
