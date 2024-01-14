package excel_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/pericles-luz/go-pdf-to-text/internal/excel"
	"github.com/stretchr/testify/require"
)

func TestSuccumbenceMustGenerateFileOfOneSource(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 1)
	succumbence := excel.NewSuccumbence(summary.Summaries(), utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	err = succumbence.ProcessFile()
	require.NoError(t, err)
	require.FileExists(t, utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx")
	require.NoError(t, os.Remove(utils.GetBaseDirectory("pdf")+"/002-Honorários.xlsx"))
}
