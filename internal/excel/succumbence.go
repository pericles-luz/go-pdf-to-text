package excel

import (
	"fmt"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application_fee"
	"github.com/xuri/excelize/v2"
)

const (
	StyleTableHeader = iota
	StyleTableLine
	StyleTableFooter
	StyleTableValue
	StyleFooterValue
	StyleTitle
)

type Succumbence struct {
	sheetName   string
	outputPath  string
	summaries   []*application_fee.Summary
	styleIDs    [10]int
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
	result.styleIDs[StyleTableHeader] = result.tableHeaderStyle()
	result.styleIDs[StyleTableLine] = result.tableLineStyle()
	result.styleIDs[StyleTableFooter] = result.tableFooterStyle()
	result.styleIDs[StyleTableValue] = result.tableLineValueStyle()
	result.styleIDs[StyleFooterValue] = result.tableFooterValueStyle()
	result.styleIDs[StyleTitle] = result.tableTitleStyle()
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

func (s *Succumbence) tableTitleStyle() int {
	styleID, err := s.getFile().NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "000000",
			Size:  12,
		},
	},
	)
	if err != nil {
		fmt.Println(err)
	}
	return styleID
}

func (s *Succumbence) tableHeaderStyle() int {
	styleID, err := s.getFile().NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type: "pattern",
			Color: []string{
				"#EEEE00",
			},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "000000",
			Size:  11,
		},
		Alignment: &excelize.Alignment{
			WrapText: true,
			Vertical: "center",
		},
	},
	)
	if err != nil {
		fmt.Println(err)
	}
	return styleID
}

func (s *Succumbence) tableLineStyle() int {
	styleID, err := s.getFile().NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type: "pattern",
			Color: []string{
				"#CDDFFC",
			},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "000000",
			Size:  9,
		},
	},
	)
	if err != nil {
		fmt.Println(err)
	}
	return styleID
}

func (s *Succumbence) tableFooterStyle() int {
	styleID, err := s.getFile().NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type: "pattern",
			Color: []string{
				"#62B0FF",
			},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "000000",
			Size:  11,
		},
	},
	)
	if err != nil {
		fmt.Println(err)
	}
	return styleID
}

func (s *Succumbence) tableLineValueStyle() int {
	styleID, err := s.getFile().NewStyle(&excelize.Style{
		NumFmt: 353,
		Fill: excelize.Fill{
			Type: "pattern",
			Color: []string{
				"#CDDFFC",
			},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "000000",
			Size:  9,
		},
	},
	)
	if err != nil {
		fmt.Println(err)
	}
	return styleID
}

func (s *Succumbence) tableFooterValueStyle() int {
	styleID, err := s.getFile().NewStyle(&excelize.Style{
		NumFmt: 353,
		Fill: excelize.Fill{
			Type: "pattern",
			Color: []string{
				"#62B0FF",
			},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "000000",
			Size:  11,
		},
	},
	)
	if err != nil {
		fmt.Println(err)
	}
	return styleID
}

func (s *Succumbence) writeHeader(summary *application_fee.Summary) error {
	if err := summary.Validate(); err != nil {
		return err
	}
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine+1), s.styleIDs[StyleTitle])
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Execução")
	s.getFile().SetCellInt(s.sheetName, s.cell("B", s.currentLine), s.currentLine)
	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), "Cumprimento de Sentença nº")
	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), summary.ExecutionNumber())
	s.currentLine++
	s.getFile().MergeCell(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine))
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Planilha Consolidada")
	s.currentLine++
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleTableHeader])
	s.getFile().SetColWidth(s.sheetName, "A", "B", 10)
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Nº PESSOAS")
	s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), "SEQ")
	s.getFile().SetColWidth(s.sheetName, "C", "C", 35)
	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), "NOME")
	s.getFile().SetColWidth(s.sheetName, "D", "E", 15)
	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), "CPF")
	s.getFile().SetCellStr(s.sheetName, s.cell("E", s.currentLine), "IDENTIFICADOR ÚNICO")
	s.getFile().SetColWidth(s.sheetName, "F", "I", 20)
	s.getFile().SetCellStr(s.sheetName, s.cell("F", s.currentLine), "PRINCIPAL ATUALIZADO (COM DESÁGIO)")
	s.getFile().SetCellStr(s.sheetName, s.cell("G", s.currentLine), "JUROS DE MORA ATUALIZADO (COM DESÁGIO)")
	s.getFile().SetCellStr(s.sheetName, s.cell("H", s.currentLine), "TOTAL COM DESÁGIO")
	s.getFile().SetCellStr(s.sheetName, s.cell("I", s.currentLine), "HONORÁRIOS DE SUCUMBÊNCIA 10%")
	s.currentLine++
	return nil
}

func (s *Succumbence) writeLine(line *application_fee.Line, index int) error {
	if err := line.Validate(); err != nil {
		return err
	}
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleTableLine])
	s.getFile().SetCellInt(s.sheetName, s.cell("A", s.currentLine), index)
	s.getFile().SetCellInt(s.sheetName, s.cell("B", s.currentLine), int(line.Sequence()))
	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), line.Name())
	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), line.CPF())
	s.getFile().SetCellStr(s.sheetName, s.cell("E", s.currentLine), line.UniqueID())
	s.getFile().SetCellStyle(s.sheetName, s.cell("F", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleTableValue])
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
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleTableFooter])
	s.getFile().SetCellStyle(s.sheetName, s.cell("F", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleFooterValue])
	s.getFile().SetCellFloat(s.sheetName, s.cell("F", s.currentLine), float64(summary.Total().Main())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("G", s.currentLine), float64(summary.Total().Interest())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("H", s.currentLine), float64(summary.Total().Total())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("I", s.currentLine), float64(summary.Total().Fees())/100, 2, 64)
	s.currentLine += 3
	return nil
}

func (s *Succumbence) writeSummary(summary *application_fee.Summary) error {
	if err := s.writeHeader(summary); err != nil {
		return err
	}
	for i, line := range summary.Table() {
		if err := s.writeLine(line, i+1); err != nil {
			return err
		}
	}
	if err := s.writeFooter(summary); err != nil {
		return err
	}
	index, err := s.getFile().GetSheetIndex(s.sheetName)
	if err != nil {
		return err
	}
	s.getFile().SetActiveSheet(index)
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
