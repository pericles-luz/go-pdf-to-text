package convert_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/pericles-luz/go-pdf-to-text/pkg/convert"
	"github.com/stretchr/testify/require"
)

func TestWalkMustProcessDirectoryWithCalculo(t *testing.T) {
	err := convert.Walk(utils.GetBaseDirectory("pdf"), &service.Calculo{})
	require.NoError(t, err)
}

func TestWalkMustProcessPDFDirectory(t *testing.T) {
	t.Skip("This test takes too long to run")
	err := convert.Walk("/mnt/c/Users/peric/Downloads/PDFs", &service.Calculo{})
	require.NoError(t, err)
}

func TestWalkMustProcessTestDirectoryWithSummary(t *testing.T) {
	processor := service_fee.NewSummary()
	err := convert.Walk(utils.GetBaseDirectory("pdf"), processor)
	require.NoError(t, err)
	require.Len(t, processor.Summaries(), 6)
}

func TestWalkMustProcessRealDirectoryWithSummary(t *testing.T) {
	t.Skip("This test takes too long to run")
	processor := service_fee.NewSummary()
	err := convert.Walk("/mnt/c/Users/peric/Downloads/PDFs", processor)
	require.NoError(t, err)
	require.Len(t, processor.Summaries(), 2)
}

func TestWalkMustProcessTestDirectoryWithSummaryAndGenerateExcel(t *testing.T) {
	processor := service_fee.NewSummary()
	err := convert.Walk(utils.GetBaseDirectory("pdf"), processor)
	require.NoError(t, err)
	require.NoError(t, convert.GenerateSuccumbence(utils.GetBaseDirectory("pdf"), processor))
	require.FileExists(t, utils.GetBaseDirectory("pdf")+"/honorarios consolidados.xlsx")
	require.NoError(t, os.Remove(utils.GetBaseDirectory("pdf")+"/honorarios consolidados.xlsx"))
}

func TestWalkMustProcessRealDirectoryWithSummaryAndGenerateExcel(t *testing.T) {
	t.Skip("This test takes too long to run")
	processor := service_fee.NewSummary()
	err := convert.Walk("/mnt/c/Users/peric/Downloads/PDFs", processor)
	require.NoError(t, err)
	require.Len(t, processor.Summaries(), 2)
}
