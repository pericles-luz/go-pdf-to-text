package excel

import (
	"fmt"
	"log"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
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
	styleIDs    [10]int
	rows        [][]string
	summaries   []*application_fee.Summary
	existents   []*application_fee.Line
	file        *excelize.File
	mustHave    *application_fee.MustHaveList
	currentLine int
	countLines  int
}

func NewSuccumbence(s []*application_fee.Summary, o string) *Succumbence {
	result := &Succumbence{
		summaries:  s,
		outputPath: o,
		mustHave:   application_fee.NewMustHaveList(),
	}
	result.existents = make([]*application_fee.Line, 0)
	result.currentLine = 1
	result.countLines = 0
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
		return fmt.Errorf("summary invalido em writeheader: %w", err)
	}
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine+1), s.styleIDs[StyleTitle])
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Execução")
	s.getFile().SetCellInt(s.sheetName, s.cell("B", s.currentLine), int(summary.LocalExecutionNumber()))
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

func (s *Succumbence) writeLine(line *application_fee.Line) error {
	if err := line.Validate(); err != nil {
		return fmt.Errorf("linha invalida em writeline: %w", err)
	}
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleTableLine])
	s.getFile().SetCellInt(s.sheetName, s.cell("A", s.currentLine), s.countLines)
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

func (s *Succumbence) writeFooter(total *application_fee.TotalLine, countingLines int) error {
	if err := total.Validate(); err != nil {
		return fmt.Errorf("total invalido em writefooter: %w", err)
	}
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleTableFooter])
	s.getFile().SetCellStyle(s.sheetName, s.cell("F", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleFooterValue])
	s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), fmt.Sprintf("%d beneficiários na execução", countingLines))
	s.getFile().MergeCell(s.sheetName, s.cell("B", s.currentLine), s.cell("E", s.currentLine))
	s.getFile().SetCellFloat(s.sheetName, s.cell("F", s.currentLine), float64(total.Main())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("G", s.currentLine), float64(total.Interest())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("H", s.currentLine), float64(total.Total())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("I", s.currentLine), float64(total.Fees())/100, 2, 64)
	s.currentLine += 3
	return nil
}

func (s *Succumbence) writeSummary(summary *application_fee.Summary) error {
	if !summary.HasLines() {
		return nil
	}
	if err := s.writeHeader(summary); err != nil {
		return err
	}
	countLines := 0
	for _, line := range summary.Table() {
		if !line.Useble() {
			s.existents = append(s.existents, line)
			continue
		}
		s.countLines++
		if err := s.writeLine(line); err != nil {
			return err
		}
		countLines++
	}
	if err := s.writeFooter(summary.CalculateTotal(), countLines); err != nil {
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
	hasSummary := false
	err := s.mustHave.FromJSONFile(utils.GetBaseDirectory("assets") + "/mustHave.json")
	if err != nil {
		return err
	}
	for _, mustHave := range s.mustHave.List() {
		mustHave.SetClassification(application_fee.MustHaveNotFound)
	}
	for _, summary := range s.summaries {
		for _, line := range summary.Table() {
			// log.Println("CPF:", line.CPF())
			if s.ExistsOnPrevious(line.CPF()) {
				s.existents = append(s.existents, line)
				line.SetUseble(false)
				continue
			} else {
				s.mustHave.SetClassification(line.CPF(), application_fee.MustHaveFound)
			}
		}
		if !summary.HasLines() {
			continue
		}
		if err := s.writeSummary(summary); err != nil {
			return err
		}
		hasSummary = true
		s.currentLine++
	}
	if !hasSummary {
		return nil
	}
	if err := s.writeTotal(s.CalculateTotal()); err != nil {
		return err
	}
	if err := s.writeOutOfList(); err != nil {
		return err
	}
	hasSummary = false
	s.sheetName = "Listagens Anteriores"
	s.countLines = 0
	s.currentLine = 1
	s.getFile().NewSheet(s.sheetName)
	for _, summary := range s.summaries {
		for _, line := range summary.Table() {
			line.SetUseble(!line.Useble())
			if line.Useble() {
				s.mustHave.SetClassification(line.CPF(), application_fee.MustHaveAlreadyPaid)
			}
		}
		if !summary.HasLines() {
			log.Println("summary sem linhas em listagens anteriores: ", summary.LocalExecutionNumber())
			continue
		}
		if err := s.writeSummary(summary); err != nil {
			return err
		}
		hasSummary = true
		s.currentLine++
	}
	if hasSummary {
		if err := s.writeTotal(s.CalculateTotal()); err != nil {
			return err
		}
	}
	// if err := s.writeExistents(); err != nil {
	// 	return err
	// }
	if err := s.writeMustHave(); err != nil {
		return err
	}
	if err := s.write(s.outputPath); err != nil {
		return err
	}
	return nil
}

func (s *Succumbence) writeTotal(total *application_fee.TotalLine) error {
	if err := total.Validate(); err != nil {
		return fmt.Errorf("total invalido em writetotal: %w", err)
	}
	s.currentLine += 3
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleTableFooter])
	s.getFile().SetCellStyle(s.sheetName, s.cell("F", s.currentLine), s.cell("I", s.currentLine), s.styleIDs[StyleFooterValue])
	s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), fmt.Sprintf("%d beneficiários totais", s.countLines))
	s.getFile().MergeCell(s.sheetName, s.cell("B", s.currentLine), s.cell("D", s.currentLine))
	s.getFile().SetCellStr(s.sheetName, s.cell("E", s.currentLine), "TOTAL GERAL")
	s.getFile().SetCellFloat(s.sheetName, s.cell("F", s.currentLine), float64(total.Main())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("G", s.currentLine), float64(total.Interest())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("H", s.currentLine), float64(total.Total())/100, 2, 64)
	s.getFile().SetCellFloat(s.sheetName, s.cell("I", s.currentLine), float64(total.Fees())/100, 2, 64)
	return nil
}

func (s *Succumbence) close() {
	if err := s.getFile().Close(); err != nil {
		fmt.Println(err)
	}
}

func (s *Succumbence) LoadPrevious(file string) error {
	if !utils.FileExists(file) {
		return application.ErrArquivoNaoEncontrado
	}
	f, err := excelize.OpenFile(file)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetName := f.GetSheetName(f.GetActiveSheetIndex())
	s.rows, err = f.GetRows(sheetName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Succumbence) ExistsOnPrevious(cpf string) bool {
	cpf = utils.GetOnlyNumbers(cpf)
	if !utils.ValidateCPF(cpf) {
		return false
	}
	for _, row := range s.rows {
		if len(row) < 5 {
			continue
		}
		if utils.GetOnlyNumbers(row[4]) == cpf {
			return true
		}
	}
	return false
}

// func (s *Succumbence) writeExistents() error {
// 	if len(s.existents) == 0 {
// 		return nil
// 	}
// 	s.currentLine = 1
// 	s.sheetName = "Listagens Anteriores"
// 	s.getFile().NewSheet(s.sheetName)
// 	s.getFile().SetColWidth(s.sheetName, "A", "B", 10)
// 	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "Nº PESSOAS")
// 	s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), "SEQ")
// 	s.getFile().SetColWidth(s.sheetName, "C", "C", 35)
// 	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), "NOME")
// 	s.getFile().SetColWidth(s.sheetName, "D", "E", 15)
// 	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), "CPF")
// 	s.getFile().SetCellStr(s.sheetName, s.cell("E", s.currentLine), "IDENTIFICADOR ÚNICO")
// 	s.getFile().SetColWidth(s.sheetName, "F", "I", 20)
// 	s.getFile().SetCellStr(s.sheetName, s.cell("F", s.currentLine), "PRINCIPAL ATUALIZADO (COM DESÁGIO)")
// 	s.getFile().SetCellStr(s.sheetName, s.cell("G", s.currentLine), "JUROS DE MORA ATUALIZADO (COM DESÁGIO)")
// 	s.getFile().SetCellStr(s.sheetName, s.cell("H", s.currentLine), "TOTAL COM DESÁGIO")
// 	s.getFile().SetCellStr(s.sheetName, s.cell("I", s.currentLine), "HONORÁRIOS DE SUCUMBÊNCIA 10%")
// 	s.currentLine++
// 	total := application_fee.NewTotalLine()
// 	total.SetMain(0)
// 	total.SetInterest(0)
// 	total.SetTotal(0)
// 	total.SetFees(0)
// 	s.countLines = 0
// 	for _, line := range s.existents {
// 		line.SetUseble(!line.Useble())
// 		if !line.Useble() {
// 			continue
// 		}
// 		if err := s.writeLine(line); err != nil {
// 			return err
// 		}
// 		total.Add(line)
// 		s.countLines++
// 	}
// 	fmt.Println(len(s.existents))
// 	fmt.Println("existents:", len(s.existents))
// 	fmt.Println("principal:", total.Main())
// 	fmt.Println("interest:", total.Interest())
// 	fmt.Println("total:", total.Total())
// 	fmt.Println("fees:", total.Fees())
// 	if err := s.writeFooter(total, len(s.existents)); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *Succumbence) CalculateTotal() *application_fee.TotalLine {
	total := application_fee.NewTotalLine()
	for _, summary := range s.summaries {
		total.Add(summary.CalculateTotal())
	}
	return total
}

func (s *Succumbence) writeMustHave() error {
	s.currentLine = 1
	s.sheetName = "Lista Sucumbência"
	s.getFile().NewSheet(s.sheetName)
	s.getFile().SetColWidth(s.sheetName, "A", "A", 35)
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "NOME")
	s.getFile().SetColWidth(s.sheetName, "B", "B", 15)
	s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), "CPF")
	s.getFile().SetColWidth(s.sheetName, "C", "C", 35)
	s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), "PROCESSO")
	s.getFile().SetColWidth(s.sheetName, "D", "D", 15)
	s.getFile().SetCellStr(s.sheetName, s.cell("D", s.currentLine), "EXECUÇÃO")
	s.getFile().SetColWidth(s.sheetName, "E", "E", 15)
	s.getFile().SetCellStr(s.sheetName, s.cell("E", s.currentLine), "SEQUÊNCIA")
	s.getFile().SetColWidth(s.sheetName, "F", "F", 15)
	s.getFile().SetCellStr(s.sheetName, s.cell("F", s.currentLine), "STATUS")
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("F", s.currentLine), s.styleIDs[StyleTableHeader])
	s.currentLine++
	for _, mustHave := range s.mustHave.List() {
		s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), mustHave.Name)
		s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), mustHave.CPF)
		s.getFile().SetCellStr(s.sheetName, s.cell("C", s.currentLine), mustHave.ProcessNumber)
		s.getFile().SetCellInt(s.sheetName, s.cell("D", s.currentLine), int(mustHave.Execution))
		s.getFile().SetCellInt(s.sheetName, s.cell("E", s.currentLine), int(mustHave.Sequence))
		s.getFile().SetCellStr(s.sheetName, s.cell("F", s.currentLine), mustHave.Classification())
		s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("F", s.currentLine), s.styleIDs[StyleTableLine])
		s.currentLine++
		fmt.Println(mustHave.CPF, mustHave.Classification())
	}
	return nil
}

func (s *Succumbence) writeOutOfList() error {
	s.currentLine = 1
	s.sheetName = "Lista Fora da Sucumbência"
	s.getFile().NewSheet(s.sheetName)
	s.getFile().SetColWidth(s.sheetName, "A", "A", 35)
	s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), "NOME")
	s.getFile().SetColWidth(s.sheetName, "B", "B", 15)
	s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), "CPF")
	s.getFile().SetColWidth(s.sheetName, "C", "C", 35)
	s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("C", s.currentLine), s.styleIDs[StyleTableHeader])
	s.currentLine++
	for _, sumary := range s.summaries {
		for _, line := range sumary.Table() {
			if !line.Useble() {
				continue
			}
			if s.mustHave.Has(line.CPF()) {
				continue
			}
			s.getFile().SetCellStr(s.sheetName, s.cell("A", s.currentLine), line.Name())
			s.getFile().SetCellStr(s.sheetName, s.cell("B", s.currentLine), line.CPF())
			s.getFile().SetCellStyle(s.sheetName, s.cell("A", s.currentLine), s.cell("C", s.currentLine), s.styleIDs[StyleTableLine])
			s.currentLine++
		}
	}
	return nil
}
