package extract

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/xuri/excelize/v2"
)

func MuitoDiferente(a, b uint64) bool {
	if a > b {
		return a-b > 1000
	}
	return b-a > 1000
}

func ReadLinesFromFile(path string) ([]string, error) {
	if filepath.Ext(path) == ".txt" {
		return ReadLinesFromTextFile(path)
	}
	if filepath.Ext(path) == ".xlsm" {
		return ReadLinesFromExcelFile(path)
	}
	return nil, errors.New("arquivo inválido")
}

func ReadLinesFromTextFile(path string) ([]string, error) {
	if filepath.Ext(path) != ".txt" {
		return nil, errors.New("arquivo inválido")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadLinesFromExcelFile(path string) ([]string, error) {
	if filepath.Ext(path) != ".xlsm" {
		return nil, errors.New("arquivo inválido")
	}
	if !utils.FileExists(path) {
		return nil, errors.New("arquivo não encontrado")
	}
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetName := "Honorários"
	feeSheet, err := f.GetSheetIndex(sheetName)
	if err != nil {
		return nil, err
	}
	if feeSheet == 0 {
		return nil, errors.New("sheet não encontrada")
	}
	f.SetActiveSheet(feeSheet)
	excludeCols := []string{"J", "I", "H", "G", "F", "E", "D", "C", "B", "A"}
	for _, col := range excludeCols {
		colVisible, err := f.GetColVisible(sheetName, col)
		if err != nil {
			return nil, err
		}
		if !colVisible {
			f.RemoveCol(sheetName, col)
			continue
		}
		if colWidth, err := f.GetColWidth(sheetName, col); err == nil {
			if colWidth < 5 {
				f.RemoveCol(sheetName, col)
				continue
			}
		}
	}
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	for i, row := range rows {
		visible, err := f.GetRowVisible(sheetName, i+1)
		if err != nil {
			continue
		}
		if !visible {
			continue
		}
		lines = append(lines, strings.Join(row, "       "))
	}
	return lines, nil
}

func StringToInt(s string) uint64 {
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "-", "")
	return uint64(utils.StringToInt(s))
}
