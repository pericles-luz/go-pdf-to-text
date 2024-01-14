package application_fee_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application_fee"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
	"github.com/stretchr/testify/require"
)

func readOriginFile(t *testing.T, file string) []string {
	lines, err := extract.ReadLinesFromFile(utils.GetBaseDirectory("pdf") + "/" + file)
	require.NoError(t, err)
	return lines
}

func TestSummaryMustParse(t *testing.T) {
	lines := readOriginFile(t, "002-Honorários.txt")
	summary := application_fee.NewSummary()
	err := summary.Parse(lines)
	require.NoError(t, err)
}
