package excel

import (
	"fmt"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application_fee"
	"github.com/xuri/excelize/v2"
)

type Succumbence struct {
	sheetName   string
	outputPath  string
	summaries   []*application_fee.Summary
	file        *excelize.File
	currentLine int
}

func NewSuccumbence(s []*application_fee.Summary, o string) *Succumbence {
	result := &Succumbence{
		summaries:  s,
		outputPath: o,
	}
	result.currentLine = 1
	result.sheetName = "Listagens"
	return result
}

func (s *Succumbence) getFile() *excelize.File {
	if s.file == nil {
		s.file = excelize.NewFile()
		s.file.NewSheet(s.sheetName)
	}
	return s.file
}

func (s *Succumbence) cell(column string, line int) string {
	return fmt.Sprintf("%s%d", column, line)
}

func (s *Succumbence) write(path string) error {
	return s.getFile().SaveAs(path)
}

func (s *Succumbence) writeHeader(summary *application_fee.Summary) error {
	if err := summary.Validate(); err != nil {
		return err
	}
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Execução")
	s.getFile().SetCellInt(s.sheetName, s.cell("B", s.currentLine), s.currentLine)
	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), "Cumprimento de Sentença nº")
	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), summary.ExecutionNumber())
	s.currentLine++
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Planilha Consolidada")
	s.currentLine++
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Nº PESSOAS")
	s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), "SEQ")
	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), "NOME")
	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), "CPF")
	s.getFile().SetCellStr(s.sheetName, s.cell("E", s.currentLine), "IDENTIFICADOR ÚNICO")
	s.getFile().SetCellStr(s.sheetName, s.cell("F", s.currentLine), "PRINCIPAL ATUALIZADO (COM DESÁGIO)")
	s.getFile().SetCellStr(s.sheetName, s.cell("G", s.currentLine), "JUROS DE MORA ATUALIZADO (COM DESÁGIO)")
	s.getFile().SetCellStr(s.sheetName, s.cell("H", s.currentLine), "TOTAL COM DESÁGIO")
	s.getFile().SetCellStr(s.sheetName, s.cell("I", s.currentLine), "HONORÁRIOS DE SUCUMBÊNCIA 10%")
	s.currentLine++
	return nil
}

func (s *Succumbence) writeLine(line *application_fee.Line) error {
	if err := line.Validate(); err != nil {
		return err
	}
	s.getFile().SetCellInt(s.sheetName, s.cell("A", s.currentLine), 0)
	s.getFile().SetCellInt(s.sheetName, s.cell("B", s.currentLine), int(line.Sequence()))
	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), line.Name())
	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), line.CPF())
	s.getFile().SetCellStr(s.sheetName, s.cell("E", s.currentLine), line.UniqueID())
	s.getFile().SetCellFloat(s.sheetName, s.cell("F", s.currentLine), float64(line.Main())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("G", s.currentLine), float64(line.Interest())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("H", s.currentLine), float64(line.Total())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("I", s.currentLine), float64(line.Fees())/100, 2, 64)
	s.currentLine++
	return nil
}

func (s *Succumbence) writeFooter(summary *application_fee.Summary) error {
	if err := summary.Validate(); err != nil {
		return err
	}
	s.getFile().SetCellFloat(s.sheetName, s.cell("F", s.currentLine), float64(summary.Total().Main())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("G", s.currentLine), float64(summary.Total().Interest())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("H", s.currentLine), float64(summary.Total().Total())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("I", s.currentLine), float64(summary.Total().Fees())/100, 2, 64)
	s.currentLine++
	return nil
}

func (s *Succumbence) writeSummary(summary *application_fee.Summary) error {
	if err := s.writeHeader(summary); err != nil {
		return err
	}
	for _, line := range summary.Table() {
		if err := s.writeLine(line); err != nil {
			return err
		}
	}
	if err := s.writeFooter(summary); err != nil {
		return err
	}
	return nil
}

func (s *Succumbence) ProcessFile() error {
	defer s.close()
	for _, summary := range s.summaries {
		if err := s.writeSummary(summary); err != nil {
			return err
		}
		s.currentLine++
	}
	if err := s.write(s.outputPath); err != nil {
		return err
	}
	return nil
}

func (s *Succumbence) close() {
	if err := s.getFile().Close(); err != nil {
		fmt.Println(err)
	}
}
