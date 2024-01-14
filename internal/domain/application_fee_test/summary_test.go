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

func TestSummaryMustParseWithOtherInfotmations(t *testing.T) {
	lines := readOriginFile(t, "038-Honorários.txt")
	summary := application_fee.NewSummary()
	err := summary.Parse(lines)
	require.NoError(t, err)
	require.NoError(t, summary.Validate())
}

func TestSummaryMustParseWithOtherInfotmations2(t *testing.T) {
	lines := readOriginFile(t, "035-Honorários.txt")
	summary := application_fee.NewSummary()
	err := summary.Parse(lines)
	require.NoError(t, err)
	require.NoError(t, summary.Validate())
}

func TestSummaryMustParseWithNumberWithoutLabel(t *testing.T) {
	lines := readOriginFile(t, "069-Honorários.txt")
	summary := application_fee.NewSummary()
	err := summary.Parse(lines)
	require.NoError(t, err)
	require.NoError(t, summary.Validate())
}

func TestSummaryMustParseWithNumberWithoutTotal(t *testing.T) {
	lines := readOriginFile(t, "202-Honorários.txt")
	summary := application_fee.NewSummary()
	err := summary.Parse(lines)
	require.NoError(t, err)
	require.NoError(t, summary.Validate())
}

func TestSummaryMustDetectFeesFile(t *testing.T) {
	lines := readOriginFile(t, "002-Honorários.txt")
	summary := application_fee.NewSummary()
	require.True(t, summary.IsFeesFile(lines))
}

func TestSummaryMustNotDetectFeesFile(t *testing.T) {
	lines := readOriginFile(t, "086-11861592-C.txt")
	summary := application_fee.NewSummary()
	require.False(t, summary.IsFeesFile(lines))
}
