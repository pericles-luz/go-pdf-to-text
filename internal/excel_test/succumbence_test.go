package excel_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/pericles-luz/go-pdf-to-text/internal/excel"
	"github.com/stretchr/testify/require"
)

func TestSuccumbenceMustGenerateFileOfTwoEqualSources(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	require.NoError(t, err)
	err = summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 1)
	succumbence := excel.NewSuccumbence(summary.Summaries(), utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	succumbence.LoadPrevious(utils.GetBaseDirectory("pdf") + "/anterior.xlsx")
	err = succumbence.ProcessFile()
	require.NoError(t, err)
	require.FileExists(t, utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	require.NoError(t, os.Remove(utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx"))
}

func TestSuccumbenceMustGenerateFileOfTextAndExcelSources(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	require.NoError(t, err)
	err = summary.Parse(utils.GetBaseDirectory("pdf") + "/Execução 130.xlsm")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 2)
	succumbence := excel.NewSuccumbence(summary.Summaries(), utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	succumbence.LoadPrevious(utils.GetBaseDirectory("pdf") + "/anterior.xlsx")
	err = succumbence.ProcessFile()
	require.NoError(t, err)
	require.FileExists(t, utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	require.NoError(t, os.Remove(utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx"))
}

func TestSuccumbenceMustGenerateFileOfOneExcelSource(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/Execução 130.xlsm")
	require.NoError(t, err)
	// err = summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	// require.NoError(t, err)
	require.Len(t, summary.Summaries(), 1)
	succumbence := excel.NewSuccumbence(summary.Summaries(), utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	succumbence.LoadPrevious(utils.GetBaseDirectory("pdf") + "/anterior.xlsx")
	err = succumbence.ProcessFile()
	require.NoError(t, err)
	require.FileExists(t, utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	require.NoError(t, os.Remove(utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx"))
}

func TestSuccumbenceMustGenerateFileOfOneExcelSource2(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/Execução 193-SEQ.29.xlsm")
	require.NoError(t, err)
	// err = summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	// require.NoError(t, err)
	require.Len(t, summary.Summaries(), 1)
	succumbence := excel.NewSuccumbence(summary.Summaries(), utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	succumbence.LoadPrevious(utils.GetBaseDirectory("pdf") + "/anterior.xlsx")
	err = succumbence.ProcessFile()
	require.NoError(t, err)
	require.FileExists(t, utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	require.NoError(t, os.Remove(utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx"))
}

func TestSuccumbenceMustLoadPreviousAndFindCPF(t *testing.T) {
	succumbence := excel.NewSuccumbence(nil, "")
	err := succumbence.LoadPrevious(utils.GetBaseDirectory("pdf") + "/anterior.xlsx")
	require.NoError(t, err)
	require.True(t, succumbence.ExistsOnPrevious("153.542.101-00"))
	require.False(t, succumbence.ExistsOnPrevious("000.000.001-91"))
}

// arquivos onde o cabeçalho está como a seguir:
// "Planilha Consolidada- Honorários", com o hífen colado na palavra "Consolidada"
func TestSuccumbenceMustLoadFilesWithNewText(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/012-Exec-25.txt")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 1)
}
