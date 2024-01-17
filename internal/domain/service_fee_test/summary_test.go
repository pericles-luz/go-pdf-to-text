package service_fee_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/stretchr/testify/require"
)

func TestSummaryMustParse(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 1)
}

func TestSummaryMustValidateTotal(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	require.NoError(t, err)
	require.Equal(t, uint64(190371314), summary.CalculateTotal().Main())
	require.Equal(t, uint64(246568023), summary.CalculateTotal().Interest())
	require.Equal(t, uint64(436939336), summary.CalculateTotal().Total())
	require.Equal(t, uint64(43693935), summary.CalculateTotal().Fees())
}

func TestSummaryMustNotParse(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/009-11804009-C.txt")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 0)
}
